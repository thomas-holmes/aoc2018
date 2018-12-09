package main

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/thomas-holmes/aoc2018/aoc"
)

func main() {
	t0 := time.Now()
	p01()
	log.Println("P01:", time.Since(t0))
}

func p01() {
	lines, err := aoc.ReadStrings("8.txt")
	if err != nil {
		log.Panicln("Failed to read data", err)
	}

	data := strings.NewReader(lines[0])
	var numbers []int
	for {
		var d int
		_, err := fmt.Fscanf(data, "%d", &d)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Panicln("Failed to scan file")
		}

		numbers = append(numbers, d)
	}

	var md []int
	var sum int
	var sumMD func(start int) int
	sumMD = func(start int) int {
		consumed := 0
		numChildren := numbers[start]
		numMetadata := numbers[start+1]
		consumed += 2

		for i := 0; i < numChildren; i++ {
			consumed += sumMD(start + consumed)
		}

		for i := 0; i < numMetadata; i++ {
			mdNum := numbers[start+consumed]
			md = append(md, mdNum)
			sum += mdNum
			consumed++
		}
		return consumed
	}

	consumed := sumMD(0)

	log.Printf("P01: Consumed %d, Total %d", consumed, sum)
}

// new stack stuff
type stack struct {
	stack []int
	ptr   int
}

func (s *stack) push(r int) {
	s.ptr++
	if s.ptr > len(s.stack) {
		log.Panicln("push failure", s.ptr, ">", len(s.stack))
	}

	s.stack[s.ptr] = r
}

func (s *stack) pop() int {
	if s.ptr < 0 {
		log.Panicln("pop off empty stack")
	}
	var erase int
	r := s.stack[s.ptr]
	s.stack[s.ptr] = erase
	s.ptr--

	return r
}

func (s *stack) peek() (int, bool) {
	if s.ptr < 0 {
		return 0, false
	}

	return s.stack[s.ptr], true
}

func (s *stack) len() int {
	return s.ptr + 1
}

func newStack(l int) *stack {
	return &stack{
		stack: make([]int, l, l),
		ptr:   -1,
	}
}
