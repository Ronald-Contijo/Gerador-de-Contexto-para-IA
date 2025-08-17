package main

import (
	"os"
	"path/filepath"
)

func renderMarkdown(dir, outFile string, treeLines []string, codeFiles []codeFile, maxBytes int64) error {
	return renderMarkdownWithProgress(dir, outFile, treeLines, codeFiles, maxBytes, nil, nil, nil, nil)
}

// Versão com callbacks (para emitir progresso por arquivo)
func renderMarkdownWithProgress(
	dir, outFile string,
	treeLines []string,
	codeFiles []codeFile,
	maxBytes int64,
	onStart func(), onEnd func(),
	beforeFile func(), afterFile func(),
) error {
	if onStart != nil {
		onStart()
	}

	dirBase := filepath.Base(dir)
	f, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer f.Close()

	write := func(s string) { _, _ = f.WriteString(s) }

	write("# " + dirBase + "\n\n")
	write("## Árvore de arquivos (visão geral)\n\n```\n")
	for _, ln := range treeLines {
		write(ln + "\n")
	}
	write("```\n\n")

	write("## Conteúdo dos arquivos de código\n\n")
	write("> Apenas extensões reconhecidas têm conteúdo exibido; demais aparecem apenas na árvore acima.\n\n")

	for _, cf := range codeFiles {
		if beforeFile != nil {
			beforeFile()
		}

		toRead := cf.Size
		truncated := false
		if maxBytes > 0 && cf.Size > maxBytes {
			toRead = maxBytes
			truncated = true
		}

		var data []byte
		if toRead > 0 {
			file, err := os.Open(cf.Path)
			if err == nil {
				buf := make([]byte, toRead)
				n, _ := file.Read(buf)
				data = buf[:n]
				_ = file.Close()
			}
		}

		write("### `" + cf.Rel + "`\n\n")
		if cf.Lang == "md" {
			write("```\n")
		} else {
			write("```" + cf.Lang + "\n")
		}
		write(string(data))
		if truncated {
			write("\n… (conteúdo truncado)")
		}
		write("\n```\n\n")

		if afterFile != nil {
			afterFile()
		}
	}

	if onEnd != nil {
		onEnd()
	}
	return nil
}
