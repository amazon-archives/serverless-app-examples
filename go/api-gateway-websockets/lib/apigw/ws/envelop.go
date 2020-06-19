// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// AWS Samples / Serverless Application Examples / Go WebSockets
//
// Go source code, Infrastructure as Code templates, build scripts, and configuration
// files for deploying a minimal example demonstrating how to use WebSockets with Amazon
// API Gateway, AWS Lambda, Amazon ElastiCache for Redis, and Amazon VPC.

// Package ws provides common resources for working with Amazon API Gateway WebSockets
package ws

import "encoding/json"

// InputEnvelop defines the expected structure for incoming messages sent over the WebSocket connection. The envelop
// provides additional metadata in addition to the message data.
type InputEnvelop struct {
	Echo bool            `json:"echo"`
	Type int             `json:"type"`
	Data json.RawMessage `json:"data"`
}

// Decode decodes and populates the InputEnvelop from the provided bytes.
func (e *InputEnvelop) Decode(data []byte) (*InputEnvelop, error) {
	err := json.Unmarshal(data, e)
	return e, err
}

// OutputEnvelop defines the structure for messages sent over the WebSocket connection from the backend service. The
// envelop provides additional metadata in addition to the message data.
type OutputEnvelop struct {
	Type     int             `json:"type"`
	Data     json.RawMessage `json:"data"`
	Received int64           `json:"received"`
}

// Encode encodes the OutputEnvelop as JSON. The output is suitable for sending over the wire.
func (e *OutputEnvelop) Encode() ([]byte, error) {
	return json.Marshal(e)
}
