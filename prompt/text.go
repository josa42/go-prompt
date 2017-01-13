package prompt

import "github.com/josa42/go-prompt/elements"

// Text :
func Text(label string, defaultValue string) (value string, canceled bool) {

	// selection = climenu.GetText(label, defaultValue)

	input := elements.Input{
		Label:        label,
		DefaultValue: defaultValue,
		MultiLine:    false,
	}

	value, canceled = input.Run()

	return
}
