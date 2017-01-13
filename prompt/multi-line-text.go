package prompt

import "github.com/josa42/go-prompt/elements"

// MultiLineText :
func MultiLineText(label string, defaultValue string) (value string, canceled bool) {

	input := elements.Input{
		Label:        label,
		DefaultValue: defaultValue,
		MultiLine:    true,
	}

	value, canceled = input.Run()

	return
}
