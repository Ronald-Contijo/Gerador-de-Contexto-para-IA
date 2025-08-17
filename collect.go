package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type codeFile struct {
	Path string
	Rel  string
	Lang string
	Size int64
}

func collectCodeFiles(dir, outFile string, langByExt map[string]string) ([]codeFile, error) {
	var list []codeFile
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return nil
		}
		// não incluir o próprio markdown de saída na coleta de conteúdos
		if filepath.Base(path) == outFile && filepath.Dir(path) == "." {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if lang, ok := langByExt[ext]; ok {
			if info, err := os.Stat(path); err == nil && info != nil {
				rel, _ := filepath.Rel(dir, path)
				list = append(list, codeFile{Path: path, Rel: rel, Lang: lang, Size: info.Size()})
			}
		}
		return nil
	})
	return list, err
}
