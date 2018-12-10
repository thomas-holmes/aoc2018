package main

import (
	"fmt"
	"log"
	"math"
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
	lines, err := aoc.ReadStrings("10.txt")
	if err != nil {
		log.Panicln("Failed to read file", err)
	}

	var particles []particle
	for _, line := range lines {
		var p particle
		fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &p.x, &p.y, &p.dx, &p.dy)

		particles = append(particles, p)
	}

	space := make(map[point][]particle)

	for _, p := range particles {
		existing := space[p.point]
		space[p.point] = append(existing, p)
	}

	for {
		var missingAdjacency bool
		for k := range space {
			if !hasAdjacentPoint(k, space) {
				missingAdjacency = true
				break
			}
		}

		if !missingAdjacency {
			break
		}

		// Still going, so advance
		newSpace := make(map[point][]particle)
		for _, existing := range space {
			for _, p := range existing {
				k2 := p.point
				k2.x += p.dx
				k2.y += p.dy
				p.point = k2
				dest := newSpace[k2]
				newSpace[k2] = append(dest, p)
			}
		}
		space = newSpace
	}

	printSpace(space)
}

func printSpace(space map[point][]particle) {
	minX, minY, maxX, maxY := math.MaxInt64, math.MaxInt64, math.MinInt64, math.MinInt64
	for k := range space {
		if k.x > maxX {
			maxX = k.x
		}
		if k.x < minX {
			minX = k.x
		}

		if k.y > maxY {
			maxY = k.y
		}
		if k.y < minY {
			minY = k.y
		}
	}

	for y := minY; y <= maxY; y++ {
		var row string
		for x := minX; x <= maxX; x++ {
			p := point{x, y}
			_, ok := space[p]
			if ok {
				row += "#"
			} else {
				row += "."
			}
		}
		log.Println(row)
	}

}

func hasAdjacentPoint(p point, g map[point][]particle) bool {
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			t := point{p.x + x, p.y + y}
			if t == p {
				continue
			}
			if _, ok := g[t]; ok {
				return true
			}
		}
	}

	return false
}

func p02() {
	lines, err := aoc.ReadStrings("10.txt")
	if err != nil {
		log.Panicln("Failed to read file", err)
	}

	var particles []particle
	for _, line := range lines {
		var p particle
		fmt.Sscanf(line, "position=<%d,%d> velocity=<%d,%d>", &p.x, &p.y, &p.dx, &p.dy)

		particles = append(particles, p)
	}

	space := make(map[point][]particle)

	for _, p := range particles {
		existing := space[p.point]
		space[p.point] = append(existing, p)
	}

	var loops int
	for {
		var missingAdjacency bool
		for k := range space {
			if !hasAdjacentPoint(k, space) {
				missingAdjacency = true
				break
			}
		}

		if !missingAdjacency {
			break
		}

		// Still going, so advance
		newSpace := make(map[point][]particle)
		for _, existing := range space {
			for _, p := range existing {
				k2 := p.point
				k2.x += p.dx
				k2.y += p.dy
				p.point = k2
				dest := newSpace[k2]
				newSpace[k2] = append(dest, p)
			}
		}
		loops++
		space = newSpace
	}

	log.Println("P02:", loops)
}

type point struct {
	x, y int
}
type particle struct {
	point
	dx, dy int
}
