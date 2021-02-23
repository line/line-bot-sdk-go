// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package linebot

import (
	"context"
	"net/url"
)

// GetFollowerIDs method
func (client *Client) GetFollowerIDs(continuationToken string) *GetFollowerIDsCall {
	return &GetFollowerIDsCall{
		c:                 client,
		continuationToken: continuationToken,
	}
}

// GetFollowerIDsCall type
type GetFollowerIDsCall struct {
	c   *Client
	ctx context.Context

	continuationToken string
}

// WithContext method
func (call *GetFollowerIDsCall) WithContext(ctx context.Context) *GetFollowerIDsCall {
	call.ctx = ctx
	return call
}

// Do method
func (call *GetFollowerIDsCall) Do() (*UserIDsResponse, error) {
	var q url.Values
	if call.continuationToken != "" {
		q = url.Values{"start": []string{call.continuationToken}}
	}
	res, err := call.c.get(call.ctx, call.c.endpointBase, APIEndpointGetFollowerIDs, q)
	if err != nil {
		return nil, err
	}
	defer closeResponse(res)
	return decodeToUserIDsResponse(res)
}

// NewScanner returns Group IDs scanner.
func (call *GetFollowerIDsCall) NewScanner() *UserIDsScanner {
	var ctx context.Context
	if call.ctx != nil {
		ctx = call.ctx
	} else {
		ctx = context.Background()
	}

	c2 := &GetFollowerIDsCall{}
	*c2 = *call
	c2.ctx = ctx

	return &UserIDsScanner{
		caller: c2,
		ctx:    ctx,
	}
}

func (call *GetFollowerIDsCall) setContinuationToken(token string) {
	call.continuationToken = token
}

type userIDsCaller interface {
	Do() (*UserIDsResponse, error)
	setContinuationToken(string)
}

// UserIDsScanner type
type UserIDsScanner struct {
	caller userIDsCaller
	ctx    context.Context
	start  int
	ids    []string
	next   string
	called bool
	done   bool
	err    error
}

// Scan method
func (s *UserIDsScanner) Scan() bool {
	if s.done {
		return false
	}

	select {
	case <-s.ctx.Done():
		s.err = s.ctx.Err()
		s.done = true
		return false
	default:
	}

	s.start++
	if len(s.ids) > 0 && len(s.ids) > s.start {
		return true
	}

	if s.next == "" && s.called {
		s.done = true
		return false
	}

	s.start = 0
	res, err := s.caller.Do()
	if err != nil {
		s.err = err
		s.done = true
		return false
	}

	s.called = true
	s.ids = res.UserIDs
	s.next = res.Next
	s.caller.setContinuationToken(s.next)

	return true
}

// ID returns member id.
func (s *UserIDsScanner) ID() string {
	if len(s.ids) == 0 {
		return ""
	}
	return s.ids[s.start : s.start+1][0]
}

// Err returns scan error.
func (s *UserIDsScanner) Err() error {
	return s.err
}
