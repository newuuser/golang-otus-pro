package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args

	env, err := ReadDir(args[1])
	if err != nil {
		log.Fatal(err)
	}
	ret := RunCmd(args[2:], env)
	if ret != 0 {
		log.Printf("Program exited with exit code %d", ret)
	}
	os.Exit(ret)
}
