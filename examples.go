package main

import (
	"fmt"

	"github.com/josa42/go-prompt/prompt"
)

func main() {
	selection1, _ := prompt.Select("Choose one option", prompt.Options{
		{"1", "One"},
		{"2", "Two"},
		{"3", "Three"},
		{"4", "Four"},
		{"5", "Five"},
		{"6", "Six"},
		{"7", "Seven"},
		{"8", "Eight"},
		{"9", "Nine"},
	})
	fmt.Println()
	fmt.Println("Selection:", selection1)
	fmt.Println()

	selection2, _ := prompt.MultiSelect("Choose multiple options", prompt.Options{
		{"1", "One"},
		{"2", "Two"},
		{"3", "Three"},
	})
	fmt.Println()
	fmt.Println("Selection:", selection2)
	fmt.Println()

	selection3, _ := prompt.Confirm("Do you want to go on?")
	fmt.Println()
	fmt.Println("Selection:", selection3)
	fmt.Println()

	selection4, _ := prompt.Text("Write single line text", "")
	fmt.Println()
	fmt.Println("Selection:", selection4)
	fmt.Println()

	selection5, _ := prompt.MultiLineText("Write multiple line text", "")
	fmt.Println()
	fmt.Println("Selection:\n" + selection5)
	fmt.Println()
}
