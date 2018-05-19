package main

import (
	"fmt"

	"gopkg.in/AlecAivazis/survey.v1"
)

var qs = []*survey.Question{
	{
		Name: "name",
		Prompt: &survey.Select{
			Message: "Choose A or B",
			Options: []string{"A", "B"},
			Default: "bjtb",
		},
	},
	{
		Name: "color",
		Prompt: &survey.Select{
			Message: "Choose a color:",
			Options: []string{"nginx", "project"},
			Default: "red",
		},
	},
	{
		Name:   "host",
		Prompt: &survey.Input{Message: "How old are you?"},
	},
}

func main() {
	answers := struct {
		Name          string `survey:"name"`
		FavoriteColor string `survey:"color"`
		Age           int
	}{}
	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	fmt.Printf("%s chose %s.", answers.Name, answers.FavoriteColor)
}
