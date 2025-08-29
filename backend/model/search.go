package model

type Match struct {
	FilePath string   `json:"filePath"`
	Lines    []string `json:"matches"`
}
