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

	// NOVO
	ignoreContentDirs []string

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
		startScan(m.dir, m.outFile, m.ignoreContentDirs),
	)
	
}

func tickColor() tea.Cmd {
	return tea.Tick(90*time.Millisecond, func(time.Time) tea.Msg { return tickColorMsg{} })
}

func tickProgress() tea.Cmd {
	// anima barra durante o render; o valor real é atualizado no Update
	return tea.Tick(70*time.Millisecond, func(time.Time) tea.Msg { return tickProgressMsg{} })
}


func startScan(dir, outFile string, ignoreContentDirs []string) tea.Cmd {
	return func() tea.Msg {
		root, err := buildTree(dir, defaultSkipDirs())
		if err != nil {
			return errMsg{err}
		}
		var lines []string
		printTree(root, "", &lines)

		// passa ignoreContentDirs aqui
		cf, err := collectCodeFiles(dir, outFile, defaultLangByExt(), ignoreContentDirs)
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
