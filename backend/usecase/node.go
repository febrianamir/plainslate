package usecase

import (
	"io/fs"
	"os"
	"path/filepath"
	"plainslate/backend/dto"
	"plainslate/backend/model"
	"strings"
)

func (u *Usecase) buildTree(root string) (*model.Node, error) {
	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	rootNode := &model.Node{
		Name:     info.Name(),
		Path:     root,
		Type:     "directory",
		Children: []*model.Node{},
		IsRoot:   true,
	}

	err = filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == root {
			return nil // skip root
		}

		rel, _ := filepath.Rel(root, path)
		parts := strings.Split(rel, string(os.PathSeparator))
		u.insert(rootNode, parts, path, entry)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return rootNode, nil
}

func (u *Usecase) insert(parent *model.Node, parts []string, fullPath string, entry fs.DirEntry) {
	if len(parts) == 0 {
		return
	}

	currentName := parts[0]
	// Find or create child
	var child *model.Node
	for _, c := range parent.Children {
		if c.Name == currentName {
			child = c
			break
		}
	}

	if child == nil {
		child = &model.Node{
			Name:     currentName,
			Path:     fullPath,
			Type:     "file",
			Children: []*model.Node{},
		}
		if entry.IsDir() {
			child.Type = "directory"
		}
		parent.Children = append(parent.Children, child)
	}

	// Recurse into next level
	if len(parts) > 1 {
		u.insert(child, parts[1:], fullPath, entry)
	}
}

func (u *Usecase) GetRootNodeTree() (*model.Node, error) {
	err := u.checkBaseDirectory()
	if err != nil {
		return nil, err
	}

	nodeTree, err := u.buildTree(u.Config.RootPath)
	if err != nil {
		return nil, err
	}

	return nodeTree, err
}

func (u *Usecase) GetNodeTree(req dto.GetNodeTreeReq) (*model.Node, error) {
	nodeTree, err := u.buildTree(req.Path)
	if err != nil {
		return nil, err
	}

	return nodeTree, err
}
