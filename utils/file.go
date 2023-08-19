package util

import (
	"io"
	"mime/multipart"
)

func ReadFile(file *multipart.FileHeader) (content []byte, err error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}