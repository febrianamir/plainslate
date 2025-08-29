package usecase

import (
	"errors"
	"io"
	"os"
	"path/filepath"
	"plainslate/backend/dto"
)

func (u *Usecase) checkBaseDirectory() error {
	if u.Config.RootPath == "" {
		return errors.New("base directory not set")
	}
	return nil
}

func (u *Usecase) buildFilePath(fileName string) string {
	return filepath.Join(u.Config.RootPath, fileName)
}

// openOrCreateFile open/create file if it doesn't exist.
func (u *Usecase) openOrCreateFile(filePath string) (*os.File, error) {
	err := u.checkBaseDirectory()
	if err != nil {
		return nil, err
	}

	if filePath == "" {
		return nil, errors.New("file path is required")
	}

	return os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0644)
}

func (u *Usecase) saveFile(filePath string, content string) error {
	err := u.checkBaseDirectory()
	if err != nil {
		return err
	}

	if filePath == "" {
		return errors.New("file path is required")
	}

	return os.WriteFile(filePath, []byte(content), 0644)
}

func (u *Usecase) OpenOrCreateFile(req dto.OpenOrCreateFileReq) (string, error) {
	file, err := u.openOrCreateFile(req.FilePath)
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

func (u *Usecase) SaveFile(req dto.SaveFileReq) error {
	return u.saveFile(req.FilePath, req.Content)
}
