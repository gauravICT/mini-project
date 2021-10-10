package employee

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

// Endpoint for the Employee service.

// MakeCreateEmployeeEndpoint :
func MakeCreateEmployeeEndpoint(s EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateEmployeeRequest)
		msg, err := s.CreateEmployee(ctx, req.employee)
		return CreateEmployeeResponse{Employee: msg, Err: err}, nil
	}
}

// MakeGetEmployeeByIDEndpoint :
func MakeGetEmployeeByIDEndpoint(s EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetEmployeeByIDRequest)
		employeeDetails, err := s.GetEmployeeByID(ctx, req.Id)
		if err != nil {
			return GetEmployeeByIDResponse{Employee: employeeDetails, Err: "Id not found"}, nil
		}
		return GetEmployeeByIDResponse{Employee: employeeDetails, Err: ""}, nil
	}
}

// MakeGetAllEmployeeEndpoint :
func MakeGetAllEmployeeEndpoint(s EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		employeeDetails, err := s.GetAllEmployee(ctx)
		if err != nil {
			return GetAllEmployeeResponse{Employee: employeeDetails, Err: "no data found"}, nil
		}
		return GetAllEmployeeResponse{Employee: employeeDetails, Err: ""}, nil
	}
}

// MakeDeleteEmployeeEndpoint :
func MakeDeleteEmployeeEndpoint(s EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEmployeeRequest)
		msg, err := s.DeleteEmployee(ctx, req.EmployeeID)
		if err != nil {
			return DeleteEmployeeResponse{Msg: msg, Err: err}, nil
		}
		return DeleteEmployeeResponse{Msg: msg, Err: nil}, nil
	}
}

// MakeUpdateEmployeeEndpoint :
func MakeUpdateEmployeeEndpoint(s EmployeeService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEmployeeRequest)
		msg, err := s.UpdateEmployee(ctx, req.employee)
		return msg, err
	}
}

// DecodeCreateEmployeeRequest :
func DecodeCreateEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req CreateEmployeeRequest
	fmt.Println("-------->>>>into Decoding")
	if err := json.NewDecoder(r.Body).Decode(&req.employee); err != nil {
		return nil, err
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	return req, nil
}

// DecodeGetEmployeeIDRequest :
func DecodeGetEmployeeByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetEmployeeByIDRequest
	fmt.Println("-------->>>>into GetById Decoding")
	vars := mux.Vars(r)
	req = GetEmployeeByIDRequest{
		Id: vars["employee_id"],
	}
	return req, nil
}

// DecodeGetAllEmployeeRequest :
func DecodeGetAllEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into GETALL Decoding")
	var req GetAllEmployeeRequest
	return req, nil
}

// DecodeDeleteEmployeeRequest :
func DecodeDeleteEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Delete Decoding")
	var req DeleteEmployeeRequest
	vars := mux.Vars(r)
	req = DeleteEmployeeRequest{
		EmployeeID: vars["employee_id"],
	}
	return req, nil
}

// DecodeUpdateEmployeeRequest :
func DecodeUpdateEmployeeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	fmt.Println("-------->>>> Into Update Decoding")
	var req UpdateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req.employee); err != nil {
		return nil, err
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	return req, nil
}

// EncodeResponse
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Println("into Encoding <<<<<<----------------")
	return json.NewEncoder(w).Encode(response)
}

type (
	// CreateEmployeeRequest :
	CreateEmployeeRequest struct {
		employee Employee
	}

	// CreateEmployeeResponse :
	CreateEmployeeResponse struct {
		//Msg string `json:"msg"`
		Employee interface{} `json:"employee,omitempty"`
		Err      error       `json:"error,omitempty"`
	}

	// GetEmployeeByIDRequest :
	GetEmployeeByIDRequest struct {
		Id string `json:"employee_id"`
	}

	// GetEmployeeByIDResponse :
	GetEmployeeByIDResponse struct {
		Employee interface{} `json:"employee,omitempty"`
		Err      string      `json:"error,omitempty"`
	}

	// GetAllEmployeesRequest :
	GetAllEmployeeRequest struct{}

	// GetAllEmployeeResponse :
	GetAllEmployeeResponse struct {
		Employee interface{} `json:"employee,omitempty"`
		Err      string      `json:"error,omitempty"`
	}

	// DeleteEmployeeRequest :
	DeleteEmployeeRequest struct {
		EmployeeID string `json:"employee_id"`
	}

	// DeleteEmployeeRequest :
	DeleteEmployeeResponse struct {
		Msg string `json:"response"`
		Err error  `json:"error,omitempty"`
	}

	// UpdateEmployeeRequest :
	UpdateEmployeeRequest struct {
		employee Employee
	}

	// DeleteEmployeeResponse :
	UpdateEmployeeResponse struct {
		Msg string `json:"status,omitempty"`
		Err error  `json:"error,omitempty"`
	}
)
