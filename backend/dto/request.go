package dto

// Config
type (
	SetRootPathReq struct {
		RootPath string `json:"rootPath"`
	}
)

// Directory
type (
	CreateDirectoryReq struct {
		Path string `json:"path"`
	}

	CopyDirectoryReq struct {
		SourcePath string `json:"sourcePath"`
		DestPath   string `json:"destPath"`
	}
)

// File
type (
	OpenOrCreateFileReq struct {
		FilePath string `json:"filePath"`
	}

	SaveFileReq struct {
		FilePath string `json:"filePath"`
		Content  string `json:"content"`
	}

	CopyFileReq struct {
		SourcePath string `json:"sourcePath"`
		DestPath   string `json:"destPath"`
	}
)

// Log
type (
	WriteLogReq struct {
		Level     string         `json:"level"`
		LogFields map[string]any `json:"logFields"`
	}
)

// Path
type (
	RenamePathReq struct {
		OldPath string `json:"oldPath"`
		NewPath string `json:"newPath"`
	}

	MoveToTrashReq struct {
		Path string `json:"path"`
	}

	CheckPathReq struct {
		Path string `json:"path"`
	}
)

// Search
type (
	SearchInFilesReq struct {
		Query           string `json:"query"`
		IsCaseSensitive bool   `json:"isCaseSensitive"`
	}
)

// Node
type (
	GetNodeTreeReq struct {
		Path string `json:"path"`
	}
)
