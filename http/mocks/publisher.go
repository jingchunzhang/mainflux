//
// Copyright (c) 2018
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package mocks

import "github.com/mainflux/mainflux"

var _ (mainflux.MessagePublisher) = (*mockPublisher)(nil)

type mockPublisher struct{}

// NewPublisher returns mock message publisher.
func NewPublisher() mainflux.MessagePublisher {
	return mockPublisher{}
}

func (pub mockPublisher) Publish(msg mainflux.RawMessage) error {
	return nil
}
