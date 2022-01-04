package mediator

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Mediator struct {
	handlers map[reflect.Type]Handler
}

func NewMediator(handlers map[reflect.Type]Handler) *Mediator {
	return &Mediator{
		handlers: handlers,
	}
}

func (m *Mediator) Handle(c *gin.Context, command interface{}) (interface{}, error) {
	if handler, ok := m.handlers[reflect.TypeOf(command)]; ok {

		context := context.Background()
		var traceId interface{}

		if c != nil {
			context = c.Request.Context()
			traceId, _ = c.Get("traceId")
		}

		logger := logrus.StandardLogger()

		logger.WithContext(context).WithFields(logrus.Fields{
			"command":     command,
			"commantType": reflect.TypeOf(command),
			"traceId":     traceId,
		}).Info("start executing command...")

		response, err := handler.Handle(context, command)

		entry := logger.WithContext(context).WithFields(logrus.Fields{
			"traceId": traceId,
		})

		if err != nil {
			entry.WithError(err).Error("stop executing command...")
		} else {
			entry.Info("stop executing command...")
		}

		return response, err
	}
	return nil, &HandlerError{
		StatusCode: http.StatusInternalServerError,
		Message:    fmt.Sprintf("can't get handler for command/query type %v", reflect.TypeOf(command)),
	}
}
