package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type triangle struct {
	a, b, c int
}

func parseTriangle(input string) *triangle {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanWords)
	var dims []string
	for s.Scan() {
		dims = append(dims, s.Text())
	}
	if len(dims) != 3 {
		panic("Number of input dimensions != 3")
	}
	var intDims []int
	for _, rawDim := range dims {
		dim, err := strconv.Atoi(rawDim)
		if err != nil {
			panic(err)
		}
		intDims = append(intDims, dim)
	}
	return &triangle{intDims[0], intDims[1], intDims[2]}
}

func (t *triangle) isValid() bool {
	return (t.a+t.b > t.c) && (t.b+t.c > t.a) && (t.a+t.c > t.b)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	valid := 0
	for s.Scan() {
		if parseTriangle(s.Text()).isValid() {
			valid++
		}
	}
	fmt.Printf("Valid triangles: %d\n", valid)
}
