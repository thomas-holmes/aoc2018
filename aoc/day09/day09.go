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

func mod(new, max int) int {
	m := new % max
	if m < 0 {
		return m + max
	}

	return m
}

func p01() {
	players, lastMarble := 473, 7090400

	var circle = []int{0}
	scores := make(map[int]int)
	currentMarble, currentPlayer := 0, 0

	// I'm pretty certain this is trashing my performance with allocations
	for i := 1; i <= lastMarble; i++ {
		//log.Printf("M=%d, C=%d, %v", i, currentMarble, circle)
		if i%23 == 0 {
			scores[currentPlayer] += i
			currentMarble = mod((currentMarble - 7), len(circle))
			removedMarble := circle[currentMarble]
			scores[currentPlayer] += removedMarble
			circle = append(circle[:currentMarble], circle[currentMarble+1:]...)

			currentMarble = mod(currentMarble, len(circle))
			currentPlayer = mod(currentPlayer+1, players)
		} else {
			currentMarble = mod(currentMarble+1, len(circle)) // move clockwise one

			front := append([]int{}, circle[:currentMarble+1]...)
			back := append([]int{}, circle[currentMarble+1:]...)
			// log.Printf("B1: C=%d CV=%d F=%v, B=%v Cr=%v", currentMarble, circle[currentMarble], front, back, circle)

			var newCircle []int
			newCircle = append(front, i)
			newCircle = append(newCircle, back...)
			// log.Printf("B2: C=%d CV=%d F=%v, B=%v Cr=%v", currentMarble, circle[currentMarble], front, back, circle)
			circle = newCircle
			// log.Printf("B3: C=%d CV=%d F=%v, B=%v Cr=%v", currentMarble, circle[currentMarble], front, back, circle)

			currentMarble++

			// log.Printf("A: C=%d CV=%d Circle=%v", currentMarble, circle[currentMarble], circle)

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

			scores[currentPlayer] += removedMarble.val

			currentPlayer = mod(currentPlayer+1, players)
		} else {
			currentMarble = currentMarble.next

			newMarble := newNode(i, currentMarble, currentMarble.next)

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
