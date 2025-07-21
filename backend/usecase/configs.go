package usecase

import "os"

// SetRootPath will set root path config & create the directory if not exists
func (u *Usecase) SetRootPath(rootPath string) error {
	err := os.MkdirAll(rootPath, 0755)
	if err != nil {
		return err
	}

	u.Config.RootPath = rootPath
	return nil
}
