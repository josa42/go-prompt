package elements

import (
	"fmt"
	"os"

	"github.com/buger/goterm"
	"github.com/josa42/go-prompt/input"
)

type InputAction int

const (
	Default InputAction = iota
	Return
	Cancel
	None
)

// Select :
type Select struct {
	Multi          bool
	Label          string
	MaxVisible     int
	InputHandler   func(Select, input.Sequence) InputAction
	FilterHandler  func(Select, string) bool
	cursorPosition int
	options        []option
	visibleIndex   int
}

type option struct {
	key      string
	label    string
	selected bool
}

// AddOption :
func (m *Select) AddOption(key string, text string) {
	m.options = append(m.options, option{
		label: text,
		key:   key,
	})
}

// Run :
func (m *Select) Run() (results []string, canceled bool) {

	defer (func() {
		// m.cursorPosition = -1
		// m.redrawOptions()
		fmt.Println()
		m.showCursor()
	})()

	m.hideCursor()

	m.drawLabel()
	m.drawOptions()

	for {
		sequence, err := input.ReadSequence()

		if m.InputHandler != nil {
			switch m.InputHandler(*m, sequence) {
			case Return:
				return m.selectedKeys(), false
			case Cancel:
				return m.cancelKeys(), true
			case None:
				continue
			}
		}

		if sequence.IsEtx() {
			os.Exit(0)

		} else if sequence.IsEsc() || err != nil {
			return m.cancelKeys(), true

		} else if sequence.IsReturn() {
			return m.selectedKeys(), false

		} else if sequence.IsDown() {
			m.moveCursor(1)

		} else if sequence.IsUp() {
			m.moveCursor(-1)

		} else if sequence.IsSpace() {
			if m.Multi {
				m.toggleSelection()
				m.redrawOptions()
			}

		} else if sequence.IsNumber() {

			number := sequence.Number()

			if number < len(m.options) {
				m.setCursorPosition(number)

				if m.Multi {
					// Toggle option at index
					m.toggleSelection()
					m.redrawOptions()

				} else {
					// Select option at index and return
					m.redrawOptions()
					return m.selectedKeys(), false
				}
			}
			// } else if sequence.IsString("y") {
			// 	m.setCursorPosition(0)
			// 	m.redrawOptions()
			// } else if sequence.IsString("n") {
			// 	m.setCursorPosition(1)
			// 	m.redrawOptions()
		}
	}
}

func (m *Select) SelectIndex(index int) {
	m.setCursorPosition(index)
	m.redrawOptions()
}

func (m *Select) maxLen() int {
	l := len(m.options)
	if m.MaxVisible != 0 && m.MaxVisible < l {
		return m.MaxVisible
	}
	return l
}

func (m *Select) selectedKeys() (selections []string) {
	if m.Multi {
		for _, option := range m.options {
			if option.selected {
				selections = append(selections, option.key)
			}
		}
	} else {
		option := &m.options[m.cursorPosition]
		selections = append(selections, option.key)
	}

	return
}

func (m *Select) cancelKeys() []string {
	if m.Multi {
		return []string{}
	}
	return []string{""}
}

func (m *Select) setCursorPosition(position int) {
	length := len(m.options)
	m.cursorPosition = (position + length) % length

	if m.cursorPosition < m.visibleIndex {
		m.visibleIndex = m.cursorPosition
	}
	if m.cursorPosition >= m.visibleIndex+m.maxLen() {
		m.visibleIndex = (m.cursorPosition - m.maxLen() + 1)
	}
}

func (m *Select) moveCursor(diff int) {
	m.setCursorPosition(m.cursorPosition + diff)
	m.redrawOptions()
}

func (m *Select) toggleSelection() {
	option := &m.options[m.cursorPosition]

	if option.selected {
		option.selected = false
	} else {
		option.selected = true
	}
}

func (m *Select) showCursor() {
	fmt.Printf("\033[?25h")
}

func (m *Select) hideCursor() {
	fmt.Printf("\033[?25l")
}

func (m *Select) drawnOptionCount() int {
	c := len(m.options)
	if c > m.maxLen() {
		c = m.maxLen()
	}
	return c
}

func (m *Select) resetCursor() {
	fmt.Printf("\033[%dA", m.drawnOptionCount()-1)
}

func (m *Select) redrawOptions() {
	m.resetCursor()
	m.drawOptions()
}

func (m *Select) drawLabel() {
	fmt.Println(goterm.Color(goterm.Bold(m.Label)+":", goterm.GREEN))
}

func (m *Select) drawOptions() {

	count := 0

	for index, option := range m.options {

		if index < m.visibleIndex || index >= (m.visibleIndex+m.maxLen()) {
			continue
		}

		if m.FilterHandler != nil && !m.FilterHandler(*m, option.key) {
			continue
		}

		// Actually draw option
		count = count + 1

		// :Erase line
		fmt.Printf("\033[2K")

		prefix := ""
		if m.Multi {
			if option.selected {
				// TODO make configurable
				prefix = "\u25c9 "
			} else {
				// TODO make configurable
				prefix = "\u25ef "
			}
		}

		if index == m.cursorPosition {
			// TODO make configurable
			cursor := goterm.Color("\u276F ", goterm.CYAN)
			fmt.Printf("\r%s%s%s", cursor, prefix, goterm.Color(option.label, goterm.CYAN))

		} else {
			fmt.Printf("\r%s%s%s", "  ", prefix, option.label)
		}

		if count != m.maxLen() {
			fmt.Println()
		}
	}
}
