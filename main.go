package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// flags
	dirFlag := flag.String("d", ".", "Diretório raiz a processar")
	maxBytes := flag.Int64("m", 1<<20, "Tamanho máx por arquivo de código (bytes)")
	ignoreContent := flag.String("i", "", "Pastas (relativas) cujo conteúdo NÃO será lido (separadas por vírgula)")
	flag.Parse()

	// diretório alvo
	dir := filepath.Clean(*dirFlag)

	outFile := "contexto.md"

	// normaliza lista de pastas ignoradas por conteúdo
	var ignoreContentDirs []string
	if *ignoreContent != "" {
		for _, p := range strings.Split(*ignoreContent, ",") {
			p = strings.TrimSpace(p)
			if p != "" {
				ignoreContentDirs = append(ignoreContentDirs, filepath.Clean(p))
			}
		}
	}

	// inicializa modelo TUI
	m := NewModel(dir, outFile, *maxBytes)
	m.ignoreContentDirs = ignoreContentDirs

	// executa TUI
	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}
}
