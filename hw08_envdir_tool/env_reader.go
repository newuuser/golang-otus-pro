package main

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"
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
		ok := true
		// Skip if name contains '='
		for _, c := range name {
			if c == '=' {
				ok = false
				break
			}
		}

		if !ok {
			continue
		}

		path := filepath.Join(dir, name)
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		// Check if empty
		stat, err := f.Stat()

		if err != nil {
			return nil, err
		}

		if stat.Size() == 0 {
			env[name] = EnvValue{"", true}
			continue
		}

		br := bufio.NewScanner(f)
		line := ""
		ok = br.Scan()
		if ok {
			line = br.Text()
		}

		line = strings.TrimRight(line, "\t ")
		line = string(bytes.Replace([]byte(line), []byte("\000"), []byte("\n"), -1))
		env[name] = EnvValue{line, false}
	}
	return env, nil
}
