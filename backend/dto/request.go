package dto

type SearchInFilesReq struct {
	Query           string `json:"query"`
	IsCaseSensitive bool   `json:"is_case_sensitive"`
}
