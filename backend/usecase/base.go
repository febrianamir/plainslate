package usecase

import "plainslate/backend/lib"

type Usecase struct {
	Config *lib.Config
}

func NewUsecase(c *lib.Config) *Usecase {
	return &Usecase{
		Config: c,
	}
}
