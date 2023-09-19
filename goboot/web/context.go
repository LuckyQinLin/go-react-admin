package web

import (
	"goboot/domain/interfaces/engine"
	"sync"
)

type Pool sync.Pool

func (p *Pool) GetContext() engine.IContext {

}
