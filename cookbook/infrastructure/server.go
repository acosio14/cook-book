package infrastructure

import "github.com/acosio14/cook-book/cookbook/domain"

type Server interface {
	ServeContent(*domain.Recipe) error
	ServeList([]*domain.Recipe) error
}
