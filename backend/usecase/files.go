package usecase

import (
	"errors"
	"os"
	"path/filepath"
)

func (u *Usecase) checkBaseDirectory() error {
	if u.Config.RootPath == "" {
		return errors.New("base directory not set")
	}
	return nil
}

func (u *Usecase) buildFilePath(filename string) string {
	return filepath.Join(u.Config.RootPath, filename)
}

// GetOrCreateFile opens the file for reading/writing and creates it if it doesn't exist.
func (u *Usecase) GetOrCreateFile(filename string) (*os.File, error) {
	err := u.checkBaseDirectory()
	if err != nil {
		return nil, err
	}

	if filename == "" {
		return nil, errors.New("filename is required")
	}

	return os.OpenFile(u.buildFilePath(filename), os.O_RDWR|os.O_CREATE, 0644)
}

func (u *Usecase) SaveFile(filename string, content string) error {
	err := u.checkBaseDirectory()
	if err != nil {
		return err
	}

	if filename == "" {
		return errors.New("filename is required")
	}

	return os.WriteFile(u.buildFilePath(filename), []byte(content), 0644)
}
