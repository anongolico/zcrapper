package main

import (
	"log"
	"time"

	"github.com/anongolico/ib/config"
	"github.com/anongolico/ib/helpers"
	"github.com/anongolico/ib/schemas"
)

var (
	FormatsToDownload []string
	Buffer            []byte

	TotalFilesToDownload int
)

func main() {
	id, err := ReadPostId()
	if err != nil {
		log.Panicf("cannot read post id, exiting program")
	}

	// reads the post's id
	post, err := schemas.ScanPost(id, Buffer)
	if err != nil {
		log.Panicf("cannot scan post")
	}

	// scans all the formats in the post
	Formats, numberOfFiles := helpers.ScanFormats(post)
	log.Printf("%d comments with attachments found", numberOfFiles)

	err = ReadFormats(Formats)
	if err != nil {
		log.Panicf("cannot scan formats")
	}

	// filter
	FilterFormats(Formats)

	// count
	CountFilesToDownload(Formats)

	// get files
	ts := time.Now()
	DownloadFiles(post, Formats)
	log.Printf("Hasta la prótsimaaaa (tiempo total: %v)\n", time.Since(ts))
}

func init() {
	err := config.ReadUrlParameters()
	if err != nil {
		log.Fatalf("cannot rad config file\n")
	}
	Buffer = make([]byte, config.BufferSize)
}
