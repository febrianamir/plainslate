package model

type Node struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Type     string  `json:"type"` // "file" or "directory"
	Children []*Node `json:"children,omitempty"`
	IsRoot   bool    `json:"is_root"`
}

const (
	NodeTypeFile      = "file"
	NodeTypeDirectory = "directory"
)
