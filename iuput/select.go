package input

import (
	"github.com/AlecAivazis/survey/v2"
)

func SingleSelect(label string, options []string) (string, error) {
	var selected string
	prompt := &survey.Select{
		Message: label,
		Options: options,
	}
	if err := survey.AskOne(prompt, &selected); err != nil {
		return "", err
	}
	return selected, nil
}

func MultiSelect(label string, options []string) ([]string, error) {
	var selected []string
	prompt := &survey.MultiSelect{
		Message: label,
		Options: options,
	}
	if err := survey.AskOne(prompt, &selected); err != nil {
		return nil, err
	}
	return selected, nil
}
