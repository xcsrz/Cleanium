package main

import (
	"io/ioutil"
	"os"
)

func userDir() string {
	tmpdir := os.Getenv("TMPDIR")
	if tmpdir == "" {
		tmpdir = "/tmp"
	}
	dir, err := ioutil.TempDir(tmpdir, "cleanium")
	if err != nil {
		bomb("Could not create a clean temporary directory.")
	}
	return dir
}
