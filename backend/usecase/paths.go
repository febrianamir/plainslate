package usecase

import "os"

func (u *Usecase) RenamePath(oldPath, newPath string) error {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}

	return nil
}
