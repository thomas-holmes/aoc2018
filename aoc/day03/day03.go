package main

import (
	"fmt"
	"image"
	"log"

	"github.com/thomas-holmes/aoc2018/aoc"
)

func main() {
	p01()
	p02()
}

type claim struct {
	id int
	image.Rectangle
}

func p01() {

	data, err := aoc.ReadStrings("3.txt")
	if err != nil {
		log.Panicln("Couldn't load data", err)
	}

	var claims []claim

	for _, s := range data {

		var id, x, y, w, h int
		_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)

		c := claim{
			id: id,
			Rectangle: image.Rect(
				x, y, x+w, y+h,
			),
		}

		if err != nil {
			log.Panicln("Couldn't parse string", s)
		}

		claims = append(claims, c)
	}

	var intersections []image.Rectangle
	for i, c1 := range claims {
		for j := i + 1; j < len(claims); j++ {
			c2 := claims[j]

			if !c1.Overlaps(c2.Rectangle) {
				continue
			}

			intersection := c1.Intersect(c2.Rectangle)

			intersections = append(intersections, intersection)
		}
	}

	var space image.Rectangle
	for _, r := range intersections {
		space = space.Union(r)
	}

	var found int
	for x := space.Min.X; x < space.Max.X; x++ {
		for y := space.Min.Y; y < space.Max.Y; y++ {
			point := image.Pt(x, y)
			for _, r := range intersections {
				if point.In(r) {
					found++
					break
				}
			}
		}
	}

	log.Println("p01", found)
}

func p02() {
	data, err := aoc.ReadStrings("3.txt")
	if err != nil {
		log.Panicln("Couldn't load data", err)
	}

	var claims []claim

	for _, s := range data {

		var id, x, y, w, h int
		_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)

		c := claim{
			id: id,
			Rectangle: image.Rect(
				x, y, x+w, y+h,
			),
		}

		if err != nil {
			log.Panicln("Couldn't parse string", s)
		}

		claims = append(claims, c)
	}

	for _, c1 := range claims {
		var overlapped bool
		for _, c2 := range claims {

			// I spent a 30 damn minutes wondering why I couldn't
			// find a non-verlapping rectangle to finally realize
			// I was always comparing to myself. FML
			if c1.id == c2.id {
				continue
			}

			if c1.Overlaps(c2.Rectangle) {
				overlapped = true
				break
			}
		}

		if !overlapped {
			log.Println("I don't overlap!", c1.id)
			return
		}
	}
}
