package main

import (
	"log"
	"os"
	"strings"
)

func createDirectory(name string) {
	name = strings.ReplaceAll(name, "/", "")
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = os.Mkdir(name, 0755)
		if err != nil {
			log.Panicf("cannot create folder with name %s\n", name)
		}
	}
}
