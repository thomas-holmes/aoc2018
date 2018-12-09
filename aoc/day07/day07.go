package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/thomas-holmes/aoc2018/aoc"
)

func main() {
	// p01()
	p02()
}

func p01() {
	lines, err := aoc.ReadStrings("7.txt")
	if err != nil {
		log.Panicln("Failed to read file", err)
	}

	deps := make([]dep, 0, len(lines))
	for _, s := range lines {
		var d dep
		fmt.Sscanf(s, "Step %s must be finished before step %s can begin.", &d.b, &d.a)

		deps = append(deps, d)
	}

	var start *step
	graph := make(map[string]*step)
	for _, d := range deps {
		// after
		a, ok := graph[d.a]
		if !ok {
			a = newStep(d.a)
		}

		// before
		b, ok := graph[d.b]
		if !ok {
			b = newStep(d.b)
		}

		a.depends[d.b] = b
		b.frees[d.a] = a

		graph[d.a] = a
		graph[d.b] = b

		if len(b.depends) == 0 {
			if start == nil {
				start = b
			} else {
				if b.name < start.name {
					start = b
				}
			}
		}
	}

	var order string
	for {
		if len(graph) == 0 {
			break
		}

		var lowest *step

		for _, v := range graph {
			if len(v.depends) != 0 {
				continue
			}

			if lowest == nil || v.name < lowest.name {
				lowest = v
			}
		}

		order += lowest.name
		delete(graph, lowest.name)

		for _, s := range lowest.frees {
			delete(s.depends, lowest.name)
		}
	}

	log.Println("P01:", order)
}

func p02() {
	lines, err := aoc.ReadStrings("7.txt")
	if err != nil {
		log.Panicln("Failed to read file", err)
	}

	deps := make([]dep, 0, len(lines))
	for _, s := range lines {
		var d dep
		fmt.Sscanf(s, "Step %s must be finished before step %s can begin.", &d.b, &d.a)

		deps = append(deps, d)
	}

	var start *step
	graph := make(map[string]*step)
	for _, d := range deps {
		// after
		a, ok := graph[d.a]
		if !ok {
			a = newStep(d.a)
		}

		// before
		b, ok := graph[d.b]
		if !ok {
			b = newStep(d.b)
		}

		a.depends[d.b] = b
		b.frees[d.a] = a

		graph[d.a] = a
		graph[d.b] = b

		if len(b.depends) == 0 {
			if start == nil {
				start = b
			} else {
				if b.name < start.name {
					start = b
				}
			}
		}
	}

	var order string
	const MaxWorkers = 5
	var currentWorkers int
	var working []*step
	var curr int

	var stop int
	for {
		var lowest *step

		log.Println("len(graph)", len(graph), "len(working)", len(working))
		for _, v := range graph {
			if len(v.depends) != 0 {
				continue
			}

			if lowest == nil || v.name < lowest.name {
				lowest = v
			}
		}

		if len(graph) == 0 && len(working) == 0 {
			break
		}

		if currentWorkers == MaxWorkers || lowest == nil {
			log.Printf("currentWorkers=%d MaxWorkers=%d, lowest=%v", currentWorkers, MaxWorkers, lowest)
			// do some work
			sort.Slice(working, func(i, j int) bool { return working[i].done < working[j].done })

			targetTime := working[0].done
			curr = targetTime
			var completed []*step

			log.Printf("Advancing to %d. First up is %s", targetTime, working[0].name)

			var removalIndex int
			for removalIndex = 0; removalIndex < len(working); removalIndex++ {
				s := working[removalIndex]
				if s.done > targetTime {
					// log.Println("working: ", len(working), "before")
					break
				}

				stop++
				if stop > 10 {
					// return
				}
				order += s.name
				log.Println("***Completing [", s.name, "] at T:", curr)
				completed = append(completed, s)
			}

			working = working[removalIndex:]
			currentWorkers = len(working)
			curr = targetTime + 1

			for _, c := range completed {
				for _, d := range c.frees {
					delete(d.depends, c.name)
				}
			}

			continue
		}

		log.Printf("Adding %s to the work queue, T:%d", lowest.name, curr)
		doneTime(lowest, curr)
		working = append(working, lowest)
		delete(graph, lowest.name)
		currentWorkers++
	}

	log.Println("P02:", order, "in T:", curr)
}

func doneTime(s *step, currentTime int) {
	completedTime := currentTime + int(([]rune(s.name)[0]-rune('A'))+60)
	s.done = completedTime
}

type dep struct {
	b, a string
}

type step struct {
	name    string
	depends map[string]*step
	frees   map[string]*step
	done    int
}

func newStep(name string) *step {
	return &step{
		name:    name,
		depends: make(map[string]*step),
		frees:   make(map[string]*step),
	}
}