package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	// Define the size of our playing field, add 2 for the borders.
	tableWidth  = 60 + 2
	tableHeight = 30 + 2

	// Define our symbols for representing the playing field.
	corner         = '+'
	lineVertical   = '|'
	lineHorizontal = '-'
	empty          = ' '
	player         = '0'
	item           = '$'
	enemy          = 'X'
)

type model struct {
	table [tableHeight][tableWidth]rune

	playerRow int
	playerCol int
}

func (m *model) move(row, col int) {
	m.table[m.playerRow][m.playerCol] = 0

	m.table[row][col] = player
	m.playerRow = row
	m.playerCol = col
}

func (m *model) playerUp() {
	if m.playerRow <= 1 {
		return
	}
	m.move(m.playerRow-1, m.playerCol)
}

func (m *model) playerDown() {
	if m.playerRow >= tableHeight-2 {
		return
	}
	m.move(m.playerRow+1, m.playerCol)
}

func (m *model) playerLeft() {
	if m.playerCol <= 1 {
		return
	}
	m.move(m.playerRow, m.playerCol-1)
}

func (m *model) playerRight() {
	if m.playerCol >= tableWidth-2 {
		return
	}
	m.move(m.playerRow, m.playerCol+1)
}

func newModel() *model {
	model := &model{}

	model.table[0][0] = corner
	model.table[0][tableWidth-1] = corner
	model.table[tableHeight-1][0] = corner
	model.table[tableHeight-1][tableWidth-1] = corner

	for col := 1; col < tableWidth-1; col++ {
		model.table[0][col] = lineHorizontal
		model.table[tableHeight-1][col] = lineHorizontal
	}

	for row := 1; row < tableHeight-1; row++ {
		model.table[row][0] = lineVertical
		model.table[row][tableWidth-1] = lineVertical
	}

	model.playerRow = 10
	model.playerCol = 10
	model.table[model.playerRow][model.playerCol] = player

	return model
}

func (m *model) Init() tea.Cmd {
	return nil
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "Ctrl+C", "q":
			return m, tea.Quit
		case "up":
			m.playerUp()
		case "down":
			m.playerDown()
		case "left":
			m.playerLeft()
		case "right":
			m.playerRight()

		}

	}
	return m, nil
}

func (m *model) View() string {
	builder := strings.Builder{}

	for _, row := range m.table {
		for _, cell := range row {
			if cell == 0 {
				builder.WriteRune(empty)
			} else {
				builder.WriteRune(cell)
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func main() {
	model := newModel()

	program := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Println("Unexpected error: %v", err)
		os.Exit(1)
	}
}
