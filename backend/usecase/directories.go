package usecase

import "os"

func (u *Usecase) CreateDirectory(path string) error {
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	return nil
}
