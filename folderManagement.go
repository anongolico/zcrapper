package main

import (
	"github.com/anongolico/base"
	"os"
	"strings"
)

func createDirectory(name string) {
	name = strings.ReplaceAll(name, "/", "")
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		err = os.Mkdir(name, 0755)
		base.Handle(err, "")
	}
}
