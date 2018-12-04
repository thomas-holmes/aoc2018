package main

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/thomas-holmes/aoc2018/aoc"
)

func main() {
	p01()
	p02()
}

type record struct {
	id    int
	t     time.Time
	state string
}

func p01() {
	data, err := aoc.ReadStrings("4.txt")
	if err != nil {
		log.Panicln("Failed to read data", err)
	}

	var records []record
	for _, line := range data {

		var rec record

		letters := []rune(line)
		timeString := string(letters[1:17])
		rest := string(letters[19:])

		ts, err := time.Parse("2006-01-02 15:04", timeString)
		if err != nil {
			log.Panicln("Failed to parse timestamp", err)
		}

		rec.t = ts

		switch string(rest[:5]) {
		case "wakes":
			rec.state = "w"
		case "falls":
			rec.state = "f"
		case "Guard":
			rec.state = "g"
			var guardId int
			_, err := fmt.Sscanf(string(rest[6:11]), "#%d", &guardId)
			if err != nil {
				log.Panicln("Failed to scan", string(rest[6:11]))
			}
			rec.id = guardId
		}

		records = append(records, rec)
	}

	sort.Slice(records, func(i, j int) bool { return records[i].t.Before(records[j].t) })

	sleeping := make(map[int]int)
	minuteMap := make(map[int][60]int)
	// var state string
	var guard int

	var startSleep time.Time
	for _, r := range records {
		switch r.state {
		case "g":
			guard = r.id
		case "f":
			startSleep = r.t
		case "w":
			asleep := r.t.Sub(startSleep)
			minutes := sleeping[guard]
			sleeping[guard] = minutes + int(asleep.Minutes())

			mm := minuteMap[guard]
			iter := startSleep
			for {
				if !iter.Before(r.t) {
					break
				}
				mm[iter.Minute()]++
				iter = iter.Add(1 * time.Minute)
			}
			minuteMap[guard] = mm
		}
	}

	var maxTime int
	for k, v := range sleeping {
		if v > maxTime {
			guard = k
			maxTime = v
		}
	}

	var mostSlept int
	var sleepiestMinute int
	for i, v := range minuteMap[guard] {
		if v > mostSlept {
			mostSlept = v
			sleepiestMinute = i
		}
	}

	log.Println("Guard", guard, "slept for", maxTime, "minutes", "map", minuteMap[guard], "sleepiest minute", sleepiestMinute, "for", mostSlept)
}

func p02() {
	data, err := aoc.ReadStrings("4.txt")
	if err != nil {
		log.Panicln("Failed to read data", err)
	}

	var records []record
	for _, line := range data {

		var rec record

		letters := []rune(line)
		timeString := string(letters[1:17])
		rest := string(letters[19:])

		ts, err := time.Parse("2006-01-02 15:04", timeString)
		if err != nil {
			log.Panicln("Failed to parse timestamp", err)
		}

		rec.t = ts

		switch string(rest[:5]) {
		case "wakes":
			rec.state = "w"
		case "falls":
			rec.state = "f"
		case "Guard":
			rec.state = "g"
			var guardId int
			_, err := fmt.Sscanf(string(rest[6:11]), "#%d", &guardId)
			if err != nil {
				log.Panicln("Failed to scan", string(rest[6:11]))
			}
			rec.id = guardId
		}

		records = append(records, rec)
	}

	sort.Slice(records, func(i, j int) bool { return records[i].t.Before(records[j].t) })

	sleeping := make(map[int]int)
	minuteMap := make(map[int][60]int)
	// var state string
	var guard int

	var startSleep time.Time
	for _, r := range records {
		switch r.state {
		case "g":
			guard = r.id
		case "f":
			startSleep = r.t
		case "w":
			asleep := r.t.Sub(startSleep)
			minutes := sleeping[guard]
			sleeping[guard] = minutes + int(asleep.Minutes())

			mm := minuteMap[guard]
			iter := startSleep
			for {
				if !iter.Before(r.t) {
					break
				}
				mm[iter.Minute()]++
				iter = iter.Add(1 * time.Minute)
			}
			minuteMap[guard] = mm
		}
	}

	var mostSlept int
	var sleepiestMinute int
	var sleepiestGuard int
	for g, mm := range minuteMap {
		for minute, slept := range mm {
			if slept > mostSlept {
				mostSlept = slept
				sleepiestMinute = minute
				sleepiestGuard = g
			}
		}
	}

	log.Println("Guard", sleepiestGuard, "slept for", mostSlept, "minutes", " on sleepiest minute", sleepiestMinute)

}
