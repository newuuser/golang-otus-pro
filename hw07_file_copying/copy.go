package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

const (
	bufferSize = 2048
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	//os.FileMode
	in, err := os.OpenFile(fromPath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer in.Close()
	info, err := in.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 || info.IsDir() {
		return ErrUnsupportedFile
	} else if offset > info.Size() {
		return ErrOffsetExceedsFileSize
	}

	//out, err := os.OpenFile(toPath, os.O_WRONLY, os.ModePerm)
	out, err := os.Create(toPath)
	/*if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			out, err = os.Create(toPath)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}*/
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = in.Seek(offset, 0)
	if err != nil {
		return nil
	}

	if limit == 0 {
		limit = info.Size()
	}

	limit = min(limit, info.Size())
	_, err = io.CopyN(out, in, limit)
	if err != nil {
		return err
	}

	return nil
}
