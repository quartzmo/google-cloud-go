// Copyright 2016 Google LLC
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

package testutil

import (
	"fmt"
	"testing"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func TestNewServer(t *testing.T) {
	srv, err := NewServer()
	if err != nil {
		t.Fatal(err)
	}
	srv.Start()

	conn, err := grpc.NewClient(srv.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		conn.Close()
		srv.Close()
	})
}

func TestNewServerWithAddress(t *testing.T) {
	addresses := []string{
		":8181",
		"0.0.0.0:8181",
		"127.0.0.1:8181",
		"localhost:8181",
	}

	for _, a := range addresses {
		t.Run(fmt.Sprintf("GIVEN host %s THEN succeed to init new server", a), func(t *testing.T) {
			srv, err := NewServerWithAddress(a)
			if err != nil {
				t.Fatal(err)
			}
			srv.Start()

			conn, err := grpc.NewClient(srv.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				t.Fatal(err)
			}

			t.Cleanup(func() {
				conn.Close()
				srv.Close()
			})
		})
	}
}

func TestPageBounds(t *testing.T) {
	const length = 10
	for _, test := range []struct {
		size     int
		tok      string
		wantFrom int
		wantTo   int
		wantTok  string
	}{
		{5, "",
			0, 5, "5"},
		{11, "",
			0, 10, ""},
		{5, "2",
			2, 7, "7"},
		{5, "8",
			8, 10, ""},
		{11, "8",
			8, 10, ""},
		{1, "11",
			10, 10, ""},
	} {
		gotFrom, gotTo, gotTok, err := PageBounds(test.size, test.tok, length)
		if err != nil {
			t.Fatal(err)
		}
		if got, want := gotFrom, test.wantFrom; got != want {
			t.Errorf("%+v: from: got %d, want %d", test, got, want)
		}
		if got, want := gotTo, test.wantTo; got != want {
			t.Errorf("%+v: to: got %d, want %d", test, got, want)
		}
		if got, want := gotTok, test.wantTok; got != want {
			t.Errorf("%+v: got %q, want %q", test, got, want)
		}
	}

	_, _, _, err := PageBounds(4, "xyz", 5)
	if status.Code(err) != codes.InvalidArgument {
		t.Errorf("want invalid argument, got <%v>", err)
	}
}
