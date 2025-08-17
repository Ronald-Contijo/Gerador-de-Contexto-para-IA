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

// collectCodeFiles varre o diretório `dir`, coleta apenas os arquivos de código
// (com extensões mapeadas em langByExt) e retorna a lista de arquivos que devem
// ser renderizados no Markdown. Qualquer arquivo que esteja dentro de uma pasta
// listada em ignoreContentDirs é ignorado (aparece na árvore, mas não é lido).
func collectCodeFiles(
	dir, outFile string,
	langByExt map[string]string,
	ignoreContentDirs []string, // pastas cujo CONTEÚDO não será embutido
) ([]codeFile, error) {
	var list []codeFile

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// ignora erros de leitura de entrada
			return nil
		}
		if d.IsDir() {
			// não pula diretórios aqui para que eles apareçam na árvore;
			// a lógica de "não ler conteúdo" é feita por arquivo (abaixo).
			return nil
		}

		// não incluir o próprio markdown de saída na coleta de conteúdos (na raiz)
		if filepath.Base(path) == outFile && filepath.Dir(path) == "." {
			return nil
		}

		// caminho relativo em relação à raiz
		rel, _ := filepath.Rel(dir, path)

		// se o arquivo está dentro de alguma pasta “ignorar conteúdo”, pula
		if isUnderAny(rel, ignoreContentDirs) {
			return nil
		}

		// filtra por extensão reconhecida
		ext := strings.ToLower(filepath.Ext(d.Name()))
		if lang, ok := langByExt[ext]; ok {
			if info, statErr := os.Stat(path); statErr == nil && info != nil {
				list = append(list, codeFile{
					Path: path,
					Rel:  rel,
					Lang: lang,
					Size: info.Size(),
				})
			}
		}
		return nil
	})

	return list, err
}

// isUnderAny verifica se relPath (relativo a dir) está contido em alguma
// das pastas passadas em dirs. A comparação é por prefixo de caminho seguro.
func isUnderAny(relPath string, dirs []string) bool {
	if relPath == "" || len(dirs) == 0 {
		return false
	}
	rp := filepath.ToSlash(filepath.Clean(relPath))
	for _, d := range dirs {
		dd := filepath.ToSlash(filepath.Clean(d))
		// caso o próprio diretório seja apontado (arquivo direto dentro dele)
		if rp == dd {
			return true
		}
		// arquivo/subpasta dentro do diretório ignorado
		if strings.HasPrefix(rp, dd+"/") {
			return true
		}
	}
	return false
}
