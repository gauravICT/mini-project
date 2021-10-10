package schedule

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/mini-project/pulsar"
)

type (
	// ScheduleService ...
	ScheduleService interface {
		ScheduleMicro(ctx context.Context, payload *CreateScheduleRequest) (string, error)
	}

	// ScheduleSvc :
	ScheduleSvc struct {
		logger log.Logger
	}

	// CreateScheduleRequest :
	CreateScheduleRequest struct {
		ReqType    string `json:"req_type"`
		URL        string `json:"url"`
		EmployeeID string `json:"employee_id"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Department string `json:"department"`
		Address    string `json:"address"`
		Email      string ` json:"email"`
	}

	// SchedulePayloadRequest :
	SchedulePayloadRequest struct {
		EmployeeID string `json:"employee_id"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Department string `json:"department"`
		Address    string `json:"address"`
		Email      string ` json:"email"`
	}
)

// NewScheduleService :
func NewScheduleService(logger log.Logger) ScheduleService {

	return &ScheduleSvc{
		logger: logger,
	}
}

// ScheduleMicro :
func (ss *ScheduleSvc) ScheduleMicro(ctx context.Context, payload *CreateScheduleRequest) (string, error) {

	// Schedular
	s := gocron.NewScheduler(time.Local)
	_, _ = s.Every(4).Second().LimitRunsTo(1).Do(CallMicro(ctx, payload))

	s.StartAsync()
	time.Sleep(time.Second * 7)

	return "Successfully Completed The Task", nil
}

// CallMicro :
func CallMicro(ctx context.Context, payload *CreateScheduleRequest) (string, error) {
	var typeAPI string
	var payload1 SchedulePayloadRequest
	if payload.ReqType == "POST" {
		typeAPI = "http.MethodPost"

	} else if payload.ReqType == "PUT" {
		typeAPI = "http.MethodPut"
	} else if payload.ReqType == "DELETE" {
		typeAPI = "http.MethodDelete"
	} else if payload.ReqType == "GET" {
		typeAPI = "http.MethodGet"
	}

	payload1.EmployeeID = payload.EmployeeID
	payload1.FirstName = payload.FirstName
	payload1.LastName = payload.LastName
	payload1.Email = payload.Email

	jsonReq, err := json.Marshal(payload1)
	if err != nil {
		fmt.Print("Unable to marshal json")
	}

	req, err := http.NewRequest(typeAPI, payload.URL, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)

	// Publishing msg into topic
	err = pulsar.PulsarProducer(ctx, bodyString)
	if err != nil {
		fmt.Print("Producer failed to produce the message")
	}
	fmt.Println(bodyString)
	return bodyString, nil
}
