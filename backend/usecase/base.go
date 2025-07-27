package usecase

import (
	"context"
	"plainslate/backend/lib"
)

type Usecase struct {
	Ctx    context.Context
	Config *lib.Config
}

func NewUsecase(c *lib.Config) *Usecase {
	return &Usecase{
		Config: c,
	}
}
