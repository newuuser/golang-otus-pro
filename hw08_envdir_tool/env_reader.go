package main

import (
	"bufio"
	"os"
	"path/filepath"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)
	for _, file := range files {
		name := file.Name()
		path := filepath.Join(dir, name)
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		br := bufio.NewScanner(f)
		line := ""
		ok := br.Scan()
		if ok {
			line = br.Text()
		}

		env[name] = EnvValue{line, ok}
		//os.ReadFile(path)
	}
	return env, nil
}
