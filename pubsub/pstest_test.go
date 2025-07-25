// Copyright 2018 Google LLC
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

package pubsub_test

import (
	"context"
	"errors"
	"strconv"
	"sync"
	"testing"

	"cloud.google.com/go/internal/testutil"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/pubsub/pstest"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// withGRPCHeadersAssertionAlt is named differently than
// withGRPCHeadersAssertion in integration_test.go, because integration_test.go
// doesn't perform an external test i.e. its package name is "pubsub" while
// this one's is "pubsub_test", and when using Go Modules, without this rename
// go test won't find the function "withGRPCHeadersAssertion".
func withGRPCHeadersAssertionAlt(t *testing.T, opts ...option.ClientOption) []option.ClientOption {
	grpcHeadersEnforcer := &testutil.HeadersEnforcer{
		OnFailure: t.Fatalf,
		Checkers: []*testutil.HeaderChecker{
			testutil.XGoogClientHeaderChecker,
		},
	}
	return append(grpcHeadersEnforcer.CallOptions(), opts...)
}

func TestPSTest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	srv := pstest.NewServer()
	defer srv.Close()

	conn, err := grpc.Dial(srv.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	opts := withGRPCHeadersAssertionAlt(t, option.WithGRPCConn(conn))
	client, err := pubsub.NewClient(ctx, "some-project", opts...)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	topic, err := client.CreateTopic(ctx, "test-topic")
	if err != nil {
		panic(err)
	}

	sub, err := client.CreateSubscription(ctx, "sub-name", pubsub.SubscriptionConfig{Topic: topic})
	if err != nil {
		panic(err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			srv.Publish("projects/some-project/topics/test-topic", []byte(strconv.Itoa(i)), nil)
		}
	}()

	ctx, cancel := context.WithCancel(ctx)
	var mu sync.Mutex
	count := 0
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		mu.Lock()
		count++
		if count >= 10 {
			cancel()
		}
		mu.Unlock()
		m.Ack()
	})
	if err != nil && !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
