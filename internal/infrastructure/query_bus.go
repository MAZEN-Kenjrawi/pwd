package infrastructure

import (
	"fmt"

	"github.com/MAZEN-Kenjrawi/pwd/internal/application"
)

type QueryBus struct {
	getLoginHandler    application.GetLoginHandler
	searchLoginHandler application.SearchLoginHandler
}

func (b *QueryBus) Handle(q interface{}) (interface{}, error) {
	qType := getType(q)
	switch qType {
	case "application.GetLoginQuery":
		return b.getLoginHandler.Handle(q.(application.GetLoginQuery))
	case "application.SearchLoginQuery":
		return b.searchLoginHandler.Handle(q.(application.SearchLoginQuery))
	}

	return nil, fmt.Errorf("no handler found for query %s", qType)
}
