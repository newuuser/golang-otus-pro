package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
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

	out, err := os.Create(toPath)
	if err != nil {
		return err
	}
	infoOut, err := out.Stat()
	if err != nil {
		return err
	}
	if infoOut.IsDir() {
		return ErrUnsupportedFile
	}

	defer out.Close()
	_, err = in.Seek(offset, 0)
	if err != nil {
		return err
	}

	if limit == 0 {
		limit = info.Size()
	}

	limit = min(limit, info.Size()-offset)

	bar := pb.StartNew(1000)
	defer bar.Finish()
	barReader := bar.NewProxyReader(in)

	_, err = io.CopyN(out, barReader, limit)
	if err != nil {
		return err
	}

	return nil
}
