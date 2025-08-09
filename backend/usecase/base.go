package usecase

import (
	"context"
	"plainslate/backend/lib"
)

type Usecase struct {
	Ctx      context.Context
	Config   *lib.Config
	Searcher *lib.Searcher
}

func NewUsecase(c *lib.Config, s *lib.Searcher) *Usecase {
	return &Usecase{
		Config:   c,
		Searcher: s,
	}
}
