package main

import (
	"log"
	"os"
)

var (
	DEBUG    = false
	APP_NAME = "Cleanium"
)

func init() {
	if os.Getenv("DEBUG") != "" {
		DEBUG = true
	}
	if os.Getenv("APP_NAME") != "" {
		APP_NAME = os.Getenv("APP_NAME")
	}
}

func buildEnv() (env []string) {
	want := []string{"PATH", "SHELL", "TMPDIR", "Apple_PubSub_Socket_Render", "USER", "__CF_USER_TEXT_ENCODING", "PWD", "LANG", "XPC_FLAGS", "XPC_SERVICE_NAME", "HOME", "LOGNAME"}
	for _, key := range want {
		env = append(env, key+"="+os.Getenv(key))
	}
	return
}

func main() {
	open()
}

func bomb(message string) {
	log.Fatalln("ERROR IN ", APP_NAME+":", message)
	os.Exit(1)
}
