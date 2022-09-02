package ofd

import (
	"archive/zip"
	"bytes"
	"io"
)

func LoadZipFileContent(rc *zip.ReadCloser, path string) ([]byte, error) {
	filePath := path
	if filePath[0] == '/' {
		filePath = filePath[1:]
	}
	fs, err := rc.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, fs)
	if err != nil {
		return nil, err
	}
	content := buf.Bytes()
	return content, nil
}
