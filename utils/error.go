package utils

import (
	"fmt"
)

type InvalidParamsError struct {
	Message string
}

func NewInvalidParamsError(message string) *InvalidParamsError {
	return &InvalidParamsError{Message: message}
}

func (i *InvalidParamsError) Error() string {
	return fmt.Sprintf("Invalid params: %s", i.Message)
}

type DbError struct {
	Message string
}

func NewDbError(message string) *DbError {
	return &DbError{Message: message}
}

func (d *DbError) Error() string {
	return fmt.Sprintf("Db error: %s", d.Message)
}

type DbRowNotFound struct {
	Message string
}

func NewDbRowNotFound(message string) *DbRowNotFound {
	return &DbRowNotFound{Message: message}
}

func (d *DbRowNotFound) Error() string {
	return fmt.Sprintf("Db Row Not Found error: %s", d.Message)
}

type Status struct {
	Code    int64
	Message string
}

func GetStatus(err error) *Status {
	if err == nil {
		return &Status{
			Code:    0,
			Message: "ok",
		}
	}
	switch err.(type) {
	case *InvalidParamsError:
		return &Status{
			Code:    400,
			Message: err.Error(),
		}
	case *DbError:
		return &Status{
			Code:    501,
			Message: err.Error(),
		}
	case *DbRowNotFound:
		return &Status{
			Code:    404,
			Message: err.Error(),
		}
	default:
		return &Status{
			Code:    500,
			Message: err.Error(),
		}
	}
}
