package main

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type heading int

const (
	north heading = iota
	east
	south
	west
)

type state struct {
	x, y int
	h    heading
}

// parseDirection takes a direction of the form 'L7' or 'R20' and updates the
// position/heading state accordingly
func (s *state) applyDirection(direction string) {
	switch direction[0] {
	case 'L':
		if s.h == north {
			s.h = west
		} else {
			s.h--
		}
	case 'R':
		if s.h == west {
			s.h = north
		} else {
			s.h++
		}
	default:
		log.Panicf("unexpected input: %s", string(direction[0]))
	}
	dist, err := strconv.Atoi(direction[1:])
	if err != nil {
		log.Panicf("Error converting to int: " + err.Error())
	}
	switch s.h {
	case north:
		s.y += dist
	case east:
		s.x += dist
	case south:
		s.y -= dist
	case west:
		s.x -= dist
	}
}

func iabs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (s *state) blocksAway() (dist int) {
	return iabs(s.x) + iabs(s.y)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	directions := strings.Split(strings.Trim(string(data), "\n"), ", ")
	s := new(state)
	for _, d := range directions {
		s.applyDirection(d)
	}
	log.Printf("%d blocks away!", s.blocksAway())
}
