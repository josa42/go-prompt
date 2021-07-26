package prompt

import "github.com/josa42/go-prompt/elements"

// MultiSelect :
func MultiSelect(label string, options Options) (selection []string, canceled bool) {

	menu := elements.Select{
		Label: label,
		Multi: true,
	}

	for _, option := range options {
		menu.AddOption(option[0], option[0])
	}

	selection, canceled = menu.Run()

	return
}
