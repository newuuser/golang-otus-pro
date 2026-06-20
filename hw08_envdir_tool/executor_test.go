package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		ini := []string{"go", "run", "./testdata/good_program/main.go"} //nolint
		code := RunCmd(ini, make(Environment))
		require.Zero(t, code)
	})
	t.Run("Invalid", func(t *testing.T) {
		ini := []string{"go", "run", "./testdata/failing_program/main.go"}
		code := RunCmd(ini, make(Environment))
		require.Equal(t, 1, code)
	})
	t.Run("Testing environment", func(t *testing.T) {
		ini := []string{"go", "run", "./testdata/testing_program/main.go", "arg1=1", "arg2=2"}
		env, _ := ReadDir("./testdata/env")
		code := RunCmd(ini, env)
		require.Equal(t, 0, code)
	})
}
