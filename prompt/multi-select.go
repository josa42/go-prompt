package prompt

import "github.com/josa42/go-prompt/elements"

// MultiSelect :
func MultiSelect(label string, options Options) (selection []string, canceled bool) {

	menu := elements.Select{
		Label: "Select options",
		Multi: true,
	}

	for value, label := range options {
		menu.AddOption(value, label)
	}

	selection, canceled = menu.Run()

	return
}
