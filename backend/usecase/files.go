package usecase

import (
	"errors"
	"io"
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

// openFile open/create file it if it doesn't exist.
func (u *Usecase) openFile(filepath string) (*os.File, error) {
	err := u.checkBaseDirectory()
	if err != nil {
		return nil, err
	}

	if filepath == "" {
		return nil, errors.New("filepath is required")
	}

	return os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
}

func (u *Usecase) saveFile(filepath string, content string) error {
	err := u.checkBaseDirectory()
	if err != nil {
		return err
	}

	if filepath == "" {
		return errors.New("filepath is required")
	}

	return os.WriteFile(filepath, []byte(content), 0644)
}

func (u *Usecase) OpenFile(filepath string) (string, error) {
	file, err := u.openFile(filepath)
	if err != nil {
		return "", err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	content := string(data)
	return content, err
}

func (u *Usecase) SaveFile(filepath string, content string) error {
	return u.saveFile(filepath, content)
}
