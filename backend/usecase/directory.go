package usecase

import (
	"os"
	"path/filepath"
	"plainslate/backend/dto"
)

func (u *Usecase) CreateDirectory(req dto.CreateDirectoryReq) error {
	err := os.MkdirAll(req.Path, 0755)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) CopyDirectory(req dto.CopyDirectoryReq) error {
	return filepath.Walk(req.SourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(req.SourcePath, path)
		if err != nil {
			return err
		}

		targetPath := filepath.Join(req.DestPath, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}
		return u.CopyFile(dto.CopyFileReq{
			SourcePath: path,
			DestPath:   targetPath,
		})
	})
}
