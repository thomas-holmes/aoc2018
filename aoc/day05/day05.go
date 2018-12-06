package main

import (
	"log"

	"github.com/thomas-holmes/aoc2018/aoc"
)

func main() {
	p01()
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

	polymerChain := []rune(lines[0])

	var adjusted bool
	var last rune
	var i int

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

	log.Println(len(polymerChain))
}
