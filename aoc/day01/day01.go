package main

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
	"time"
)

func main() {
	p1()
	p2()
}

func p1() {
	f, err := os.Open(filepath.Join("data", "1-2.txt"))
	if err != nil {
		log.Panicln("Failed to open file", err)
	}
	defer f.Close()

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

		numbers = append(numbers, i)
	}

	var acc int
	for _, i := range numbers {
		acc += i
	}

	log.Println("P1: Computed total freq of", acc)
}

func p2() {
	t1 := time.Now()
	f, err := os.Open(filepath.Join("data", "1-2.txt"))
	if err != nil {
		log.Panicln("Failed to open file", err)
	}
	defer f.Close()

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

		numbers = append(numbers, i)
	}

	result := scan(0, numbers, make(map[int]struct{}))

	log.Println(time.Since(t1))

	log.Println("P2: Found duplicate frequency", result)
}

func scan(acc int, numbers []int, seen map[int]struct{}) int {
	for _, i := range numbers {
		acc += i
		if _, ok := seen[acc]; ok {
			return acc
		}
		seen[acc] = struct{}{}
	}
	return scan(acc, numbers, seen)
}
