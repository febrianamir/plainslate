package usecase

import (
	"os"
	"plainslate/backend/dto"
)

func (u *Usecase) CreateDirectory(req dto.CreateDirectoryReq) error {
	err := os.MkdirAll(req.Path, 0755)
	if err != nil {
		return err
	}

	return nil
}
