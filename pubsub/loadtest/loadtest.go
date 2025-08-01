// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package loadtest implements load testing for pubsub,
// following the interface defined in https://github.com/GoogleCloudPlatform/pubsub/tree/master/load-test-framework/ .
//
// This package is experimental.
package loadtest

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"cloud.google.com/go/pubsub"
	pb "cloud.google.com/go/pubsub/loadtest/pb"
	"golang.org/x/time/rate"
)

type pubServerConfig struct {
	topic     *pubsub.Topic
	msgData   []byte
	batchSize int32
	ordered   bool
}

// PubServer is a dummy Pub/Sub server for load testing.
type PubServer struct {
	ID string

	cfg    atomic.Value
	seqNum int32
	pb.UnimplementedLoadtestWorkerServer
}

// Start starts the server.
func (l *PubServer) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	log.Println("received start")
	c, err := pubsub.NewClient(ctx, req.Project)
	if err != nil {
		return nil, err
	}
	l.init(c, req.Topic, req.GetPublisherOptions().GetMessageSize(), req.GetPublisherOptions().GetBatchSize(), req.GetPublisherOptions().GetBatchDuration().AsDuration(), false)
	log.Println("started")
	return &pb.StartResponse{}, nil
}

func (l *PubServer) init(c *pubsub.Client, topicName string, msgSize, batchSize int32, batchDur time.Duration, ordered bool) {
	topic := c.Topic(topicName)
	topic.PublishSettings = pubsub.PublishSettings{
		DelayThreshold: batchDur,
		CountThreshold: 950,
		ByteThreshold:  9500000,
	}
	topic.EnableMessageOrdering = ordered

	l.cfg.Store(pubServerConfig{
		topic:     topic,
		msgData:   bytes.Repeat([]byte{'A'}, int(msgSize)),
		batchSize: batchSize,
		ordered:   ordered,
	})
}

func (l *PubServer) publishBatch() ([]int64, error) {
	var cfg pubServerConfig
	if c, ok := l.cfg.Load().(pubServerConfig); ok {
		cfg = c
	} else {
		return nil, errors.New("config not loaded")
	}

	start := time.Now()
	latencies := make([]int64, cfg.batchSize)
	startStr := strconv.FormatInt(start.UnixNano()/1e6, 10)
	seqNum := atomic.AddInt32(&l.seqNum, cfg.batchSize) - cfg.batchSize

	rs := make([]*pubsub.PublishResult, cfg.batchSize)
	for i := int32(0); i < cfg.batchSize; i++ {
		msg := &pubsub.Message{
			Data: cfg.msgData,
			Attributes: map[string]string{
				"sendTime":       startStr,
				"clientId":       l.ID,
				"sequenceNumber": strconv.Itoa(int(seqNum + i)),
			},
		}
		if cfg.ordered {
			msg.OrderingKey = fmt.Sprintf("key-%d", seqNum+i)
		}
		rs[i] = cfg.topic.Publish(context.TODO(), msg)
	}
	for i, r := range rs {
		_, err := r.Get(context.Background())
		if err != nil {
			return nil, err
		}
		// TODO(jba,pongad): fix latencies
		// Later values will be skewed by earlier ones, since we wait for the
		// results in order. (On the other hand, it may not matter much, since
		// messages are added to bundles in order and bundles get sent more or
		// less in order.) If we want more accurate values, we can either start
		// a goroutine for each result (similar to the original code using a
		// callback), or call reflect.Select with the Ready channels of the
		// results.
		latencies[i] = time.Since(start).Nanoseconds() / 1e6
	}
	return latencies, nil
}

// SubServer is a dummy Pub/Sub server for load testing.
type SubServer struct {
	// TODO(deklerk): what is this actually for?
	lim *rate.Limiter

	mu        sync.Mutex
	idents    []*pb.MessageIdentifier
	latencies []int64
	pb.UnimplementedLoadtestWorkerServer
}

// Start starts the server.
func (s *SubServer) Start(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	log.Println("received start")
	s.lim = rate.NewLimiter(rate.Every(time.Second), 1)

	c, err := pubsub.NewClient(ctx, req.Project)
	if err != nil {
		return nil, err
	}

	// Load test API doesn't define any way to stop right now.
	go func() {
		sub := c.Subscription(req.GetPubsubOptions().Subscription)
		sub.ReceiveSettings.NumGoroutines = 10 * runtime.GOMAXPROCS(0)
		err := sub.Receive(context.Background(), s.callback)
		log.Fatal(err)
	}()

	log.Println("started")
	return &pb.StartResponse{}, nil
}

func (s *SubServer) callback(_ context.Context, m *pubsub.Message) {
	id, err := strconv.ParseInt(m.Attributes["clientId"], 10, 64)
	if err != nil {
		log.Println(err)
		m.Nack()
		return
	}

	seqNum, err := strconv.ParseInt(m.Attributes["sequenceNumber"], 10, 32)
	if err != nil {
		log.Println(err)
		m.Nack()
		return
	}

	sendTimeMillis, err := strconv.ParseInt(m.Attributes["sendTime"], 10, 64)
	if err != nil {
		log.Println(err)
		m.Nack()
		return
	}

	latency := time.Now().UnixNano()/1e6 - sendTimeMillis
	ident := &pb.MessageIdentifier{
		PublisherClientId: id,
		SequenceNumber:    int32(seqNum),
	}

	s.mu.Lock()
	s.idents = append(s.idents, ident)
	s.latencies = append(s.latencies, latency)
	s.mu.Unlock()
	m.Ack()
}
