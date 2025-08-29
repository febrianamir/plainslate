package usecase

import (
	"os"
	"plainslate/backend/dto"
	"plainslate/backend/lib"
)

func (u *Usecase) GetConfig() *lib.Config {
	return u.Config
}

// SetRootPath will set root path config & create the directory if not exists
func (u *Usecase) SetRootPath(req dto.SetRootPathReq) error {
	err := os.MkdirAll(req.RootPath, 0755)
	if err != nil {
		return err
	}

	u.Config.RootPath = req.RootPath
	return nil
}
