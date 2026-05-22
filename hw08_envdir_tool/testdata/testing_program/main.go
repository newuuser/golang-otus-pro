package main

import "os"

func main() {
	expectedArgs := []string{
		"arg1=1",
		"arg2=2",
	}

	args := os.Args
	expected := map[string]string{
		"BAR":   "bar", //nolint
		"EMPTY": "",
		"FOO":   "   foo\nwith new line",
		"HELLO": "\"hello\"",
	}
	for k, v := range expected {
		if os.Getenv(k) != v {
			os.Exit(1)
		}
	}
	_, ok := os.LookupEnv("UNSET")
	if ok {
		os.Exit(1)
	}

	if len(args) != len(expectedArgs)+1 {
		os.Exit(1)
	}

	for i := 1; i < len(args); i++ {
		if args[i] != expectedArgs[i-1] {
			os.Exit(1)
		}
	}
}
