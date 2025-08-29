package usecase

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"plainslate/backend/dto"
	"strconv"
	"strings"
	"time"
)

func (u *Usecase) RenamePath(req dto.RenamePathReq) error {
	err := os.Rename(req.OldPath, req.NewPath)
	if err != nil {
		return err
	}

	return nil
}

// MoveToTrash moves the file at the given path to the trash.
func (u *Usecase) MoveToTrash(req dto.MoveToTrashReq) error {
	// Find an open file name.
	tname := filepath.Base(req.Path)
	for i := 2; u.exists(tname); i++ {
		tname = filepath.Base(req.Path) + "." + strconv.Itoa(i)
	}

	// Write the trashinfo file.
	abs, err := filepath.Abs(req.Path)
	if err != nil {
		return err
	}
	infoString := u.trashInfo(abs, time.Now())
	err = os.WriteFile(u.info2path(tname), []byte(infoString), os.ModePerm)
	if err != nil {
		return err
	}

	// Next, move the file to the trash.
	return os.Rename(req.Path, u.file2path(tname))
}

func (u *Usecase) exists(tname string) bool {
	_, err := os.Stat(u.file2path(tname))
	return !os.IsNotExist(err)
}

func (u *Usecase) file2path(s string) string {
	trashDir := fmt.Sprintf("%s/.local/share/Trash", u.Config.UserHomeDir)
	return filepath.Join(trashDir, "files", s)
}

func (u *Usecase) info2path(tname string) string {
	trashDir := fmt.Sprintf("%s/.local/share/Trash", u.Config.UserHomeDir)
	return filepath.Join(trashDir, "info", tname+".trashinfo")
}

// trashInfo create trashinfo as a string in the INI format.
func (u *Usecase) trashInfo(path string, deletionDate time.Time) string {
	timeFormat := "2006-01-02T15:04:05"

	return fmt.Sprintf(
		"[Trash Info]\nPath=%v\nDeletionDate=%v\n",
		queryEscape(path),
		deletionDate.Format(timeFormat),
	)
}

// queryEscape is a wrapper function around url.QueryEscape that doesn't
// escape '/'.
func queryEscape(s string) string {
	// The first for loop is the workaround for "/".
	a := strings.Split(s, "/")
	for i := range a {
		// The second for loop is the workaround for " ".
		b := strings.Split(a[i], " ")
		for j := range b {
			b[j] = url.QueryEscape(b[j])
		}
		a[i] = strings.Join(b, "%20")
	}
	return strings.Join(a, "/")
}
