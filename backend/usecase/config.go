package usecase

import (
	"os"
	"plainslate/backend/lib"
)

func (u *Usecase) GetConfig() *lib.Config {
	return u.Config
}

// SetRootPath will set root path config & create the directory if not exists
func (u *Usecase) SetRootPath(rootPath string) error {
	err := os.MkdirAll(rootPath, 0755)
	if err != nil {
		return err
	}

	u.Config.RootPath = rootPath
	return nil
}
