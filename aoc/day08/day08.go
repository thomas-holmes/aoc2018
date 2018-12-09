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
	t1 := time.Now()
	p02()
	log.Println("P02:", time.Since(t1))
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

func p02() {
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
	var sumMD func(start int) (int, int)
	// Forgive me, this is horrific, but it works.
	// Returns (consumed, nodeValue int)
	sumMD = func(start int) (int, int) {
		consumed := 0
		numChildren := numbers[start]
		numMetadata := numbers[start+1]
		consumed += 2

		var childValues []int
		for i := 0; i < numChildren; i++ {
			c, value := sumMD(start + consumed)
			consumed += c
			childValues = append(childValues, value)
		}

		var mdSum int
		var mdEntries []int
		for i := 0; i < numMetadata; i++ {
			mdNum := numbers[start+consumed]
			md = append(md, mdNum)
			mdEntries = append(mdEntries, mdNum)
			sum += mdNum
			mdSum += mdNum
			consumed++
		}

		if numChildren == 0 {
			return consumed, mdSum
		}
		var nodeTotal int
		for _, ni := range mdEntries {
			if ni > 0 && ni <= numChildren {
				// log.Printf("NumChildren=%d len(childValues)=%d Looking for nodeIndex %d in childValues: %v", numChildren, len(childValues), ni-1, childValues)
				nodeTotal += childValues[ni-1]
			}
		}
		return consumed, nodeTotal
	}

	consumed, total := sumMD(0)

	log.Printf("P02: Consumed %d, Total %d", consumed, total)
}
