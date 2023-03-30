package main

import (
	"log"
	"os"
)

var (
	DEBUG = false
)

func init() {
	if os.Getenv("DEBUG") != "" {
		DEBUG = true
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
	log.Fatalln("ERROR IN CLEANIUM:", message)
	os.Exit(1)
}
