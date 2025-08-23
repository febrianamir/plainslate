package usecase

import (
	"bufio"
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"plainslate/backend/dto"
	"plainslate/backend/model"
	"slices"
	"strings"
	"sync"
)

func (u *Usecase) SearchInFiles(req dto.SearchInFilesReq) ([]model.Match, error) {
	// Cancel any ongoing search
	u.Searcher.Mu.Lock()
	if u.Searcher.Cancel != nil {
		u.Searcher.Cancel()
	}
	searchCtx, cancel := context.WithCancel(context.Background())
	u.Searcher.Cancel = cancel
	u.Searcher.Mu.Unlock()

	res := []model.Match{}
	var resMu sync.Mutex
	var wg sync.WaitGroup

	err := filepath.Walk(u.Config.RootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return nil // Skip file on error
		}
		if info.IsDir() {
			return nil // Skip directory
		}

		wg.Add(1)
		go func(path string) {
			defer wg.Done()

			f, err := os.Open(path)
			if err != nil {
				return
			}
			defer f.Close()

			// Skip non text file
			if !isTextFile(f) {
				return
			}
			f.Seek(0, io.SeekStart) // Reset file read offset after being used in isTextFile

			var matches []string
			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				select {
				case <-searchCtx.Done():
					return
				default:
				}
				line := scanner.Text()
				isMatch := strings.Contains(strings.ToLower(line), strings.ToLower(req.Query))
				if req.IsCaseSensitive {
					isMatch = strings.Contains(line, req.Query)
				}
				if isMatch {
					matches = append(matches, line)
				}
			}
			if len(matches) > 0 {
				resMu.Lock()
				res = append(res, model.Match{
					FilePath: path,
					Lines:    matches,
				})
				resMu.Unlock()
			}
		}(path)

		return nil
	})
	if err != nil {
		return res, err
	}

	wg.Wait()
	return res, nil
}

// isTextFile use binary check (null bytes in first KB)
func isTextFile(f *os.File) bool {
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	return !slices.Contains(buf[:n], 0)
}
