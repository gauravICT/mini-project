package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// MakeConsumerEndpoint:
func MakeConsumerEndpoint(consumerSvc ConsumerService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ConsumerRequest)

		msg, err := consumerSvc.WriteMSG(ctx, &req.msg)
		if err != nil {
			return nil, err
		}

		return ScheduleResponse{Msg: msg, Err: err}, nil
	}
}

type (
	// ConsumerRequest :
	ConsumerRequest struct {
		msg CreateConsumerRequest
	}
	// ScheduleResponse :
	ScheduleResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
)

// decodeScheduleRequest ....
func DecodeConsumerRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req ConsumerRequest

	if err := json.NewDecoder(r.Body).Decode(&req.msg); err != nil {
		return nil, err
	}
	return req, nil
}

// EncodeResponse :
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}
