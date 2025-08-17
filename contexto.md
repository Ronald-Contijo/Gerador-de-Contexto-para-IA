# .

## Árvore de arquivos (visão geral)

```
📁 .
  📁 teste
    📄 main.py
  📄 .gitattributes
  📄 README.md
  📄 collect.go
  📄 contexto.md
  📄 filters.go
  📄 go.mod
  📄 go.sum
  📄 lista_de_arquivos.md
  📄 main
  📄 main.go
  📄 render.go
  📄 tree.go
  📄 tui.go
```

## Conteúdo dos arquivos de código

> Apenas extensões reconhecidas têm conteúdo exibido; demais aparecem apenas na árvore acima.

### `README.md`

```
Este projeto é uma pequena ferramenta em Go com interface em terminal (TUI) construída sobre a biblioteca Bubble Tea 
O objetivo é gerar automaticamente documentação de um diretório de código, exibindo:

- Árvore de arquivos e diretórios.
- Conteúdo dos arquivos reconhecidos (códigos, scripts, etc.).
- Arquivo final em formato Markdown.

A idéia é mandar de contexto pra alguma IA, como o ChatGePeto, para que ele tenha um contexto


```

### `collect.go`

```go
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

```

### `filters.go`

```go
package main

func defaultLangByExt() map[string]string {
	return map[string]string{
		".go": "go", ".html": "html", ".htm": "html", ".css": "css",
		".js": "javascript", ".mjs": "javascript", ".cjs": "javascript",
		".ts": "typescript", ".tsx": "tsx", ".jsx": "jsx",
		".py": "python", ".rb": "ruby", ".php": "php",
		".java": "java", ".c": "c", ".h": "c",
		".hpp": "cpp", ".cpp": "cpp", ".cc": "cpp",
		".cs": "csharp", ".rs": "rust", ".kt": "kotlin",
		".swift": "swift", ".sh": "bash", ".bat": "bat", ".ps1": "powershell",
		".sql": "sql", ".json": "json", ".yaml": "yaml", ".yml": "yaml",
		".toml": "toml", ".ini": "ini", ".md": "md", ".tex": "latex",
		".r": "r", ".m": "matlab",
	}
}

func defaultSkipDirs() map[string]bool {
	return map[string]bool{
		".git": true, "node_modules": true, "dist": true, "build": true,
		"out": true, "venv": true, ".venv": true, ".idea": true, ".vscode": true,"":true,
	}
}

```

### `lista_de_arquivos.md`

```
# .

## Árvore de arquivos (visão geral)

```
📁 .
  📁 teste
    📄 main.py
  📄 ..md
  📄 .gitattributes
  📄 collect.go
  📄 filters.go
  📄 go.mod
  📄 go.sum
  📄 main
  📄 main.go
  📄 render.go
  📄 tree.go
  📄 tui.go
```

## Conteúdo dos arquivos de código

> Apenas extensões reconhecidas têm conteúdo exibido; demais aparecem apenas na árvore acima.

### `..md`

```
# .

## Árvore de arquivos (visão geral)

```
📁 .
  📁 teste
    📄 main.py
  📄 ..md
  📄 .gitattributes
  📄 collect.go
  📄 filters.go
  📄 go.mod
  📄 go.sum
  📄 main
  📄 main.go
  📄 render.go
  📄 tree.go
  📄 tui.go
```

## Conteúdo dos arquivos de código

> Apenas extensões reconhecidas têm conteúdo exibido; demais aparecem apenas na árvore acima.

### `collect.go`

```go
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

```

### `filters.go`

```go
package main

func defaultLangByExt() map[string]string {
	return map[string]string{
		".go": "go", ".html": "html", ".htm": "html", ".css": "css",
		".js": "javascript", ".mjs": "javascript", ".cjs": "javascript",
		".ts": "typescript", ".tsx": "tsx", ".jsx": "jsx",
		".py": "python", ".rb": "ruby", ".php": "php",
		".java": "java", ".c": "c", ".h": "c",
		".hpp": "cpp", ".cpp": "cpp", ".cc": "cpp",
		".cs": "csharp", ".rs": "rust", ".kt": "kotlin",
		".swift": "swift", ".sh": "bash", ".bat": "bat", ".ps1": "powershell",
		".sql": "sql", ".json": "json", ".yaml": "yaml", ".yml": "yaml",
		".toml": "toml", ".ini": "ini", ".md": "md", ".tex": "latex",
		".r": "r", ".m": "matlab",
	}
}

func defaultSkipDirs() map[string]bool {
	return map[string]bool{
		".git": true, "node_modules": true, "dist": true, "build": true,
		"out": true, "venv": true, ".venv": true, ".idea": true, ".vscode": true,"":true,
	}
}

```

### `main.go`

```go
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
	dirBase := filepath.Base(dir)
	outFile := dirBase + ".md"

	m := NewModel(dir, outFile, *maxBytes)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}
}

```

### `render.go`

```go
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

```

### `teste/main.py`

```python
print("cocaina")
```

### `tree.go`

```go
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

```

### `tui.go`

```go
package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type stage int

const (
	stageScan stage = iota
	stageRender
	stageDone
)

type model struct {
	// config
	dir      string
	outFile  string
	maxBytes int64

	// ui
	stg         stage
	status      string
	err         error
	spin        spinner.Model
	prog        progress.Model
	doneSteps   int
	totalSteps  int
	treeLines   []string // árvore para exibir
	maxTreeShow int      // quantas linhas mostrar
	colorIdx    int      // animação de cor do banner

	// dados
	codeFiles []codeFile
}

// mensagens internas
type (
	tickColorMsg    struct{}
	tickProgressMsg struct{}
	scanDoneMsg     struct{ lines []string; codeFiles []codeFile }
	renderDoneMsg   struct{}
	errMsg          struct{ err error }
)

func NewModel(dir, outFile string, maxBytes int64) model {
	// spinner (bubbles)
	sp := spinner.New()
	sp.Spinner = spinner.MiniDot

	// progress (bubbles)
	p := progress.New(progress.WithDefaultGradient())
	p.Width = 60

	return model{
		dir:         dir,
		outFile:     outFile,
		maxBytes:    maxBytes,
		stg:         stageScan,
		status:      "🔎 Iniciando varredura…",
		spin:        sp,
		prog:        p,
		maxTreeShow: 60,
	}
}

// ASCII art do “logo” (pode trocar)
var asciiArt = []string{
	".............................................................. ........................",
	"............................................................  .........................",
	".............................................................-%@@@%-...................",
	"........  ...=*##*=-.................................... ..=@@@@#%@@+..................",
	"........  .=@@@%@@@@@@#:..............................  .:%@@@#===%@@-.................",
	"....... ..+@@+++#%@@%@@@@%-.............................*@@@@+==++=@@%.................",
	"..........%@*======+%@##%@@@*..................... . .*@@@@@+=+**+=%@@-................",
	"..........%@*========+@@#*#@@@%+:..................-#@@@%%@+=+**===*@@+................",
	"..........#@#===+*****+*@@#*#%@@@@#-..:-+*%@@@@@@@@@@@@##@+==**====*@@*................",
	"..........=@@=======+***=+#%@@%#%@@@@@@@@@@@@%%+===-=*%@@@#*+======+@@#................",
	"...........%@#========+**+====+%@#*##*##*****=:::::::::::=@@%**+===+@@%................",
	"...........=@@+=========+++=+*#%%#*******=:::::::::::::.:..-#@%#*+=+%@%................",
	"............#@%===========+*#%@#******=:::::.::::::::::::=@@@@@@%#*+%@@:...............",
	".............%@%=========**%@#*****+-+%@@%@%=::::::::::::@@:+@@*#@@#%@@=.. ............",
	".............:%@#+#**++=**@@******=-@@@@+.*@@@-::::::::::%@*@@@@=:%@@@@@...............",
	"..............-@@*+*#@#*#@%*******=@@@@%@@@@@@%:::::::::::+%@@%#-:::*@@@%-...  ........",
	"...............=@@*=+@##@#*******=:.-#%%@@@@%*:.:::.::::::--:.:::::::.-@@@*..  ........",
	"................#@@*+##@#******##+::::::::::.:::::::*@%@@%%@=.::::::::::*@@*...........",
	"............. . -@@%*#%*******%%%=----------:::::::::+%@@@@=:::::::::::::#@@-. ........",
	"..............  .%@@##******#%%%%----------=------------%%------=*%=------@@#..........",
	"...............  +@@%%%#**##%%%%*----------#@#++=======-#@===+#@@#--------%@@..........",
	"................ :@@%%%%%%%%%%%%+------------=+**##%%%@@@@@@%#+=----------%@@..........",
	"..................%@@%%%%%%%%%%%=----------------------------------------=@@%..........",
	"..................-@@@%%%%%%%%%#-----------------------------------------#@@=..........",
	"............... ...*@@%%%%%%%%%+----------------------------------------#@@*...........",
	"...............   ..%@@%%%%%%%+----------------------------------------%@@*............",
	"................  ...%@@%%%%#=-----------------------------===--------#@@*.............",
	"................. ...-@@@%#=-------------------------=======---------%@@*... ..........",
	".................. ..:%@#=------------------------------------------%@@=.... ..........",
	"................. ...-@@+------------------------------------------=@@#.    ...........",
	".................  ..#@%-------------------------------------------+@@*.    ...........",
	"................. ..%@@--------------------------------------------+@@*.    ...........",
	"...................*@@=--------------------------------------------=%@#.... ...........",
	"............... ..=@@+----------------------------------------------#@@:...  ..........",
	"............... .-@@+-----------------------------------------------=@@#..   ..........",
	"................:%@#-------------------------------------------------#@@:.   ..........",
	"................*@@=-------------------------------------------------=@@*.   ..........",
	"...............:@@+--------------------------------------------------=%@#..............",
	"............ ..%@%----------------------------------------------------%@%:  ...........",
	"............ .-@@=----------------------------------------------------#@%:  ...........",
	"..............*@%=----------------------------------------------------#@@-  ...........",
	"...........  .#@*-----------------------------------------------------*@@-. ...........",
	".............:%@*-----------------------------------------------------*@@-  ...........",
	"...........  :%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@-  ...........",
}

var palette = []lipgloss.Color{
	lipgloss.Color("#A6E3A1"),
	lipgloss.Color("#89B4FA"),
	lipgloss.Color("#F5C2E7"),
	lipgloss.Color("#F38BA8"),
	lipgloss.Color("#94E2D5"),
	lipgloss.Color("#F9E2AF"),
}

func (m model) Init() tea.Cmd {
	// Começa a animar o banner e o spinner, e já inicia o scan
	return tea.Batch(
		m.spin.Tick,
		tickColor(),
		startScan(m.dir, m.outFile),
	)
}

func tickColor() tea.Cmd {
	return tea.Tick(90*time.Millisecond, func(time.Time) tea.Msg { return tickColorMsg{} })
}

func tickProgress() tea.Cmd {
	// anima barra durante o render; o valor real é atualizado no Update
	return tea.Tick(70*time.Millisecond, func(time.Time) tea.Msg { return tickProgressMsg{} })
}

func startScan(dir, outFile string) tea.Cmd {
	// monta árvore e coleta os arquivos de código
	return func() tea.Msg {
		root, err := buildTree(dir, defaultSkipDirs())
		if err != nil {
			return errMsg{err}
		}
		var lines []string
		printTree(root, "", &lines)

		cf, err := collectCodeFiles(dir, outFile, defaultLangByExt())
		if err != nil {
			return errMsg{err}
		}
		return scanDoneMsg{lines: lines, codeFiles: cf}
	}
}

func startRender(dir, outFile string, tree []string, code []codeFile, maxBytes int64) tea.Cmd {
	// gera o markdown de forma síncrona; a barra é animada por tickProgress
	return func() tea.Msg {
		if err := renderMarkdown(dir, outFile, tree, code, maxBytes); err != nil {
			return errMsg{err}
		}
		return renderDoneMsg{}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// teclas de saída
	if km, ok := msg.(tea.KeyMsg); ok {
		switch km.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spin, cmd = m.spin.Update(msg)
		return m, cmd

	case tickColorMsg:
		m.colorIdx = (m.colorIdx + 1) % len(palette)
		return m, tickColor()

	case scanDoneMsg:
		m.treeLines = msg.lines
		m.codeFiles = msg.codeFiles
		m.stg = stageRender
		m.status = "📝 Gerando Markdown…"
		// barra “fake”/suave enquanto render ocorre (o render é síncrono)
		m.doneSteps = 0
		m.totalSteps = 100
		return m, tea.Batch(
			startRender(m.dir, m.outFile, m.treeLines, m.codeFiles, m.maxBytes),
			tickProgress(),
		)

	case tickProgressMsg:
		// durante o render, anima a barra
		if m.stg == stageRender && m.doneSteps < m.totalSteps {
			m.doneSteps++
			return m, tickProgress()
		}
		return m, nil

	case renderDoneMsg:
		m.doneSteps = m.totalSteps
		m.stg = stageDone
		m.status = "✅ Concluído: " + m.outFile + " !!!\nI love you, bye! 💕"

		// não sai automaticamente (mantém funcionalidades e permanece na tela)
		return m, tea.Quit

	case errMsg:
		m.err = msg.err
		m.stg = stageDone
		m.status = "❌ Erro: " + msg.err.Error()
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	// estilos
	col := palette[m.colorIdx]
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(col)
	subStyle := lipgloss.NewStyle().Faint(true)

	// header com spinner + emojis
	header := fmt.Sprintf("%s  %s  %s",
		m.spin.View(),
		titleStyle.Render(" Cartola - Corra e Olhe o Céu"),
		"☀️",
	)

	// banner ASCII colorido (sempre visível)
	art := ""
	for _, ln := range asciiArt {
		art += titleStyle.Render(ln) + "\n"
	}

	// progresso (0..1)
	var bar string
	if m.totalSteps > 0 {
		pct := float64(m.doneSteps) / float64(m.totalSteps)
		if pct > 1 {
			pct = 1
		}
		bar = m.prog.ViewAs(pct)
	} else {
		bar = m.prog.ViewAs(0)
	}

	body := ""

	body += "\n" + bar + "\n"
	if m.err != nil {
		body += subStyle.Render(m.status) + "\n"
	} else {
		body += subStyle.Render("Status: "+m.status) + "\n"
	}

	footer := subStyle.Render("Dicas: q/esc/ctrl+c para sair • feito com Bubble Tea 🫖 + Lip Gloss 💄")

	// layout
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		art,
		body,
		"",
		footer,
	)
}

```


```

### `collect.go`

```go
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

```

### `filters.go`

```go
package main

func defaultLangByExt() map[string]string {
	return map[string]string{
		".go": "go", ".html": "html", ".htm": "html", ".css": "css",
		".js": "javascript", ".mjs": "javascript", ".cjs": "javascript",
		".ts": "typescript", ".tsx": "tsx", ".jsx": "jsx",
		".py": "python", ".rb": "ruby", ".php": "php",
		".java": "java", ".c": "c", ".h": "c",
		".hpp": "cpp", ".cpp": "cpp", ".cc": "cpp",
		".cs": "csharp", ".rs": "rust", ".kt": "kotlin",
		".swift": "swift", ".sh": "bash", ".bat": "bat", ".ps1": "powershell",
		".sql": "sql", ".json": "json", ".yaml": "yaml", ".yml": "yaml",
		".toml": "toml", ".ini": "ini", ".md": "md", ".tex": "latex",
		".r": "r", ".m": "matlab",
	}
}

func defaultSkipDirs() map[string]bool {
	return map[string]bool{
		".git": true, "node_modules": true, "dist": true, "build": true,
		"out": true, "venv": true, ".venv": true, ".idea": true, ".vscode": true,"":true,
	}
}

```

### `main.go`

```go
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
	outFile :=  "lista_de_arquivos.md"

	m := NewModel(dir, outFile, *maxBytes)

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Fatal(err)
	}
}

```

### `render.go`

```go
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

```

### `teste/main.py`

```python
print("cocaina")
```

### `tree.go`

```go
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

```

### `tui.go`

```go
package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type stage int

const (
	stageScan stage = iota
	stageRender
	stageDone
)

type model struct {
	// config
	dir      string
	outFile  string
	maxBytes int64

	// ui
	stg         stage
	status      string
	err         error
	spin        spinner.Model
	prog        progress.Model
	doneSteps   int
	totalSteps  int
	treeLines   []string // árvore para exibir
	maxTreeShow int      // quantas linhas mostrar
	colorIdx    int      // animação de cor do banner

	// dados
	codeFiles []codeFile
}

// mensagens internas
type (
	tickColorMsg    struct{}
	tickProgressMsg struct{}
	scanDoneMsg     struct{ lines []string; codeFiles []codeFile }
	renderDoneMsg   struct{}
	errMsg          struct{ err error }
)

func NewModel(dir, outFile string, maxBytes int64) model {
	// spinner (bubbles)
	sp := spinner.New()
	sp.Spinner = spinner.MiniDot

	// progress (bubbles)
	p := progress.New(progress.WithDefaultGradient())
	p.Width = 60

	return model{
		dir:         dir,
		outFile:     outFile,
		maxBytes:    maxBytes,
		stg:         stageScan,
		status:      "🔎 Iniciando varredura…",
		spin:        sp,
		prog:        p,
		maxTreeShow: 60,
	}
}

// ASCII art do “logo” (pode trocar)
var asciiArt = []string{
	".............................................................. ........................",
	"............................................................  .........................",
	".............................................................-%@@@%-...................",
	"........  ...=*##*=-.................................... ..=@@@@#%@@+..................",
	"........  .=@@@%@@@@@@#:..............................  .:%@@@#===%@@-.................",
	"....... ..+@@+++#%@@%@@@@%-.............................*@@@@+==++=@@%.................",
	"..........%@*======+%@##%@@@*..................... . .*@@@@@+=+**+=%@@-................",
	"..........%@*========+@@#*#@@@%+:..................-#@@@%%@+=+**===*@@+................",
	"..........#@#===+*****+*@@#*#%@@@@#-..:-+*%@@@@@@@@@@@@##@+==**====*@@*................",
	"..........=@@=======+***=+#%@@%#%@@@@@@@@@@@@%%+===-=*%@@@#*+======+@@#................",
	"...........%@#========+**+====+%@#*##*##*****=:::::::::::=@@%**+===+@@%................",
	"...........=@@+=========+++=+*#%%#*******=:::::::::::::.:..-#@%#*+=+%@%................",
	"............#@%===========+*#%@#******=:::::.::::::::::::=@@@@@@%#*+%@@:...............",
	".............%@%=========**%@#*****+-+%@@%@%=::::::::::::@@:+@@*#@@#%@@=.. ............",
	".............:%@#+#**++=**@@******=-@@@@+.*@@@-::::::::::%@*@@@@=:%@@@@@...............",
	"..............-@@*+*#@#*#@%*******=@@@@%@@@@@@%:::::::::::+%@@%#-:::*@@@%-...  ........",
	"...............=@@*=+@##@#*******=:.-#%%@@@@%*:.:::.::::::--:.:::::::.-@@@*..  ........",
	"................#@@*+##@#******##+::::::::::.:::::::*@%@@%%@=.::::::::::*@@*...........",
	"............. . -@@%*#%*******%%%=----------:::::::::+%@@@@=:::::::::::::#@@-. ........",
	"..............  .%@@##******#%%%%----------=------------%%------=*%=------@@#..........",
	"...............  +@@%%%#**##%%%%*----------#@#++=======-#@===+#@@#--------%@@..........",
	"................ :@@%%%%%%%%%%%%+------------=+**##%%%@@@@@@%#+=----------%@@..........",
	"..................%@@%%%%%%%%%%%=----------------------------------------=@@%..........",
	"..................-@@@%%%%%%%%%#-----------------------------------------#@@=..........",
	"............... ...*@@%%%%%%%%%+----------------------------------------#@@*...........",
	"...............   ..%@@%%%%%%%+----------------------------------------%@@*............",
	"................  ...%@@%%%%#=-----------------------------===--------#@@*.............",
	"................. ...-@@@%#=-------------------------=======---------%@@*... ..........",
	".................. ..:%@#=------------------------------------------%@@=.... ..........",
	"................. ...-@@+------------------------------------------=@@#.    ...........",
	".................  ..#@%-------------------------------------------+@@*.    ...........",
	"................. ..%@@--------------------------------------------+@@*.    ...........",
	"...................*@@=--------------------------------------------=%@#.... ...........",
	"............... ..=@@+----------------------------------------------#@@:...  ..........",
	"............... .-@@+-----------------------------------------------=@@#..   ..........",
	"................:%@#-------------------------------------------------#@@:.   ..........",
	"................*@@=-------------------------------------------------=@@*.   ..........",
	"...............:@@+--------------------------------------------------=%@#..............",
	"............ ..%@%----------------------------------------------------%@%:  ...........",
	"............ .-@@=----------------------------------------------------#@%:  ...........",
	"..............*@%=----------------------------------------------------#@@-  ...........",
	"...........  .#@*-----------------------------------------------------*@@-. ...........",
	".............:%@*-----------------------------------------------------*@@-  ...........",
	"...........  :%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@-  ...........",
}

var palette = []lipgloss.Color{
	lipgloss.Color("#A6E3A1"),
	lipgloss.Color("#89B4FA"),
	lipgloss.Color("#F5C2E7"),
	lipgloss.Color("#F38BA8"),
	lipgloss.Color("#94E2D5"),
	lipgloss.Color("#F9E2AF"),
}

func (m model) Init() tea.Cmd {
	// Começa a animar o banner e o spinner, e já inicia o scan
	return tea.Batch(
		m.spin.Tick,
		tickColor(),
		startScan(m.dir, m.outFile),
	)
}

func tickColor() tea.Cmd {
	return tea.Tick(90*time.Millisecond, func(time.Time) tea.Msg { return tickColorMsg{} })
}

func tickProgress() tea.Cmd {
	// anima barra durante o render; o valor real é atualizado no Update
	return tea.Tick(70*time.Millisecond, func(time.Time) tea.Msg { return tickProgressMsg{} })
}

func startScan(dir, outFile string) tea.Cmd {
	// monta árvore e coleta os arquivos de código
	return func() tea.Msg {
		root, err := buildTree(dir, defaultSkipDirs())
		if err != nil {
			return errMsg{err}
		}
		var lines []string
		printTree(root, "", &lines)

		cf, err := collectCodeFiles(dir, outFile, defaultLangByExt())
		if err != nil {
			return errMsg{err}
		}
		return scanDoneMsg{lines: lines, codeFiles: cf}
	}
}

func startRender(dir, outFile string, tree []string, code []codeFile, maxBytes int64) tea.Cmd {
	// gera o markdown de forma síncrona; a barra é animada por tickProgress
	return func() tea.Msg {
		if err := renderMarkdown(dir, outFile, tree, code, maxBytes); err != nil {
			return errMsg{err}
		}
		return renderDoneMsg{}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// teclas de saída
	if km, ok := msg.(tea.KeyMsg); ok {
		switch km.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spin, cmd = m.spin.Update(msg)
		return m, cmd

	case tickColorMsg:
		m.colorIdx = (m.colorIdx + 1) % len(palette)
		return m, tickColor()

	case scanDoneMsg:
		m.treeLines = msg.lines
		m.codeFiles = msg.codeFiles
		m.stg = stageRender
		m.status = "📝 Gerando Markdown…"
		// barra “fake”/suave enquanto render ocorre (o render é síncrono)
		m.doneSteps = 0
		m.totalSteps = 100
		return m, tea.Batch(
			startRender(m.dir, m.outFile, m.treeLines, m.codeFiles, m.maxBytes),
			tickProgress(),
		)

	case tickProgressMsg:
		// durante o render, anima a barra
		if m.stg == stageRender && m.doneSteps < m.totalSteps {
			m.doneSteps++
			return m, tickProgress()
		}
		return m, nil

	case renderDoneMsg:
		m.doneSteps = m.totalSteps
		m.stg = stageDone
		m.status = "✅ Concluído: " + m.outFile + " !!!\nI love you, bye! 💕"

		// não sai automaticamente (mantém funcionalidades e permanece na tela)
		return m, tea.Quit

	case errMsg:
		m.err = msg.err
		m.stg = stageDone
		m.status = "❌ Erro: " + msg.err.Error()
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	// estilos
	col := palette[m.colorIdx]
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(col)
	subStyle := lipgloss.NewStyle().Faint(true)

	// header com spinner + emojis
	header := fmt.Sprintf("%s  %s  %s",
		m.spin.View(),
		titleStyle.Render(" Cartola - Corra e Olhe o Céu"),
		"☀️",
	)

	// banner ASCII colorido (sempre visível)
	art := ""
	for _, ln := range asciiArt {
		art += titleStyle.Render(ln) + "\n"
	}

	// progresso (0..1)
	var bar string
	if m.totalSteps > 0 {
		pct := float64(m.doneSteps) / float64(m.totalSteps)
		if pct > 1 {
			pct = 1
		}
		bar = m.prog.ViewAs(pct)
	} else {
		bar = m.prog.ViewAs(0)
	}

	body := ""

	body += "\n" + bar + "\n"
	if m.err != nil {
		body += subStyle.Render(m.status) + "\n"
	} else {
		body += subStyle.Render("Status: "+m.status) + "\n"
	}

	footer := subStyle.Render("Dicas: q/esc/ctrl+c para sair • feito com Bubble Tea 🫖 + Lip Gloss 💄")

	// layout
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		art,
		body,
		"",
		footer,
	)
}

```


```

### `main.go`

```go
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

```

### `render.go`

```go
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

```

### `teste/main.py`

```python
print("cocaina")
```

### `tree.go`

```go
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

```

### `tui.go`

```go
package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type stage int

const (
	stageScan stage = iota
	stageRender
	stageDone
)

type model struct {
	// config
	dir      string
	outFile  string
	maxBytes int64

	// ui
	stg         stage
	status      string
	err         error
	spin        spinner.Model
	prog        progress.Model
	doneSteps   int
	totalSteps  int
	treeLines   []string // árvore para exibir
	maxTreeShow int      // quantas linhas mostrar
	colorIdx    int      // animação de cor do banner

	// dados
	codeFiles []codeFile
}

// mensagens internas
type (
	tickColorMsg    struct{}
	tickProgressMsg struct{}
	scanDoneMsg     struct{ lines []string; codeFiles []codeFile }
	renderDoneMsg   struct{}
	errMsg          struct{ err error }
)

func NewModel(dir, outFile string, maxBytes int64) model {
	// spinner (bubbles)
	sp := spinner.New()
	sp.Spinner = spinner.MiniDot

	// progress (bubbles)
	p := progress.New(progress.WithDefaultGradient())
	p.Width = 60

	return model{
		dir:         dir,
		outFile:     outFile,
		maxBytes:    maxBytes,
		stg:         stageScan,
		status:      "🔎 Iniciando varredura…",
		spin:        sp,
		prog:        p,
		maxTreeShow: 60,
	}
}

// ASCII art do “logo” (pode trocar)
var asciiArt = []string{
	".............................................................. ........................",
	"............................................................  .........................",
	".............................................................-%@@@%-...................",
	"........  ...=*##*=-.................................... ..=@@@@#%@@+..................",
	"........  .=@@@%@@@@@@#:..............................  .:%@@@#===%@@-.................",
	"....... ..+@@+++#%@@%@@@@%-.............................*@@@@+==++=@@%.................",
	"..........%@*======+%@##%@@@*..................... . .*@@@@@+=+**+=%@@-................",
	"..........%@*========+@@#*#@@@%+:..................-#@@@%%@+=+**===*@@+................",
	"..........#@#===+*****+*@@#*#%@@@@#-..:-+*%@@@@@@@@@@@@##@+==**====*@@*................",
	"..........=@@=======+***=+#%@@%#%@@@@@@@@@@@@%%+===-=*%@@@#*+======+@@#................",
	"...........%@#========+**+====+%@#*##*##*****=:::::::::::=@@%**+===+@@%................",
	"...........=@@+=========+++=+*#%%#*******=:::::::::::::.:..-#@%#*+=+%@%................",
	"............#@%===========+*#%@#******=:::::.::::::::::::=@@@@@@%#*+%@@:...............",
	".............%@%=========**%@#*****+-+%@@%@%=::::::::::::@@:+@@*#@@#%@@=.. ............",
	".............:%@#+#**++=**@@******=-@@@@+.*@@@-::::::::::%@*@@@@=:%@@@@@...............",
	"..............-@@*+*#@#*#@%*******=@@@@%@@@@@@%:::::::::::+%@@%#-:::*@@@%-...  ........",
	"...............=@@*=+@##@#*******=:.-#%%@@@@%*:.:::.::::::--:.:::::::.-@@@*..  ........",
	"................#@@*+##@#******##+::::::::::.:::::::*@%@@%%@=.::::::::::*@@*...........",
	"............. . -@@%*#%*******%%%=----------:::::::::+%@@@@=:::::::::::::#@@-. ........",
	"..............  .%@@##******#%%%%----------=------------%%------=*%=------@@#..........",
	"...............  +@@%%%#**##%%%%*----------#@#++=======-#@===+#@@#--------%@@..........",
	"................ :@@%%%%%%%%%%%%+------------=+**##%%%@@@@@@%#+=----------%@@..........",
	"..................%@@%%%%%%%%%%%=----------------------------------------=@@%..........",
	"..................-@@@%%%%%%%%%#-----------------------------------------#@@=..........",
	"............... ...*@@%%%%%%%%%+----------------------------------------#@@*...........",
	"...............   ..%@@%%%%%%%+----------------------------------------%@@*............",
	"................  ...%@@%%%%#=-----------------------------===--------#@@*.............",
	"................. ...-@@@%#=-------------------------=======---------%@@*... ..........",
	".................. ..:%@#=------------------------------------------%@@=.... ..........",
	"................. ...-@@+------------------------------------------=@@#.    ...........",
	".................  ..#@%-------------------------------------------+@@*.    ...........",
	"................. ..%@@--------------------------------------------+@@*.    ...........",
	"...................*@@=--------------------------------------------=%@#.... ...........",
	"............... ..=@@+----------------------------------------------#@@:...  ..........",
	"............... .-@@+-----------------------------------------------=@@#..   ..........",
	"................:%@#-------------------------------------------------#@@:.   ..........",
	"................*@@=-------------------------------------------------=@@*.   ..........",
	"...............:@@+--------------------------------------------------=%@#..............",
	"............ ..%@%----------------------------------------------------%@%:  ...........",
	"............ .-@@=----------------------------------------------------#@%:  ...........",
	"..............*@%=----------------------------------------------------#@@-  ...........",
	"...........  .#@*-----------------------------------------------------*@@-. ...........",
	".............:%@*-----------------------------------------------------*@@-  ...........",
	"...........  :%@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@-  ...........",
}

var palette = []lipgloss.Color{
	lipgloss.Color("#A6E3A1"),
	lipgloss.Color("#89B4FA"),
	lipgloss.Color("#F5C2E7"),
	lipgloss.Color("#F38BA8"),
	lipgloss.Color("#94E2D5"),
	lipgloss.Color("#F9E2AF"),
}

func (m model) Init() tea.Cmd {
	// Começa a animar o banner e o spinner, e já inicia o scan
	return tea.Batch(
		m.spin.Tick,
		tickColor(),
		startScan(m.dir, m.outFile),
	)
}

func tickColor() tea.Cmd {
	return tea.Tick(90*time.Millisecond, func(time.Time) tea.Msg { return tickColorMsg{} })
}

func tickProgress() tea.Cmd {
	// anima barra durante o render; o valor real é atualizado no Update
	return tea.Tick(70*time.Millisecond, func(time.Time) tea.Msg { return tickProgressMsg{} })
}

func startScan(dir, outFile string) tea.Cmd {
	// monta árvore e coleta os arquivos de código
	return func() tea.Msg {
		root, err := buildTree(dir, defaultSkipDirs())
		if err != nil {
			return errMsg{err}
		}
		var lines []string
		printTree(root, "", &lines)

		cf, err := collectCodeFiles(dir, outFile, defaultLangByExt())
		if err != nil {
			return errMsg{err}
		}
		return scanDoneMsg{lines: lines, codeFiles: cf}
	}
}

func startRender(dir, outFile string, tree []string, code []codeFile, maxBytes int64) tea.Cmd {
	// gera o markdown de forma síncrona; a barra é animada por tickProgress
	return func() tea.Msg {
		if err := renderMarkdown(dir, outFile, tree, code, maxBytes); err != nil {
			return errMsg{err}
		}
		return renderDoneMsg{}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// teclas de saída
	if km, ok := msg.(tea.KeyMsg); ok {
		switch km.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}

	switch msg := msg.(type) {
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spin, cmd = m.spin.Update(msg)
		return m, cmd

	case tickColorMsg:
		m.colorIdx = (m.colorIdx + 1) % len(palette)
		return m, tickColor()

	case scanDoneMsg:
		m.treeLines = msg.lines
		m.codeFiles = msg.codeFiles
		m.stg = stageRender
		m.status = "📝 Gerando Markdown…"
		// barra “fake”/suave enquanto render ocorre (o render é síncrono)
		m.doneSteps = 0
		m.totalSteps = 100
		return m, tea.Batch(
			startRender(m.dir, m.outFile, m.treeLines, m.codeFiles, m.maxBytes),
			tickProgress(),
		)

	case tickProgressMsg:
		// durante o render, anima a barra
		if m.stg == stageRender && m.doneSteps < m.totalSteps {
			m.doneSteps++
			return m, tickProgress()
		}
		return m, nil

	case renderDoneMsg:
		m.doneSteps = m.totalSteps
		m.stg = stageDone
		m.status = "✅ Concluído: " + m.outFile + " !!!\nI love you, bye! 💕"

		// não sai automaticamente (mantém funcionalidades e permanece na tela)
		return m, tea.Quit

	case errMsg:
		m.err = msg.err
		m.stg = stageDone
		m.status = "❌ Erro: " + msg.err.Error()
		return m, nil
	}

	return m, nil
}

func (m model) View() string {
	// estilos
	col := palette[m.colorIdx]
	titleStyle := lipgloss.NewStyle().Bold(true).Foreground(col)
	subStyle := lipgloss.NewStyle().Faint(true)

	// header com spinner + emojis
	header := fmt.Sprintf("%s  %s  %s",
		m.spin.View(),
		titleStyle.Render(" Cartola - Corra e Olhe o Céu"),
		"☀️",
	)

	// banner ASCII colorido (sempre visível)
	art := ""
	for _, ln := range asciiArt {
		art += titleStyle.Render(ln) + "\n"
	}

	// progresso (0..1)
	var bar string
	if m.totalSteps > 0 {
		pct := float64(m.doneSteps) / float64(m.totalSteps)
		if pct > 1 {
			pct = 1
		}
		bar = m.prog.ViewAs(pct)
	} else {
		bar = m.prog.ViewAs(0)
	}

	body := ""

	body += "\n" + bar + "\n"
	if m.err != nil {
		body += subStyle.Render(m.status) + "\n"
	} else {
		body += subStyle.Render("Status: "+m.status) + "\n"
	}

	footer := subStyle.Render("Dicas: q/esc/ctrl+c para sair • feito com Bubble Tea 🫖 + Lip Gloss 💄")

	// layout
	return lipgloss.JoinVertical(lipgloss.Left,
		header,
		"",
		art,
		body,
		"",
		footer,
	)
}

```

