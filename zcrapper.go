package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/anongolico/base"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	mediaUrl             = base.BaseMediaUrl
	FormatsToDownload    []string
	TotalFilesToDownload int
)

// createRouzFolder creates a new folder to store media
func createRouzFolder(name string) {
	// algunos dawns ponen el slash en el nombre del roz y
	// eso genera error si el SO es linux
	name = strings.ReplaceAll(name, "/", "")
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = os.Mkdir(name, 0755)
		base.Handle(err, "")
	}
}

func downloadFiles(r *base.Rouz) {
	folderName := fmt.Sprintf("%s (%s)", r.Hilo.Titulo, r.Hilo.Id)
	createRouzFolder(folderName)
	err := os.Chdir(folderName)
	base.Handle(err, "")

	err = downloadFile(r.Hilo.Media.Url)

	for _, v := range FormatsToDownload {
		createRouzFolder(v)
		err = os.Chdir(v)
		base.Handle(err, "")
		for _, x := range base.Formats[v] {
			err = downloadFile(x)
			base.Handle(err, "error al descargar archivo")
			TotalFilesToDownload--
			log.Printf("%d restantes.\n", TotalFilesToDownload+1)
		}
		err = os.Chdir("..")
		base.Handle(err, "")
	}

	err = os.Chdir("..")
	base.Handle(err, "")
}

func downloadFile(url string) error {

	_, err := os.Stat(url)
	if !os.IsNotExist(err) {
		fmt.Printf("%s ya exite, omitiendo descarga\n", url)
		return nil
	}
	// Get the data
	resp, err := http.Get(mediaUrl + url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(url)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

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

	downloadFiles(r)
	log.Println("Hasta la prótsimaaaa")
}
