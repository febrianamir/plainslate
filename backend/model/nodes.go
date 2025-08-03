package model

type Node struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Type     string  `json:"type"` // Node Type
	Children []*Node `json:"children,omitempty"`
	IsRoot   bool    `json:"is_root"`
}

// Node Type
const (
	NodeTypeFile      = "file"
	NodeTypeDirectory = "directory"
)
