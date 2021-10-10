package employee

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

// Employee :
type Employee struct {
	EmployeeID string `json:"employee_id" validate:"employee_id"`
	FirstName  string `json:"first_name" validate:"first_name"`
	LastName   string `json:"last_name" validate:"last_name"`
	Department string `json:"department" validate:"department"`
	Address    string `json:"address" validate:"address"`
	Email      string ` json:"email" validate:"email"`
}

// Repository :
type Repository interface {
	CreateEmployee(ctx context.Context, employee *Employee) (interface{}, error)
	GetEmployeeByID(ctx context.Context, id string) (interface{}, error)
	GetAllEmployee(ctx context.Context) (interface{}, error)
	UpdateEmployee(ctx context.Context, employee *Employee) (string, error)
	DeleteEmployee(ctx context.Context, id string) (string, error)
}

// Employeeservice :
type Employeeservice struct {
	repository Repository
	logger     log.Logger
}

// Service describes the Employee service.
type EmployeeService interface {
	CreateEmployee(ctx context.Context, employee Employee) (interface{}, error)
	GetEmployeeByID(ctx context.Context, id string) (interface{}, error)
	GetAllEmployee(ctx context.Context) (interface{}, error)
	UpdateEmployee(ctx context.Context, employee Employee) (string, error)
	DeleteEmployee(ctx context.Context, id string) (string, error)
}

// NewService creates and returns a new Account service instance
func NewService(rep Repository, logger log.Logger) EmployeeService {
	return &Employeeservice{
		repository: rep,
		logger:     logger,
	}
}

// CreateEmployee :
func (s Employeeservice) CreateEmployee(ctx context.Context, employee Employee) (interface{}, error) {
	logger := log.With(s.logger, "method", "Create")

	var msg = "success"
	//currentdatetime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))
	//CreatedAt := currentdatetime.Format("2006-01-02 15:04:05")
	employeeDetails := Employee{

		EmployeeID: employee.EmployeeID,
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		Department: employee.Department,
		Address:    employee.Address,
		Email:      employee.Email,
	}
	if employee, err := s.repository.CreateEmployee(ctx, &employeeDetails); err != nil {
		if err != nil {
			level.Error(logger).Log("err from repo is ", err)
			return "", err
		}

		return employee, err
	}
	return msg, nil
}

// GetEmployeeByID :
func (s Employeeservice) GetEmployeeByID(ctx context.Context, id string) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetEmployeeByID")

	var employee interface{}
	var empty interface{}
	employee, err := s.repository.GetEmployeeByID(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return employee, nil
}

// GetAllEmployee :
func (s Employeeservice) GetAllEmployee(ctx context.Context) (interface{}, error) {
	logger := log.With(s.logger, "method", "GetAllEmployee")
	var employee interface{}
	var empty interface{}
	employee, err := s.repository.GetAllEmployee(ctx)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return empty, err
	}
	return employee, nil
}

// UpdateEmployee :
func (s Employeeservice) UpdateEmployee(ctx context.Context, employee Employee) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	var msg = "success"
	employeeDetails := Employee{
		EmployeeID: employee.EmployeeID,
		FirstName:  employee.FirstName,
		LastName:   employee.LastName,
		Email:      employee.Email,
	}
	msg, err := s.repository.UpdateEmployee(ctx, &employeeDetails)
	if err != nil {
		level.Error(logger).Log("err from repo is ", err)
		return "", err
	}
	return msg, nil
}

// DeleteEmployee :
func (s Employeeservice) DeleteEmployee(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "DeleteEmployee")
	msg, err := s.repository.DeleteEmployee(ctx, id)
	if err != nil {
		level.Error(logger).Log("err ", err)
		return "", err
	}
	return msg, nil
}
