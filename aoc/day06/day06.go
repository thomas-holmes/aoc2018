package main

import (
	"fmt"
	"log"

	"github.com/thomas-holmes/aoc2018/aoc"
)

func main() {
	p01()
	p02()
}

type point struct {
	x, y int
}

type mark struct {
	point
	distance  int
	collision bool
}

var unset mark
var collision = mark{
	distance: -9999,
}

func dist(p1, p2 point) int {
	xDist := p1.x - p2.x
	yDist := p1.y - p2.y
	if xDist < 0 {
		xDist *= -1
	}
	if yDist < 0 {
		yDist *= -1
	}

	return xDist + yDist
}

func p01() {
	lines, err := aoc.ReadStrings("6.txt")
	if err != nil {
		log.Panicln("Failed to read points", err)
	}

	var points []point

	for _, l := range lines {
		var p point
		fmt.Sscanf(l, "%d, %d", &p.x, &p.y)
		points = append(points, p)
	}

	minX, minY, maxX, maxY := 0xFFFFFFFF, 0xFFFFFFFF, 0, 0

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// I think there are cases that break this. Probably need to make sure i check a
	// box that would contain the triangle point off of each side.

	// Adjust?
	// minX, minY, maxX, maxY = minX, minY, maxX, maxY
	log.Printf("Min (%d,%d) Max (%d,%d)", minX, minY, maxX, maxY)

	grid := make(map[point]mark)
	for _, p := range points {
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				t := point{x, y}

				compareAndSet(grid, p, t)
			}
		}
	}

	excluded := make(map[point]struct{})
	for x := minX; x <= maxX; x++ {
		topY, botY := minY, maxY

		t1, t2 := point{x, topY}, point{x, botY}

		m1 := grid[t1]
		if !m1.collision {
			excluded[m1.point] = struct{}{}
		}

		m2 := grid[t2]
		if !m2.collision {
			excluded[m2.point] = struct{}{}
		}
	}

	for y := minY; y <= maxY; y++ {
		leftX, rightX := minX, maxX

		t1, t2 := point{leftX, y}, point{rightX, y}

		m1 := grid[t1]
		if !m1.collision {
			excluded[m1.point] = struct{}{}
		}

		m2 := grid[t2]
		if !m2.collision {
			excluded[m2.point] = struct{}{}
		}
	}

	counts := make(map[point]int)

	var biggestArea int
	var biggestPoint point

	for _, v := range grid {
		if v.collision {
			continue
		}
		_, ok := excluded[v.point]
		if !ok {
			n := counts[v.point] + 1
			if n > biggestArea {
				biggestArea = n
				biggestPoint = v.point
			}
			counts[v.point] = n
		}
	}

	log.Println(len(grid))

	log.Println("excluded", excluded)

	log.Println(biggestArea, biggestPoint)
}

func compareAndSet(grid map[point]mark, p, t point) {
	distance := dist(p, t)

	a, ok := grid[t]
	if !ok {
		grid[t] = mark{
			point:    p,
			distance: distance,
		}
		return
	}

	if distance < a.distance {
		grid[t] = mark{
			point:    p,
			distance: distance,
		}
	} else if distance == a.distance {
		a.collision = true
		grid[t] = a
	} else {
		// greater than, do nothing
	}
}

func p02() {
	lines, err := aoc.ReadStrings("6.txt")
	if err != nil {
		log.Panicln("Failed to read points", err)
	}

	var points []point

	for _, l := range lines {
		var p point
		fmt.Sscanf(l, "%d, %d", &p.x, &p.y)
		points = append(points, p)
	}

	minX, minY, maxX, maxY := 0xFFFFFFFF, 0xFFFFFFFF, 0, 0

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	// I think there are cases that break this. Probably need to make sure i check a
	// box that would contain the triangle point off of each side.

	// Adjust?
	// minX, minY, maxX, maxY = minX, minY, maxX, maxY
	log.Printf("Min (%d,%d) Max (%d,%d)", minX, minY, maxX, maxY)

	grid := make(map[point]int)
	for _, p := range points {
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				t := point{x, y}

				grid[t] += dist(p, t)
			}
		}
	}

	var safeRegion int

	for _, v := range grid {
		if v < 10000 {
			safeRegion++
		}
	}

	log.Println("Safe region:", safeRegion)
}
