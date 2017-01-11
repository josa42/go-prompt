package prompt

import "github.com/josa42/go-prompt/elements"

// Select :
func Select(label string, options Options) (selection string, canceled bool) {

	menu := elements.Select{
		Label: "Select options",
		Multi: false,
	}

	for value, label := range options {
		menu.AddOption(value, label)
	}

	var selections []string
	selections, canceled = menu.Run()

	selection = selections[0]

	return
}
