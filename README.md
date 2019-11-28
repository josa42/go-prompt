# Go Prompt

Interactive terminal inputs.

## Elements

### Select

![](docs/screenshot-select.jpg)

```Go
selection1, _ := prompt.Select("Choose one option", prompt.Options{
	{"1", "One"},
	{"2", "Two"},
	{"3", "Three"},
})
fmt.Println("Selection:", selection1)
```

### MultiSelect

![](docs/screenshot-multi-select.jpg)

```Go
selection, _ := prompt.MultiSelect("Choose multiple options", prompt.Options{
	{"1", "One"},
	{"2", "Two"},
	{"3", "Three"},
})
fmt.Println("Selection:", selection)
```

### Confirm

![](docs/screenshot-confirm.jpg)

```Go
selection, _ := prompt.Confirm("Do you want to go on?")
fmt.Println("Selection:", selection)
```

### Text Input

![](docs/screenshot-text.jpg)

```Go
selection, _ := prompt.Text("Write single line text", "")
fmt.Println("Selection:", selection)
```

### Multi Line Text Input

![](docs/screenshot-multi-line-text.jpg)

```Go
selection, _ := prompt.MultiLineText("Write multiple line text", "")
fmt.Println("Selection:\n" + selection)
```


## Prior Art

- [paulrademacher/climenu](https://github.com/paulrademacher/climenu)

## License

[MIT © Josa Gesell](LICENSE)
