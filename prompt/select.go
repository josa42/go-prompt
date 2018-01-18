package prompt

import "github.com/josa42/go-prompt/elements"

// Select :
func Select(label string, options Options) (selection string, canceled bool) {

	menu := elements.Select{
		Label: "Select options",
		Multi: false,
	}

	for _, option := range options {
		menu.AddOption(option[0], option[1])
	}

	var selections []string
	selections, canceled = menu.Run()

	selection = selections[0]

	return
}
