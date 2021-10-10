package schedule

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// MakeScheduleEndpoint :
func MakeScheduleEndpoint(scheduleSvc ScheduleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ScheduleRequest)

		msg, err := scheduleSvc.ScheduleMicro(ctx, &req.employee)
		if err != nil {
			return nil, err
		}

		return ScheduleResponse{Msg: msg, Err: err}, nil
	}
}

type (
	// ScheduleRequest :
	ScheduleRequest struct {
		employee CreateScheduleRequest
	}
	// ScheduleResponse :
	ScheduleResponse struct {
		Msg string `json:"msg"`
		Err error  `json:"error,omitempty"`
	}
)

// decodeScheduleRequest ....
func DecodeScheduleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req ScheduleRequest

	if err := json.NewDecoder(r.Body).Decode(&req.employee); err != nil {
		return nil, err
	}

	return req, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}
