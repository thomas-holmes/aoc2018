package p01

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Run() {
	f, err := os.Open(filepath.Join("data", "1-1.txt"))
	if err != nil {
		log.Panicln("Failed to open file", err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panicln("Failed to read file", err)
	}

	buf := bufio.NewReader(bytes.NewBuffer(bs))

	var acc int

	for {
		lBytes, err := buf.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Panicln("Failed to read line", err)
		}

		i, err := strconv.Atoi(strings.Trim(string(lBytes), "\n"))
		if err != nil {
			log.Panicln("Could not convert string", string(lBytes), "to int", err)
		}

		acc += i
	}

	log.Println("Computed total freq of", acc)
}
