package main

import (
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	con "github.com/mini-project/consumer"
	e "github.com/mini-project/employee"
	s "github.com/mini-project/schedule"
)

func main() {
	// logger : for logging purposes
	logger := log.NewLogfmtLogger(os.Stderr)

	// GetDBconn : For connection to the database
	db := e.GetDBconn()

	r := mux.NewRouter()

	//  Services used in project
	var empSvc e.EmployeeService
	var schSvc s.ScheduleService
	var conSvc con.ConsumerService

	empSvc = &e.Employeeservice{}
	schSvc = &s.ScheduleSvc{}
	conSvc = &con.ConsumerSvc{}
	{
		repository, err := e.NewRepo(db, logger)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}
		empSvc = e.NewService(repository, logger)
	}

	// Handlers :
	CreateEmployeeHandler := httptransport.NewServer(
		e.MakeCreateEmployeeEndpoint(empSvc),
		e.DecodeCreateEmployeeRequest,
		e.EncodeResponse,
	)
	GetByEmployeeIDHandler := httptransport.NewServer(
		e.MakeGetEmployeeByIDEndpoint(empSvc),
		e.DecodeGetEmployeeByIDRequest,
		e.EncodeResponse,
	)
	GetAllEmployeeHandler := httptransport.NewServer(
		e.MakeGetAllEmployeeEndpoint(empSvc),
		e.DecodeGetAllEmployeeRequest,
		e.EncodeResponse,
	)
	DeleteEmployeeHandler := httptransport.NewServer(
		e.MakeDeleteEmployeeEndpoint(empSvc),
		e.DecodeDeleteEmployeeRequest,
		e.EncodeResponse,
	)
	UpdateEmployeeHandler := httptransport.NewServer(
		e.MakeUpdateEmployeeEndpoint(empSvc),
		e.DecodeUpdateEmployeeRequest,
		e.EncodeResponse,
	)
	ScheduleHandler := httptransport.NewServer(
		s.MakeScheduleEndpoint(schSvc),
		s.DecodeScheduleRequest,
		s.EncodeResponse,
	)
	ConsumerHandler := httptransport.NewServer(
		con.MakeConsumerEndpoint(conSvc),
		con.DecodeConsumerRequest,
		con.EncodeResponse,
	)
	http.Handle("/", r)
	r.Handle("/employee", CreateEmployeeHandler).Methods("POST")
	r.Handle("/employee/update", UpdateEmployeeHandler).Methods("PUT")
	r.Handle("/employee/getAll", GetAllEmployeeHandler).Methods("GET")
	r.Handle("/employee/{employee_id}", GetByEmployeeIDHandler).Methods("GET")
	r.Handle("/employee/{employee_id}", DeleteEmployeeHandler).Methods("DELETE")
	r.Handle("/schedule", ScheduleHandler).Methods("POST")
	r.Handle("/consumer", ConsumerHandler).Methods("GET")
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}
