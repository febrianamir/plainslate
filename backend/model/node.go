package model

type Node struct {
	Name     string  `json:"name"`
	Path     string  `json:"path"`
	Type     string  `json:"type"` // Node Types
	Children []*Node `json:"children,omitempty"`
	IsRoot   bool    `json:"isRoot"`
}

// Node Types
const (
	NodeTypeFile      = "file"
	NodeTypeDirectory = "directory"
)
