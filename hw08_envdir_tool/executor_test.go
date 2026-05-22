package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		ini := []string{"ls"}
		code := RunCmd(ini, make(Environment))
		require.Zero(t, code)
	})
	/*t.Run("Invalid", func(t *testing.T) {
		ini := []string{"progrem"}
		code := RunCmd(ini, make(Environment))
		require.NotEqual(t, 0, code)
	})*/
}
