package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type node struct {
	Name     string
	IsDir    bool
	Children []*node
	Path     string
}

func buildTree(root string, skip map[string]bool) (*node, error) {
	rootNode := &node{Name: filepath.Base(root), IsDir: true, Path: root}
	nodes := map[string]*node{root: rootNode}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == root {
			return nil
		}
		if d.IsDir() && skip[d.Name()] {
			return filepath.SkipDir
		}
		parent := filepath.Dir(path)
		parentNode := nodes[parent]
		if parentNode == nil {
			return nil
		}
		n := &node{Name: d.Name(), IsDir: d.IsDir(), Path: path}
		parentNode.Children = append(parentNode.Children, n)
		if d.IsDir() {
			nodes[path] = n
		}
		return nil
	})
	return rootNode, err
}

func printTree(n *node, indent string, lines *[]string) {
	prefix := "📁 "
	if !n.IsDir {
		prefix = "📄 "
	}
	*lines = append(*lines, indent+prefix+n.Name)

	sort.Slice(n.Children, func(i, j int) bool {
		if n.Children[i].IsDir != n.Children[j].IsDir {
			return n.Children[i].IsDir
		}
		return n.Children[i].Name < n.Children[j].Name
	})

	for _, c := range n.Children {
		printTree(c, indent+"  ", lines)
	}
}

// Helpers para “tempo real”
func countEntries(root string, skip map[string]bool, outFile string) (int, error) {
	total := 0
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && skip[d.Name()] {
			return filepath.SkipDir
		}
		// não contar o .md de saída
		if !d.IsDir() && filepath.Base(path) == outFile && filepath.Dir(path) == "." {
			return nil
		}
		if path == root {
			return nil
		}
		total++
		return nil
	})
	return total, err
}

func streamTree(root string, skip map[string]bool, emit func(line string)) error {
	type stackItem struct {
		path   string
		indent string
	}
	stack := []stackItem{{root, ""}}

	for len(stack) > 0 {
		item := stack[0]
		stack = stack[1:]

		entries, err := os.ReadDir(item.path)

		if err != nil {
			return err
		}

		// ordenar: diretórios primeiro
		sort.Slice(entries, func(i, j int) bool {
			if entries[i].IsDir() != entries[j].IsDir() {
				return entries[i].IsDir()
			}
			return entries[i].Name() < entries[j].Name()
		})

		// raiz imprime seu nome só uma vez
		if item.path == root {
			emit(item.indent + "📁 " + filepath.Base(root))
		}

		for _, e := range entries {
			if e.IsDir() && skip[e.Name()] {
				continue
			}
			line := item.indent + "  "
			if e.IsDir() {
				line += "📁 " + e.Name()
				emit(line)
				stack = append(stack, stackItem{filepath.Join(item.path, e.Name()), item.indent + "  "})
			} else {
				emit(line + "📄 " + e.Name())
			}
		}
	}
	return nil
}

func depthOf(path, root string) int {
	rel, _ := filepath.Rel(root, path)
	if rel == "." {
		return 0
	}
	return strings.Count(rel, string(filepath.Separator))
}
