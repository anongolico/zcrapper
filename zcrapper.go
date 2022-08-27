package main

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/anongolico/base"
	"log"
	"time"
)

var (
	FormatsToDownload    []string
	TotalFilesToDownload int
)

const (
	MaxParallelDownloads = 50
	mediaUrl             = base.BaseMediaUrl
)

func listFormats() {
	f := make([]string, len(base.Formats))
	i := 0
	for k := range base.Formats {
		f[i] = k
		i++
	}

	q := []*survey.Question{
		{
			Name: "formats",
			Prompt: &survey.MultiSelect{
				Message: "Escoge los formatos a descargar:",
				Options: f,
			},
			Validate: survey.Required,
		},
	}
	_ = survey.Ask(q, &FormatsToDownload)
}

func main() {
	base.ReadAuthFile()

	var id string
	questions := []*survey.Question{
		{
			Name:     "rouz",
			Prompt:   &survey.Input{Message: "Ingresa el ID del rouz:"},
			Validate: survey.Required,
		},
	}
	_ = survey.Ask(questions, &id)
	r := base.New(id)
	_ = base.ScanFormats(r)

	listFormats()

	for _, v := range FormatsToDownload {
		TotalFilesToDownload += len(base.Formats[v])
	}

	log.Printf("%d archivos para descargar. Puede tomar un tiempo.", TotalFilesToDownload+1)

	ts := time.Now()
	downloadFiles(r)
	log.Printf("Hasta la prótsimaaaa (tiempo total: %v)\n", time.Since(ts))

}
