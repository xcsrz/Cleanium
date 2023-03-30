package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func findBrowserPath() string {
	var potentialPaths []string

	switch runtime.GOOS {
	case "darwin":
		dirPaths := []string{
			"/Applications",
			os.Getenv("HOME") + "/Applications",
		}
		appNames := []string{
			"Chromium",
			"Chrome",
			"Google Chrome",
		}
		for _, dir := range dirPaths {
			for _, app := range appNames {
				potentialPaths = append(potentialPaths, fmt.Sprintf("%s/%s.app/Contents/MacOS/%s", dir, app, app))
			}
		}
	case "windows":
		log.Fatalln("Windows is not yet supported")
	case "linux":
		log.Fatalln("Linux is not yet supported")
	default:
		log.Fatalln("Unknown operating system:", runtime.GOOS)
	}

	if DEBUG {
		fmt.Println("potentialPaths:")
		for _, path := range potentialPaths {
			fmt.Println("  -", path)
		}
		fmt.Println()
	}

	for _, path := range potentialPaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}
	bomb("Could not find a browser, this only works if you have Chrome or Chromium installed in your /Applications directory.")
	return ""
}
