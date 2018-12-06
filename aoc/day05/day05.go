package main

import (
	"log"
	"strings"
	"time"

	"github.com/thomas-holmes/aoc2018/aoc"
)

func main() {
	t0 := time.Now()
	p01()
	log.Println("p01", time.Since(t0))

	t1 := time.Now()
	p01stack()
	log.Println("p01-stack", time.Since(t1))

	t2 := time.Now()
	p02()
	log.Println("p02", time.Since(t2))

	t3 := time.Now()
	p02stack()
	log.Println("p02-stack", time.Since(t3))
}

const offset = 'a' - 'A'

func pair(a, b rune) bool {
	return (a+offset) == b || (a-offset) == b
}

func p01() {
	lines, err := aoc.ReadStrings("5.txt")
	if err != nil {
		log.Panicln("Failed to read text", err)
	}

	log.Println("p01", len(collapse(lines[0])))
}

func p02() {
	lines, err := aoc.ReadStrings("5.txt")
	if err != nil {
		log.Panicln("Failed to read text", err)
	}

	polymerChain := []rune(lines[0])

	shortest := len(polymerChain)

	polymer := lines[0]
	for c := rune('A'); c < rune('A')+26; c++ {
		candidate := strings.Replace(polymer, string([]rune{c}), "", -1)
		candidate = strings.Replace(candidate, string([]rune{c + offset}), "", -1)
		length := len(collapse(candidate))
		if length < shortest {
			shortest = length
		}
	}

	log.Println("p02:", shortest)
}

func collapse(polymer string) string {
	var adjusted bool
	var last rune
	var i int

	polymerChain := []rune(polymer)
	for {
		if len(polymerChain) < 2 {
			break
		}

		this := polymerChain[i]

		if pair(last, this) {
			adjusted = true
			if i-1 < 0 || i+1 > len(polymerChain) {
				log.Println(i, len(polymerChain), string(polymerChain))
			}
			polymerChain = append(polymerChain[:i-1], polymerChain[i+1:]...)
			last = 0
			this = 0
		}

		last = this
		i++

		if i >= len(polymerChain) {
			if !adjusted {
				break
			}
			adjusted = false
			last = 0
			i = 0
		}
	}

	return string(polymerChain)
}

func p01stack() {
	lines, err := aoc.ReadStrings("5.txt")
	if err != nil {
		log.Panicln("Failed to read text", err)
	}

	line := lines[0]
	s := newStack(len(line))

	for _, r := range line {
		f, ok := s.peek()
		if !ok {
			s.push(r)
			continue
		}

		if pair(r, f) {
			_ = s.pop() // throw it away
			continue
		}

		s.push(r)
	}

	log.Println("p01stack:", s.len())

}

func p02stack() {
	lines, err := aoc.ReadStrings("5.txt")
	if err != nil {
		log.Panicln("Failed to read text", err)
	}

	polymerChain := []rune(lines[0])

	shortest := len(polymerChain)

	polymer := lines[0]
	for c := rune('A'); c < rune('A')+26; c++ {
		candidate := strings.Replace(polymer, string([]rune{c}), "", -1)
		candidate = strings.Replace(candidate, string([]rune{c + offset}), "", -1)
		length := stackCollapse(candidate)
		if length < shortest {
			shortest = length
		}
	}

	log.Println("p02stack", shortest)
}

// new stack stuff
type stack struct {
	stack []rune
	ptr   int
}

func (s *stack) push(r rune) {
	s.ptr++
	if s.ptr > len(s.stack) {
		log.Panicln("push failure", s.ptr, ">", len(s.stack))
	}

	s.stack[s.ptr] = r
}

func (s *stack) pop() rune {
	if s.ptr < 0 {
		log.Panicln("pop off empty stack")
	}
	var erase rune
	r := s.stack[s.ptr]
	s.stack[s.ptr] = erase
	s.ptr--

	return r
}

func (s *stack) peek() (rune, bool) {
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
		stack: make([]rune, l, l),
		ptr:   -1,
	}
}

func stackCollapse(line string) int {
	s := newStack(len(line))

	for _, r := range line {
		f, ok := s.peek()
		if !ok {
			s.push(r)
			continue
		}

		if pair(r, f) {
			_ = s.pop() // throw it away
			continue
		}

		s.push(r)
	}

	return s.len()
}
