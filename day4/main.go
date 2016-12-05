package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"unicode"
)

type room struct {
	name        string
	sectorID    int
	checksum    string
	rawSectorID string
}

type parseState int

const (
	parseStateName parseState = iota
	parseStateSectorID
	parseStateChecksum
)

func parseRoom(inString string) *room {
	newroom := &room{}
	state := parseStateName
	for _, r := range inString {
		switch state {
		case parseStateName:
			if unicode.IsLower(r) || r == '-' {
				newroom.name += string(r)
			} else if unicode.IsNumber(r) {
				newroom.rawSectorID += string(r)
				state = parseStateSectorID
			} else {
				log.Panicf("Unexpected rune %s while parsing name field", string(r))
			}
		case parseStateSectorID:
			if unicode.IsNumber(r) {
				newroom.rawSectorID += string(r)
			} else if r == '[' {
				state = parseStateChecksum
			} else {
				log.Panicf("Unexpected rune %s while parsing sectorid field", string(r))
			}
		case parseStateChecksum:
			if unicode.IsLower(r) {
				newroom.checksum += string(r)
			} else if r == ']' {
				break
			} else {
				log.Panicf("Unexpected rune %s while parsing checksum field", string(r))
			}
		}
	}
	var err error
	newroom.sectorID, err = strconv.Atoi(newroom.rawSectorID)
	if err != nil {
		log.Panicf("Non-numeric sector id: %s", newroom.rawSectorID)
	}

	return newroom
}

type runeCount struct {
	theRune  rune
	theCount int
}

type runeCounts []runeCount

func (rc runeCounts) Len() int { return len(rc) }

func (rc runeCounts) Less(i, j int) bool {
	if rc[i].theCount == rc[j].theCount {
		return rc[i].theRune < rc[j].theRune
	}
	return rc[i].theCount > rc[j].theCount
}

func (rc runeCounts) Swap(i, j int) {
	rc[i], rc[j] = rc[j], rc[i]
}

func (r *room) calculateChecksum() string {
	runecountmap := make(map[rune]int)
	for _, aRune := range r.name {
		if unicode.IsLower(aRune) {
			runecountmap[aRune]++
		}
	}
	counts := make(runeCounts, len(runecountmap))
	i := 0
	for aRune, aCount := range runecountmap {
		counts[i] = runeCount{aRune, aCount}
		i++
	}
	sort.Sort(counts)

	var csum string
	for i := 0; i < 5; i++ {
		csum += string(counts[i].theRune)
	}
	return csum
}

func (r *room) isValid() bool {

	return r.calculateChecksum() == r.checksum
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	var sectorSum int
	for s.Scan() {
		if room := parseRoom(s.Text()); room.isValid() {
			sectorSum += room.sectorID
		}
	}
	log.Printf("Part1 sector sum: %d\n", sectorSum)
}
