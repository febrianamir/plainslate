package model

type Match struct {
	FilePath string   `json:"file_path"`
	Lines    []string `json:"matches"`
}
