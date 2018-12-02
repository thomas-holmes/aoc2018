package main

import (
	"github.com/thomas-holmes/aoc2018/aoc"

	"log"
)

func main() {
	p01()
	p02()
}

func p01() {
	boxIds, err := aoc.ReadStrings("2-1.txt")
	if err != nil {
		log.Panicln("Failed to run p01", err)
	}

	var twos, threes int

	for _, id := range boxIds {
		counts := make(map[rune]int)
		for _, c := range id {
			i := counts[c]
			counts[c] = i + 1
		}

		var twoInc, threeInc int
		for _, v := range counts {
			switch v {
			case 2:
				twoInc = 1
			case 3:
				threeInc = 1
			}
		}

		twos += twoInc
		threes += threeInc
	}

	log.Println("Checksum is", twos, "twos and", threes, "threes. Total =", twos*threes)
}

func p02() {
	boxIds, err := aoc.ReadStrings("2-1.txt")
	if err != nil {
		log.Panicln("Failed to run p02")
	}

	var closest int
	var matching string

	var sol1, sol2 string

	for j, id1 := range boxIds {
		for k, id2 := range boxIds {
			if j == k {
				continue
			}

			rs1 := []rune(id1)
			rs2 := []rune(id2)

			var same int

			var match []rune
			for i := 0; i < len(rs1); i++ {
				if rs1[i] == rs2[i] {
					same++
					match = append(match, rs1[i])
				}
			}

			if same > closest {
				closest = same
				sol1, sol2 = id1, id2
				matching = string(match)
			}
		}
	}

	log.Println("found solution with length", closest, "and ids", sol1, sol2, "same values", matching)
}
