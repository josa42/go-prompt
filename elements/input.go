package elements

import (
	"bufio"
	"fmt"
	"os"

	"github.com/buger/goterm"
)

// TODO handle and return cancel
// TODO add validation
// TODO Handle controll sequences, eg. cursor move movement

// Input :
type Input struct {
	Label        string
	DefaultValue string
}

// Run :
func (i *Input) Run() (value string, canceled bool) {
	fmt.Printf("%s", goterm.Color(goterm.Bold(i.Label), goterm.GREEN))

	if i.DefaultValue != "" {
		fmt.Printf(" %s%s%s",
			goterm.Color(goterm.Bold("["), goterm.GREEN),
			goterm.Color(i.DefaultValue, goterm.YELLOW),
			goterm.Color(goterm.Bold("]"), goterm.GREEN))
	}

	fmt.Printf("%s ", goterm.Color(goterm.Bold(":"), goterm.GREEN))

	reader := bufio.NewReader(os.Stdin)
	value, _ = reader.ReadString('\n')

	if value[len(value)-1] == '\n' {
		value = value[:len(value)-1]
	}

	if value == "" {
		value = i.DefaultValue
	}

	return
}
