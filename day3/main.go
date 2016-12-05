package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type triangle struct {
	a, b, c int
}

func atoiSlice(intStrings []string) (out []int) {
	for _, intString := range intStrings {
		theInt, err := strconv.Atoi(intString)
		if err != nil {
			panic(err)
		}
		out = append(out, theInt)
	}
	return
}

func newTriangle(dimensions []int) *triangle {
	if len(dimensions) != 3 {
		log.Panicf("Illegal number of dimensions for triangle: %d", len(dimensions))
	}
	return &triangle{dimensions[0], dimensions[1], dimensions[2]}
}

func (t *triangle) isValid() bool {
	return (t.a+t.b > t.c) && (t.b+t.c > t.a) && (t.a+t.c > t.b)
}

func parsePart1Triangles(reader io.Reader) (numValid int) {
	s := bufio.NewScanner(reader)
	for s.Scan() {
		if newTriangle(atoiSlice(strings.Fields(s.Text()))).isValid() {
			numValid++
		}
	}
	return
}

func parsePart2Triangles(reader io.Reader) (numValid int) {
	s := bufio.NewScanner(reader)
	var threeLines [][]string
	for s.Scan() {
		threeLines = append(threeLines, strings.Fields(s.Text()))
		if len(threeLines) == 3 {
			var extractedDimensions []string
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					extractedDimensions = append(extractedDimensions, threeLines[j][i])
				}
				if newTriangle(atoiSlice(extractedDimensions)).isValid() {
					numValid++
				}
				extractedDimensions = []string{}

			}
			threeLines = [][]string{}
		}
	}
	return
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fmt.Printf("Part 1 valid triangles: %d\n", parsePart1Triangles(file))
	file.Seek(0, 0)
	fmt.Printf("Part 2 valid triangles: %d\n", parsePart2Triangles(file))
}
