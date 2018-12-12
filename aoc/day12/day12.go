package main

import (
	"log"
	"time"

	"github.com/thomas-holmes/aoc2018/aoc"
)

func main() {
	t0 := time.Now()
	//p01()
	log.Println("P01:", time.Since(t0))
	t1 := time.Now()
	p02()
	log.Println("P02:", time.Since(t1))
}

func p01() {
	//lines, err := aoc.ReadStrings("12.txt")
	lines, err := aoc.ReadStrings("12.txt")
	if err != nil {
		log.Panicln("Failed to read file", err)
	}

	start := lines[0]
	var rules [][]bool
	for _, l := range lines[2:] {
		rule := make([]bool, 6)
		for i, r := range l {
			if i > 4 {
				break
			}

			rule[i] = r == '#'
		}

		rule[5] = []rune(l)[len(l)-1] == '#'

		if err != nil {
			log.Panicln("failed to scan", err)
		}
		rules = append(rules, rule)
	}

	gen := make([]bool, len(start)+10)
	for i, r := range start {
		gen[i+5] = r == '#'
	}

	const MaxGens = 20

	firstPot := -5

	printGen(gen)
	for g := 0; g < MaxGens; g++ {
		/*
			nextGen := make([]bool, len(gen)+4)
			for i := 0; i < len(gen); i++ {
				nextGen[i+2] = gen[i]
			}
		*/
		nextGen := make([]bool, len(gen))

		for i := 2; i < len(gen)-2; i++ {
			for _, r := range rules {
				max := i + 3
				if alive(r, gen[i-2:max]) {
					nextGen[i] = true
					continue
				}
			}
		}

		{
			for i := 0; i < 3; i++ {
				if nextGen[i] || nextGen[len(gen)-(i+1)] {
					nextGen = append([]bool{false, false, false}, nextGen...)
					nextGen = append(nextGen, false, false, false)
					firstPot -= 3
					break
				}
			}
		}

		gen = nextGen
		printGen(nextGen)
	}

	var total int
	curPot := firstPot
	for _, p := range gen {
		if p {
			total += curPot
		}
		curPot++
	}

	log.Println("Total of", total, "First Pot", firstPot)
}

func printGen(g []bool) {
	var out string
	for _, b := range g {
		if b {
			out += "#"
		} else {
			out += "."
		}
	}
	log.Println(out)
}

func alive(rule, candidate []bool) bool {
	if rule[0] == candidate[0] &&
		rule[1] == candidate[1] &&
		rule[2] == candidate[2] &&
		rule[3] == candidate[3] &&
		rule[4] == candidate[4] {
		return rule[5]
	}
	return false
}

func p02() {
	lines, err := aoc.ReadStrings("12.txt")
	if err != nil {
		log.Panicln("Failed to read file", err)
	}

	start := lines[0]
	var rules [][]bool
	for _, l := range lines[2:] {
		rule := make([]bool, 6)
		for i, r := range l {
			if i > 4 {
				break
			}

			rule[i] = r == '#'
		}

		rule[5] = []rune(l)[len(l)-1] == '#'

		if err != nil {
			log.Panicln("failed to scan", err)
		}
		rules = append(rules, rule)
	}

	gen := make([]bool, len(start)+10)
	for i, r := range start {
		gen[i+5] = r == '#'
	}

	const MaxGens = 50000000000

	firstPot := -5

	printGen(gen)
	for g := 0; g < MaxGens; g++ {
		nextGen := make([]bool, len(gen))
		if g%1000000 == 0 {
			log.Println(g)
		}

		for i := 2; i < len(gen)-2; i++ {
			for _, r := range rules {
				max := i + 3
				if alive(r, gen[i-2:max]) {
					nextGen[i] = true
					continue
				}
			}
		}

		{
			var firstPlant, lastPlant int

			for i := 0; i < len(nextGen); i++ {
				if nextGen[i] {
					firstPlant = i
					break
				}
			}

			for i := len(nextGen) - 1; i > 0; i-- {
				if nextGen[i] {
					lastPlant = i
					break
				}
			}

			// I'm going to cheat and hardcode the case I know exists, this grows forever to the right
			// log.Printf("len(nextGen)=%d lastPlant=%d sub=%d", len(nextGen), lastPlant, len(nextGen)-lastPlant)
			if len(nextGen)-lastPlant < 4 {
				needs := 4 - (len(nextGen) - lastPlant)

				// log.Printf("first %d, last %d, needs %d", firstPlant, lastPlant, needs)
				if firstPlant-2 > needs {
					firstPot += needs
					for i := 0; i < len(nextGen)-needs; i++ {
						nextGen[i] = nextGen[i+needs]
						nextGen[i+needs] = false
					}
				} else {
					nextGen = append([]bool{false, false, false, false}, nextGen...)
					nextGen = append(nextGen, false, false, false, false)
					firstPot -= 4
				}
			}
		}

		// printGen(nextGen)

		gen = nextGen

	}

	var total int
	curPot := firstPot
	for _, p := range gen {
		if p {
			total += curPot
		}
		curPot++
	}

	log.Println("Total of", total, "First Pot", firstPot)
}
