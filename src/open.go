package main

import (
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"syscall"
)

func open() {
	switch runtime.GOOS {
	case "darwin":
		open_macos()
	case "windows":
		log.Fatalln("Windows is not yet supported")
	case "linux":
		log.Fatalln("Linux is not yet supported")
	default:
		log.Fatalln("Unknown operating system:", runtime.GOOS)
	}

}

func open_macos() {
	var userdir string
	if os.Getenv("BROWSER_USER_DIR") == "" {
		// if this is the first execution from the click/app open then we re-open with the temp
		// directory as an arg allowing us to request a new instance so each opening results
		// in a new instance of the application with it's own respective data directory

		userdir := userDir()
		appPath := strings.Split(os.Args[0], string(os.PathSeparator))
		for {
			if len(appPath) < 1 {
				bomb("could not find app path")
				return
			}
			if strings.HasSuffix(appPath[len(appPath)-1], ".app") {
				break
			}
			appPath = appPath[:len(appPath)-1]
		}
		err := syscall.Exec("/usr/bin/open", []string{"-a", path.Join(appPath...), "--new"}, append(buildEnv(), "BROWSER_USER_DIR="+userdir))
		if err != nil {
			bomb("Could not reopen " + APP_NAME + ": " + err.Error())
		}
	} else {
		// now that we have a fresh new instance, let's exec to chrome/chromium
		userdir = os.Getenv("BROWSER_USER_DIR")
		path := findBrowserPath()
		args := []string{path, "--user-data-dir=" + userdir, "--no-first-run", "--no-default-browser-check"}

		args = append(args, os.Args[1:]...)

		if os.Getenv("APP_URL") != "" {
			args = append(args, os.Getenv("APP_URL"))
		}

		err := syscall.Exec(path, args, buildEnv())
		if err != nil {
			bomb("Could not launch browser.")
		}
	}
}
