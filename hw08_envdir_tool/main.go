package main

import (
	"log"
	"os"
)

// Failed to read directory : prints error and return code=2.
// Failed to run : prints error and return code=3.
func main() {
	args := os.Args

	env, err := ReadDir(args[1])
	if err != nil {
		log.Print(err)
		os.Exit(2)
	}
	ret := RunCmd(args[2:], env)
	if ret != 0 {
		log.Printf("Program exited with exit code %d", ret)
	}
	os.Exit(ret)
}
