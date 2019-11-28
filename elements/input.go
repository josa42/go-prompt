package elements

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/buger/goterm"
)

// TODO handle and return cancel
// TODO add validation
// TODO Handle controll sequences, eg. cursor move movement

// Input :
type Input struct {
	Label        string
	DefaultValue string
	MultiLine    bool
}

// Run :
func (i *Input) Run() (value string, canceled bool) {
	fmt.Printf("%s", goterm.Color(goterm.Bold(i.Label), goterm.GREEN))

	if i.DefaultValue != "" {
		fmt.Printf(" %s%s%s ",
			goterm.Color(goterm.Bold("["), goterm.GREEN),
			goterm.Color(i.DefaultValue, goterm.YELLOW),
			goterm.Color(goterm.Bold("]"), goterm.GREEN),
		)
	}

	fmt.Printf("%s\n%s ",
		goterm.Color(goterm.Bold(":"), goterm.GREEN),
		goterm.Color("\u276F", goterm.CYAN),
	)

	reader := bufio.NewReader(os.Stdin)

	if i.MultiLine {
		var lines []string

		for {
			line, _ := reader.ReadString('\n')
			if line == "\n" {
				fmt.Print("\033[1A \033[1A\n")
				break
			}

			fmt.Print(goterm.Color("\u276F ", goterm.CYAN))

			lines = append(lines, line[:len(line)-1])
		}

		value = strings.Join(lines, "\n")
	} else {
		value, _ = reader.ReadString('\n')
	}

	if len(value) > 0 && value[len(value)-1] == '\n' {
		value = value[:len(value)-1]
	}

	if value == "" {
		value = i.DefaultValue
	}

	return
}
