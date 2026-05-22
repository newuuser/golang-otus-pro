package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		//ini := []string{"ls"}
		ini := []string{"./testdata/good_program/good_program.exe"}
		code := RunCmd(ini, make(Environment))
		require.Zero(t, code)
	})
	t.Run("Invalid", func(t *testing.T) {
		ini := []string{"./testdata/failing_program/failing_program.exe"}
		code := RunCmd(ini, make(Environment))
		require.Equal(t, 1, code)
	})
	t.Run("Testing environment", func(t *testing.T) {
		ini := []string{"./testdata/testing_program/testing_program.exe", "arg1=1", "arg2=2"}
		env, _ := ReadDir("./testdata/env")
		code := RunCmd(ini, env)
		require.Equal(t, 0, code)
	})
}
