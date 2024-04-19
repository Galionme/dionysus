package util

import (
	"bufio"
	"bytes"
	"os"
)

func PreProcessImages(dat *os.File) (*bytes.Reader, error) {
	stat, err := dat.Stat()
	if err != nil {
		return nil, err
	}

	var size = stat.Size()
	b := make([]byte, size)

	bufR := bufio.NewReader(dat)
	_, err = bufR.Read(b)
	bReader := bytes.NewReader(b)

	return bReader, err
}
