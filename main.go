package main

import (
	"fmt"
	"math/rand"
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

	score    int
	gameOver bool
}

func (m *model) spawnItem() {
	row, col := m.randomFreeCoordinates()
	m.table[row][col] = item
}

func (m *model) randomFreeCoordinates() (row, col int) {
	row, col = randomCoordinates()

	for m.table[row][col] != empty {
		row, col = randomCoordinates()
	}
	return
}

func randomCoordinates() (row, col int) {
	row = rand.Intn(tableHeight-2) + 1
	col = rand.Intn(tableWidth-2) + 1

	return row, col

}

func (m *model) movePlayer(row, col int) {
	m.table[m.playerRow][m.playerCol] = empty

	if m.table[row][col] == item {
		m.score++
		m.spawnItem()
	}

	m.table[row][col] = player
	m.playerRow = row
	m.playerCol = col
}

func (m *model) playerUp() {
	if m.playerRow <= 1 {
		return
	}
	m.movePlayer(m.playerRow-1, m.playerCol)
}

func (m *model) playerDown() {
	if m.playerRow >= tableHeight-2 {
		return
	}
	m.movePlayer(m.playerRow+1, m.playerCol)
}

func (m *model) playerLeft() {
	if m.playerCol <= 1 {
		return
	}
	m.movePlayer(m.playerRow, m.playerCol-1)
}

func (m *model) playerRight() {
	if m.playerCol >= tableWidth-2 {
		return
	}
	m.movePlayer(m.playerRow, m.playerCol+1)
}

func (m *model) init() {

	for row := 0; row < tableHeight; row++ {
		for col := 0; col < tableWidth; col++ {
			m.table[row][col] = empty
		}
	}

	m.playerRow = 0
	m.playerCol = 0
	m.score = 0
	m.gameOver = false

	m.table[0][0] = corner
	m.table[0][tableWidth-1] = corner
	m.table[tableHeight-1][0] = corner
	m.table[tableHeight-1][tableWidth-1] = corner

	for col := 1; col < tableWidth-1; col++ {
		m.table[0][col] = lineHorizontal
		m.table[tableHeight-1][col] = lineHorizontal
	}

	for row := 1; row < tableHeight-1; row++ {
		m.table[row][0] = lineVertical
		m.table[row][tableWidth-1] = lineVertical
	}

	m.playerRow = 10
	m.playerCol = 10
	m.table[m.playerRow][m.playerCol] = player
	m.spawnItem()
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
			builder.WriteRune(cell)
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func main() {
	model := &model{}

	model.init()

	program := tea.NewProgram(model, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Println("Unexpected error: %v", err)
		os.Exit(1)
	}
}
