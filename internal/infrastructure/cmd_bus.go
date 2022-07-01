package infrastructure

import (
	"fmt"

	"github.com/MAZEN-Kenjrawi/pwd/internal/application"
)

type CmdBus struct {
	addLoginHandler    application.AddLoginHandler
	removeLoginHandler application.RemoveLoginHandler
	signUpHandler      application.SignUpHandler
}

func (b *CmdBus) Handle(c interface{}) error {
	cmdType := getType(c)
	switch cmdType {
	case "application.AddLoginCommand":
		return b.addLoginHandler.Handle(c.(application.AddLoginCommand))
	case "application.RemoveLoginCommand":
		return b.removeLoginHandler.Handle(c.(application.RemoveLoginCommand))
	case "application.SignUpCommand":
		return b.signUpHandler.Handle(c.(application.SignUpCommand))
	}

	return fmt.Errorf("no handler found for command %s", cmdType)
}
