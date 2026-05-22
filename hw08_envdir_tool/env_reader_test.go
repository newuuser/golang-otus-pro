package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDir(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		env, ok := ReadDir("./testdata/env")
		require.Nil(t, ok)
		for k, v := range env {
			fmt.Println(k, v.Value, v.NeedRemove)
		}
	})
}
