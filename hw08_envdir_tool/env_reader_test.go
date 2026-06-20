package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		expected := map[string]EnvValue{
			"BAR":   EnvValue{"bar", false}, //nolint
			"EMPTY": EnvValue{"", false},
			"FOO":   EnvValue{"   foo\nwith new line", false},
			"HELLO": EnvValue{"\"hello\"", false},
			"UNSET": EnvValue{"", true},
		}
		env, ok := ReadDir("./testdata/env")
		require.Nil(t, ok)
		for k, v := range env {
			require.Equal(t, expected[k], v)
		}
	})
	t.Run("Invalid", func(t *testing.T) {
		_, ok := ReadDir("./not/existing/dir")
		require.NotNil(t, ok)
	})
	t.Run("Pass file", func(t *testing.T) {
		_, ok := ReadDir("./executor.go")
		require.NotNil(t, ok)
	})
}
