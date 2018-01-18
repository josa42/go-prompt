package prompt

// Confirm :
func Confirm(label string) (confirmed bool, canceled bool) {

	selection := ""
	selection, canceled = Select(label, Options{
		{"confirm", "Confirm"},
		{"cancel", "Cancel"},
	})

	confirmed = selection == "confirm"

	return
}
