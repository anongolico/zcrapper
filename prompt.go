package main

import (
	"github.com/AlecAivazis/survey/v2"
)

// ReadPostId prompts the user to enter the id of
// the post they want to download
func ReadPostId() (string, error) {
	var id string
	q := []*survey.Question{
		{
			Name: "Id",
			Prompt: &survey.Input{
				Message: "Ingrese el ID del post:",
			},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(q, &id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// ReadFormats prompts the user to select the formats
// they want to download.
func ReadFormats(Formats map[string][]string) error {
	formats := make([]string, 0)

	for k := range Formats {
		formats = append(formats, k)
	}

	q := []*survey.Question{
		{
			Name: "Formats",
			Prompt: &survey.MultiSelect{
				Message: "Escoge los formatos para descargar:",
				Options: formats,
			},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(q, &FormatsToDownload)

	return err
}
