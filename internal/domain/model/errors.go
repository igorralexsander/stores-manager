package model

import "fmt"

const (
	NotFound      = "NOT_FOUND"
	AlreadyExists = "ALREADY_EXISTS"
)

type DomainError struct {
	Code    string      `json:"code"`
	Id      interface{} `json:"id"`
	Message string      `json:"message"`
}

func NotFoundError(id any) error {
	return DomainError{
		Code:    NotFound,
		Id:      id,
		Message: fmt.Sprintf("Record with %v not found", id),
	}
}

func AlreadyExistsError(id any) error {
	return DomainError{
		Code:    AlreadyExists,
		Id:      id,
		Message: fmt.Sprintf("Record with %v already exists", id),
	}
}

func (e DomainError) Error() string {
	return e.Message
}
