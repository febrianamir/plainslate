package lib

import (
	"context"
	"sync"
)

type Searcher struct {
	Mu     sync.Mutex
	Cancel context.CancelFunc
}

func NewSearcher() *Searcher {
	return &Searcher{}
}
