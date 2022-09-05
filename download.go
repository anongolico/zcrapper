package main

import (
	"fmt"
	"github.com/anongolico/ib/config"
	"github.com/anongolico/ib/schemas"
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
	FilePathSeparator    string
)

const (
	Separator = "|"
)

func downloadFile(path string) error {
	values := strings.Split(path, Separator)
	folder := values[0]
	url := values[1]

	localFolder := AbsMainDirectoryPath + FilePathSeparator + folder
	fileName := localFolder + FilePathSeparator + url

	_, err := os.Stat(fileName)
	if !os.IsNotExist(err) {
		fmt.Printf("%s ya exite, omitiendo descarga\n", url)
		return nil
	}
	// Get the data
	resp, err := http.Get(config.MediaUrl + url)
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

func Worker(jobs <-chan string, wg *sync.WaitGroup) {
	for job := range jobs {
		err := downloadFile(job)
		if err != nil {
			fmt.Println("here's the error")
			fmt.Println(err.Error())
		}
		TotalFilesToDownload--
		log.Printf("%d archivos restantes\n", TotalFilesToDownload)
		wg.Done()
	}
}

func DownloadFiles(post *schemas.Post, formats map[string][]string) {

	title := fmt.Sprintf("%s (%s)", post.Hilo.Titulo, post.Hilo.Id)
	title = strings.ReplaceAll(title, "/", "")
	title = strings.ReplaceAll(title, FilePathSeparator, "")
	title = strings.TrimSpace(title)
	createDirectory(title)
	err := os.Chdir(title)
	if err != nil {
		log.Panicf("")
	}

	AbsMainDirectoryPath, _ = filepath.Abs(".")

	if !strings.Contains(post.Hilo.Media.Url, "//") {
		err = downloadCover(post.Hilo.Media.Url)
		if err != nil {
			log.Panicf("no puedo descargar la portada")
		}
	}

	for _, v := range FormatsToDownload {
		createDirectory(v)
	}

	urls := make(chan string, config.MaxParallelDownloads)
	wg := sync.WaitGroup{}

	// initialize workers routines
	for i := 0; i < config.MaxParallelDownloads; i++ {
		go Worker(urls, &wg)
	}

	for format, v := range formats {
		for _, file := range v {
			wg.Add(1)
			urls <- format + Separator + file
		}
	}

	wg.Wait()
	err = os.Chdir("..")
}

func CountFilesToDownload(formats map[string][]string) {
	for _, v := range formats {
		TotalFilesToDownload += len(v)
	}
}

func init() {
	FilePathSeparator = string(filepath.Separator)
}
