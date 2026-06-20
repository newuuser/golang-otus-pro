package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	t.Run("file doesn't exist", func(t *testing.T) {
		err := Copy("./testdata/not_real.txt", "./testdata/cpy", 0, 0)
		require.NotNil(t, err)
	})
	t.Run("file is directory", func(t *testing.T) {
		err := Copy("testdata", "./testdata/cpy", 0, 0)
		require.Equal(t, ErrUnsupportedFile, err)
	})

	t.Run("offset bigger than file size", func(t *testing.T) {
		err := Copy("./testdata/input.txt", "./testdata/cpy.txt", 1000000, 10)
		require.Equal(t, ErrOffsetExceedsFileSize, err)
	})

	t.Run("offset=0, limit=10", func(t *testing.T) {
		os.Create(".testdata/cpy.txt")
		defer os.Remove("testdata/cpy.txt")
		err := Copy("./testdata/input.txt", "./testdata/cpy.txt", 0, 10)
		require.Nil(t, err)

		file1, err := os.ReadFile("./testdata/cpy.txt")
		require.Nil(t, err)
		file2, err := os.ReadFile("./testdata/out_offset0_limit10.txt")
		require.Nil(t, err)
		require.True(t, bytes.Equal(file1, file2))
	})

	t.Run("offset=0, limit=0", func(t *testing.T) {
		os.Create(".testdata/cpy.txt")
		defer os.Remove("testdata/cpy.txt")
		err := Copy("./testdata/input.txt", "./testdata/cpy.txt", 0, 0)
		require.Nil(t, err)

		file1, err := os.ReadFile("./testdata/cpy.txt")
		require.Nil(t, err)
		file2, err := os.ReadFile("./testdata/out_offset0_limit0.txt")
		require.Nil(t, err)
		require.True(t, bytes.Equal(file1, file2))
	})

	t.Run("offset=0, limit=1000", func(t *testing.T) {
		os.Create(".testdata/cpy.txt")
		defer os.Remove("testdata/cpy.txt")
		err := Copy("./testdata/input.txt", "./testdata/cpy.txt", 0, 1000)
		require.Nil(t, err)

		file1, err := os.ReadFile("./testdata/cpy.txt")
		require.Nil(t, err)
		file2, err := os.ReadFile("./testdata/out_offset0_limit1000.txt")
		require.Nil(t, err)
		require.True(t, bytes.Equal(file1, file2))
	})

	t.Run("offset=100, limit=1000", func(t *testing.T) {
		os.Create(".testdata/cpy.txt")
		defer os.Remove("testdata/cpy.txt")
		err := Copy("./testdata/input.txt", "./testdata/cpy.txt", 100, 1000)
		require.Nil(t, err)

		file1, err := os.ReadFile("./testdata/cpy.txt")
		require.Nil(t, err)
		file2, err := os.ReadFile("./testdata/out_offset100_limit1000.txt")
		require.Nil(t, err)
		require.True(t, bytes.Equal(file1, file2))
	})

	t.Run("offset=6000, limit=1000", func(t *testing.T) {
		os.Create(".testdata/cpy.txt")
		defer os.Remove("testdata/cpy.txt")
		err := Copy("./testdata/input.txt", "./testdata/cpy.txt", 6000, 1000)
		require.Nil(t, err)

		file1, err := os.ReadFile("./testdata/cpy.txt")
		require.Nil(t, err)
		file2, err := os.ReadFile("./testdata/out_offset6000_limit1000.txt")
		require.Nil(t, err)
		require.True(t, bytes.Equal(file1, file2))
	})

	t.Run("copy file to self", func(t *testing.T) {
		f, _ := os.Create(".testdata/test.txt")
		defer os.Remove("testdata/test.txt")
		f.WriteString("test")
		f.Close()

		err := Copy("test.txt", "test.txt", 0, 0)
		require.NotNil(t, err)
	})
}
