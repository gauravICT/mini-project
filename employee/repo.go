package employee

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-kit/log"
)

var (
	//repoErr       = errors.New("Unable to handle Repo Request")
	errIdNotFound = errors.New("id not found")
)

// repo :
type repo struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepo :
func NewRepo(db *sql.DB, logger log.Logger) (Repository, error) {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "postgresql"),
	}, nil
}

// CreateEmployee :
func (repo *repo) CreateEmployee(ctx context.Context, employee *Employee) (interface{}, error) {
	currentdatetime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))
	createdAt := currentdatetime.Format("2006-01-02 15:04:05")
	_, err := repo.db.ExecContext(ctx, "INSERT INTO employee_data (employee_id, first_name, last_name, department, address, email, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)", employee.EmployeeID, employee.FirstName, employee.LastName, employee.Department, employee.Address, employee.Email, createdAt)
	if err != nil {
		fmt.Println("Error occured inside CreateEmployee in repo")
		return "", err
	} else {
		fmt.Print("User Created")
	}
	return employee, nil
}

// GetEmployeeByID :
func (repo *repo) GetEmployeeByID(ctx context.Context, id string) (interface{}, error) {
	employee := Employee{}

	err := repo.db.QueryRowContext(ctx, "SELECT employee_id,first_name,last_name,department,address,email FROM employee_data  where employee_id = $1", id).Scan(&employee.EmployeeID, &employee.FirstName, &employee.LastName, &employee.Department, &employee.Address, &employee.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return employee, errIdNotFound
		}
		return employee, err
	}
	return employee, nil
}

// GetAllEmployee :
func (repo *repo) GetAllEmployee(ctx context.Context) (interface{}, error) {
	employee := Employee{}
	var res []interface{}
	rows, err := repo.db.QueryContext(ctx, "SELECT employee_id,email,first_name,last_name, department, address, email address FROM employee_data")
	if err != nil {
		if err == sql.ErrNoRows {
			return employee, errIdNotFound
		}
		return employee, err
	}

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&employee.EmployeeID, &employee.Email, &employee.FirstName, &employee.LastName, &employee.Department, &employee.Address, &employee.Email)
		if err != nil {
			fmt.Print(err)
		}
		res = append([]interface{}{employee}, res...)
	}
	return res, nil
}

// UpdateEmployee :
func (repo *repo) UpdateEmployee(ctx context.Context, employee *Employee) (string, error) {
	currentdatetime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))
	updatedAt := currentdatetime.Format("2006-01-02 15:04:05")
	res, err := repo.db.ExecContext(ctx, "UPDATE employee_data  SET email = $1, department = $2, address =$3, updated_at=$4  WHERE employee_id = $5", employee.Email, employee.Department, employee.Address, updatedAt, employee.EmployeeID)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowCnt == 0 {
		return "", errIdNotFound
	}
	str := "Successfully updated employee_id : " + employee.EmployeeID + " whose email : " + employee.Email + " department : " + employee.Department + "address : " + employee.Address + "has been updated at : " + updatedAt
	return str, err
}

// DeleteEmployee :
func (repo *repo) DeleteEmployee(ctx context.Context, id string) (string, error) {
	currentdatetime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02 15:04:05"))
	deleatedAt := currentdatetime.Format("2006-01-02 15:04:05")
	res, err := repo.db.ExecContext(ctx, "DELETE FROM employee_data WHERE employee_id = $1 ", id)
	if err != nil {
		return "", err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return "", err
	} else if rowCnt == 0 {
		return "", errIdNotFound
	}
	str := "Successfully deleted " + "employee_id : " + id + "at : " + deleatedAt
	return str, nil
}
