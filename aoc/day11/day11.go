package main

import (
	"log"
	"time"
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
	Input := 5177

	grid := make([]int, 300*300)

	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			grid[y*300+x] = power(x+1, y+1, Input)
		}
	}

	var bestSquare int
	var bestX, bestY int

	for x := 0; x <= 298; x++ {
		for y := 0; y <= 298; y++ {
			var total int
			// log.Printf("Checking (%d,%d)", x, y)
			for dx := 0; dx < 3; dx++ {
				for dy := 0; dy < 3; dy++ {
					total += power(x+dx+1, y+dy+1, Input)
				}
			}

			if total > bestSquare {
				bestSquare = total
				bestX = x + 1
				bestY = y + 1
				// log.Printf("Found new best of %d at (%d,%d)", bestSquare, bestX, bestY)
			}
		}
	}

	log.Printf("Best square is at (%d,%d)", bestX, bestY)
}

func power(x, y, serial int) int {
	// log.Printf("Calculating %d, %d serial %d", x, y, serial)
	s1 := x + 10
	s2 := s1 * y
	s3 := s2 + serial
	s4 := s3 * s1
	s5 := (s4 / 100) % 10
	s6 := s5 - 5

	// log.Printf("s1[%d] s2[%d] s3[%d] s4[%d] s5[%d] s6[%d]", s1, s2, s3, s4, s5, s6)
	return s6
}

func p02() {
	// MaxX, MaxY := 300, 300
	// Input := 18
	MaxX, MaxY := 300, 300
	Input := 5177
	grid := make([]int, MaxX*MaxY)

	for x := 0; x < MaxX; x++ {
		for y := 0; y < MaxY; y++ {
			grid[y*MaxX+x] = power(x+1, y+1, Input)
		}
	}

	var bestSquare int
	var bestX, bestY, bestSize int

	for x := 0; x < MaxX; x++ {
		for y := 0; y < MaxY; y++ {
			for size := 1; size+x <= MaxX && size+y <= MaxY; size++ {
				if x%10 == 0 && y == 0 && size == 1 {
					log.Printf("Checking (%d,%d) %d", x, y, size)
				}
				var total int
				for dx := 0; dx < size; dx++ {
					for dy := 0; dy < size; dy++ {
						total += power(x+dx+1, y+dy+1, Input)
					}
				}

				if total > bestSquare {
					bestSquare = total
					bestX = x + 1
					bestY = y + 1
					bestSize = size
					log.Printf("Found new best of %d at (%d,%d) %d power %d", bestSquare, bestX, bestY, bestSize, bestSquare)
				}
			}
		}
	}

	log.Printf("Best square is at (%d,%d)%d", bestX, bestY, bestSize)
}
