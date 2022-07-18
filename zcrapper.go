package main

import (
	"fmt"
	"github.com/anongolico/base"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var mediaUrl = base.BaseMediaUrl

// createRouzFolder creates a new folder to store media
func createRouzFolder(name string) {
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
	for i, v := range base.Formats {
		createRouzFolder(i)
		err = os.Chdir(i)
		base.Handle(err, "")
		for _, x := range v {
			err = downloadFile(x)
			base.Handle(err, "error al descargar archivo")
		}
		err = os.Chdir("..")
		base.Handle(err, "")
	}

	/*for _, v := range base.Formats["webm"] {
		fmt.Printf("%d archivos restantes\n", l)
		// fmt.Println("v:", v)
		err = downloadFile(v)
		base.Handle(err, "Fallo al descargar el archivo")
		l--
	}*/

	err = os.Chdir("..")
	base.Handle(err, "")
}

func downloadFile(url string) error {
	_, fileName, _ := strings.Cut(url, mediaUrl)
	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) {
		fmt.Printf("%s ya exite, omitiendo descarga\n", fileName)
		return nil
	}
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// Create the file
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {
	base.ReadAuthFile()
	var id string
	fmt.Println("Ingresa el ID del rouz:")
	_, _ = fmt.Scan(&id)
	r := base.New(id)
	s := base.ScanFormats(r)
	fmt.Println(s)
	for i, v := range base.Formats {
		fmt.Printf("%s: %d\n", i, len(v))
	}
	downloadFiles(r)
	log.Println("lito")
}
