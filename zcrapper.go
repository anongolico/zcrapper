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
	"sync"
	"time"
)

var (
	mediaUrl             = base.BaseMediaUrl
	FormatsToDownload    []string
	TotalFilesToDownload int
	Id                   string
)

const MaxParallelDownloads = 50

// createRouzFolder creates a new folder to store media
func createRouzFolder(name string) {
	// algunos dawns ponen el slash en el nombre del roz y
	// eso genera error si el SO es linux
	name = strings.ReplaceAll(name, "/", "")
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, " ", "_")
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = os.Mkdir(name, 0755)
		base.Handle(err, "")
	}
}

func downloadFiles(r *base.Rouz) {
	folderName := fmt.Sprintf("%s_%s", r.Hilo.Id, r.Hilo.Titulo)
	createRouzFolder(folderName)
	err := os.Chdir(folderName)
	base.Handle(err, "")

	// channel para descargas en paralelo
	openConns := make(chan int, MaxParallelDownloads)
	var wg sync.WaitGroup

	// err = downloadFile(r.Hilo.Media.Url)

	for _, v := range FormatsToDownload {
		createRouzFolder(v)
		err = os.Chdir(v)
		base.Handle(err, "")

		for i, x := range base.Formats[v] {
			openConns <- 1
			wg.Add(1)
			// fmt.Printf("Sending file %d with URL %s\n", i, x)
			// go func() {
			go downloadFile(x, openConns, &wg, i)
			// base.Handle(err, "error al descargar archivo")

			// }()
		}
		wg.Wait()
		err = os.Chdir("..")
		base.Handle(err, "")
	}

	err = os.Chdir("..")
	base.Handle(err, "")
}

func downloadFile(url string, c chan int, wg *sync.WaitGroup, index int) {

	defer wg.Done()
	_, err := os.Stat(url)
	if !os.IsNotExist(err) {
		TotalFilesToDownload--
		log.Printf("%s ya exite, omitiendo descarga (%d restantes)\n", url, TotalFilesToDownload)
		return
	}
	// Create the file
	out, err := os.Create(url)
	if err != nil {
		return
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(mediaUrl + url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	<-c

	TotalFilesToDownload--
	log.Printf("%d restantes\n", TotalFilesToDownload+1)
}

// promptUser() asks the anon the proper questions
// to start the download process.
func promptUser() {
	f := make([]string, len(base.Formats))
	i := 0
	for k := range base.Formats {
		f[i] = k
		i++
	}

	q := []*survey.Question{
		{
			Name:     "id",
			Prompt:   &survey.Input{Message: "Ingresa el ID del rouz:"},
			Validate: survey.Required,
		},
		{
			Name: "formats",
			Prompt: &survey.MultiSelect{
				Message: "Escoge los formatos a descargar:",
				Options: f,
			},
			Validate: survey.Required,
		},
	}

	answers := struct {
		id      string   `survey:"id"`      // or you can tag fields to match a specific name
		formats []string `survey:"formats"` // if the types don't match, survey will convert it
	}{
		Id,
		FormatsToDownload,
	}

	_ = survey.Ask(q, &answers)
}

func main() {
	base.ReadAuthFile()

	var id string

	questions := []*survey.Question{
		{
			Name:     "id",
			Prompt:   &survey.Input{Message: "Ingresa el ID del rouz:"},
			Validate: survey.Required,
		},
	}
	_ = survey.Ask(questions, &id)
	r := base.New(id)
	_ = base.ScanFormats(r)

	promptUser()

	timestamp := time.Now()

	for _, v := range FormatsToDownload {
		TotalFilesToDownload += len(base.Formats[v])
	}

	log.Printf("%d archivos para descargar. Puede tomar un tiempo.", TotalFilesToDownload+1)

	downloadFiles(r)
	log.Printf("Hasta la prótsimaaaa (%v)", time.Since(timestamp))
}
