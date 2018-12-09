package main

import (
	"fmt"
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

func mod(new, max int) int {
	m := new % max
	if m < 0 {
		return m + max
	}

	return m
}

func p01() {
	players, lastMarble := 473, 70904

	circle := newNode(0, nil, nil)
	circle.prev, circle.next = circle, circle

	scores := make(map[int]int)
	currentMarble := circle
	currentPlayer := 0

	// I'm pretty certain this is trashing my performance with allocations
	for i := 1; i <= lastMarble; i++ {
		//log.Printf("M=%d, C=%d, %v", i, currentMarble, circle)
		if i%23 == 0 {
			scores[currentPlayer] += i
			for j := 0; j < 7; j++ {
				currentMarble = currentMarble.prev
			}

			removedMarble := currentMarble
			currentMarble = currentMarble.next

			removedMarble.next.prev = removedMarble.prev
			removedMarble.prev.next = removedMarble.next

			removedMarble.prev, removedMarble.next = nil, nil

			scores[currentPlayer] += removedMarble.val

			currentPlayer = mod(currentPlayer+1, players)
		} else {
			currentMarble = currentMarble.next

			newMarble := newNode(i, currentMarble, currentMarble.next)

			currentMarble.next.prev = newMarble
			currentMarble.next = newMarble

			currentMarble = newMarble

			currentPlayer = mod(currentPlayer+1, players)
		}
	}

	var max int
	for _, v := range scores {
		if v > max {
			max = v
		}
	}

	log.Println("P01:", max)
}

func p02() {
	players, lastMarble := 473, 7090400

	circle := newNode(0, nil, nil)
	circle.prev, circle.next = circle, circle

	scores := make(map[int]int)
	currentMarble := circle
	currentPlayer := 0

	// I'm pretty certain this is trashing my performance with allocations
	for i := 1; i <= lastMarble; i++ {
		//log.Printf("M=%d, C=%d, %v", i, currentMarble, circle)
		if i%23 == 0 {
			scores[currentPlayer] += i
			for j := 0; j < 7; j++ {
				currentMarble = currentMarble.prev
			}

			removedMarble := currentMarble
			currentMarble = currentMarble.next

			removedMarble.next.prev = removedMarble.prev
			removedMarble.prev.next = removedMarble.next

			removedMarble.prev, removedMarble.next = nil, nil

			scores[currentPlayer] += removedMarble.val

			currentPlayer = mod(currentPlayer+1, players)
		} else {
			currentMarble = currentMarble.next

			newMarble := newNode(i, currentMarble, currentMarble.next)

			currentMarble.next.prev = newMarble
			currentMarble.next = newMarble

			currentMarble = newMarble

			currentPlayer = mod(currentPlayer+1, players)
		}
	}

	var max int
	for _, v := range scores {
		if v > max {
			max = v
		}
	}

	log.Println("P02:", max)
}

func newNode(val int, prev, next *node) *node {
	return &node{
		prev: prev,
		next: next,
		val:  val,
	}
}

type node struct {
	prev *node
	next *node
	val  int
}

func (n node) String() string {
	return fmt.Sprintf("%d <- %d -> %d", n.prev.val, n.val, n.next.val)
}
