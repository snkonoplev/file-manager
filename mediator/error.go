package mediator

import "fmt"

type HandlerError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *HandlerError) Error() string {
	return fmt.Sprintf("status %d: message: %s err %v", e.StatusCode, e.Message, e.Err)
}
