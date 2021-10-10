package consumer

import (
	"context"
	"fmt"
	"log"

	p "github.com/mini-project/pulsar"
)

type (
	// ConsumerService ...
	ConsumerService interface {
		WriteMSG(ctx context.Context, payload *CreateConsumerRequest) (string, error)
	}

	ConsumerSvc struct {
		logger log.Logger
	}

	// CreateConsumerRequest :
	CreateConsumerRequest struct {
		Topic string `json:"topic"`
		//URL   string `json:url`
	}
)

// NewConsumerService :
func NewConsumerService(logger log.Logger) ConsumerService {

	return &ConsumerSvc{
		logger: logger,
	}
}

// WriteMSG :
func (ss *ConsumerSvc) WriteMSG(ctx context.Context, payload *CreateConsumerRequest) (string, error) {

	err := p.PulsarConsumer(ctx, payload.Topic)
	if err != nil {
		fmt.Print("Failed to consume the msg")
	}

	return "Successfully Completed The Task", nil
}
