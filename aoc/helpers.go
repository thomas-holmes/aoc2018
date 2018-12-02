package aoc

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

// ReadStrings loads a file from the data directory and returns a slice of strings
func ReadStrings(fileName string) ([]string, error) {
	f, err := os.Open(filepath.Join("data", fileName))
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open file")
	}
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read file")
	}

	buf := bufio.NewReader(bytes.NewBuffer(bs))

	var data []string
	for {
		lBytes, err := buf.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, errors.Wrap(err, "Failed to read line")
		}

		data = append(data, strings.TrimSpace(string(lBytes)))
	}

	return data, nil
}
