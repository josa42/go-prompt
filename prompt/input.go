package prompt

import "github.com/josa42/go-prompt/elements"

// Input :
func Input(label string, defaultValue string) (value string, canceled bool) {

	// selection = climenu.GetText(label, defaultValue)

	input := elements.Input{
		Label:        label,
		DefaultValue: defaultValue,
	}

	value, canceled = input.Run()

	return
}
