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
	f, err := os.Open(filepath.Join("data", "1-2.txt"))
	if err != nil {
		log.Panicln("Failed to open file", err)
	}

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		log.Panicln("Failed to read file", err)
	}

	buf := bufio.NewReader(bytes.NewBuffer(bs))

	var numbers []int
	for {
		lBytes, err := buf.ReadBytes('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Panicln("Failed to read line", err)
		}

		i, err := strconv.Atoi(strings.Trim(string(lBytes), "\r\n"))
		if err != nil {
			log.Panicln("Could not convert string", string(lBytes), "to int", err)
		}

		log.Println("Number:", i)
		numbers = append(numbers, i)
	}

	result := scan(0, numbers, make(map[int]struct{}))

	log.Println("Computed total freq of", result)
}

func scan(acc int, numbers []int, seen map[int]struct{}) int {
	for _, i := range numbers {
		acc += i
		log.Println("acc", acc)
		if _, ok := seen[acc]; ok {
			log.Println("Found duplicate frequency", acc)
			return acc
		}
		seen[acc] = struct{}{}
	}
	return scan(acc, numbers, seen)
}
