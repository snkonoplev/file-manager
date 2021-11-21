package mediator

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
)

type Mediator struct {
	handlers map[reflect.Type]Handler
}

func NewMediator(handlers map[reflect.Type]Handler) *Mediator {
	return &Mediator{
		handlers: handlers,
	}
}

func (m *Mediator) Handle(context context.Context, command interface{}) (interface{}, error) {
	if handler, ok := m.handlers[reflect.TypeOf(command)]; ok {
		return handler.Handle(context, command)
	}
	return nil, &HandlerError{
		StatusCode: http.StatusInternalServerError,
		Message:    fmt.Sprintf("can't get handler for command/query type %v", reflect.TypeOf(command)),
	}
}
