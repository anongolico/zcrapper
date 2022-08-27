package main

import (
	"fmt"
	"github.com/anongolico/base"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	AbsMainDirectoryPath string
)

const (
	Separator = "|"
)

func downloadFile(path string) error {
	values := strings.Split(path, Separator)
	folder := values[0]
	url := values[1]

	localFolder := AbsMainDirectoryPath + "/" + folder
	fileName := localFolder + "/" + url

	_, err := os.Stat(fileName)
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
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func downloadCover(path string) error {
	err := downloadFile("." + Separator + path)
	return err
}

func Laburante(jobs <-chan string, wg *sync.WaitGroup) {
	for job := range jobs {
		_ = downloadFile(job)
		TotalFilesToDownload--
		log.Printf("%d archivos restantes\n", TotalFilesToDownload)
		wg.Done()
	}
}

func downloadFiles(r *base.Rouz) {

	title := fmt.Sprintf("%s (%s)", r.Hilo.Titulo, r.Hilo.Id)
	title = strings.ReplaceAll(title, "/", "")
	title = strings.TrimSpace(title)
	createDirectory(title)
	err := os.Chdir(title)
	base.Handle(err, "Imposible entrar a la carpeta del roz, badre\n>hide")

	AbsMainDirectoryPath, _ = filepath.Abs(".")

	err = downloadCover(r.Hilo.Media.Url)
	base.Handle(err, "")

	for _, v := range FormatsToDownload {
		createDirectory(v)
	}

	urls := make(chan string, MaxParallelDownloads)
	wg := sync.WaitGroup{}
	for i := 0; i < MaxParallelDownloads; i++ {
		go Laburante(urls, &wg)
	}

	for _, format := range FormatsToDownload {
		for _, file := range base.Formats[format] {
			wg.Add(1)
			urls <- format + Separator + file
		}
	}

	wg.Wait()
	err = os.Chdir("..")
}
