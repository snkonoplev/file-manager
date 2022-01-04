package mediator

import "context"

type Handler interface {
	Handle(context.Context, interface{}) (interface{}, error)
}
