package main

import (
	"flag"
	"log"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	dirFlag := flag.String("dir", ".", "Diretório raiz a processar")
	maxBytes := flag.Int64("max", 1<<20, "Tamanho máx por arquivo de código (bytes)")
	flag.Parse()

	dir := filepath.Clean(*dirFlag)
	outFile :=  "contexto.md"

	m := NewModel(dir, outFile, *maxBytes)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}
}
