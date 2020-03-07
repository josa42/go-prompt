package prompt

import (
	"github.com/josa42/go-prompt/elements"
	"github.com/josa42/go-prompt/input"
)

// Confirm :
func Confirm(label string) (confirmed bool, canceled bool) {

	// selection := ""
	// selection, canceled = Select(label, Options{
	// 	{"confirm", "Confirm"},
	// 	{"cancel", "Cancel"},
	// })
	//
	// // selection.InputHandler = func(s input.Sequenz) bool {
	// // 	return false
	// // }
	// //
	// confirmed = selection == "confirm"
	//
	// return

	menu := elements.Select{
		Label: label,
		Multi: false,
		InputHandler: func(sel elements.Select, seq input.Sequence) elements.InputAction {

			if seq.IsNumber() {
				return elements.None
			}

			if seq.IsString("y") {
				sel.SelectIndex(0)
				return elements.Return
			}

			if seq.IsString("n") {
				sel.SelectIndex(1)
				return elements.Return
			}

			return elements.Default
		},
	}

	menu.AddOption("confirm", "Confirm")
	menu.AddOption("cancel", "Cancel")

	var selections []string
	selections, canceled = menu.Run()

	confirmed = selections[0] == "confirm"
	return
}
