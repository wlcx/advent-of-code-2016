package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type keypad struct {
	layout     map[int]map[int]rune
	posx, posy int
}

func newPart1Keypad() *keypad {
	return &keypad{
		map[int]map[int]rune{
			0: map[int]rune{
				0: '1',
				1: '2',
				2: '3',
			},
			1: map[int]rune{
				0: '4',
				1: '5',
				2: '6',
			},
			2: map[int]rune{
				0: '7',
				1: '8',
				2: '9',
			},
		},
		1, 1,
	}
}

func newPart2Keypad() *keypad {
	return &keypad{
		map[int]map[int]rune{
			0: map[int]rune{
				2: '1',
			},
			1: map[int]rune{
				1: '2',
				2: '3',
				3: '4',
			},
			2: map[int]rune{
				0: '5',
				1: '6',
				2: '7',
				3: '8',
				4: '9',
			},
			3: map[int]rune{
				1: 'A',
				2: 'B',
				3: 'C',
			},
			4: map[int]rune{
				2: 'D',
			},
		},
		2, 0,
	}
}

func (k *keypad) applyInput(reader io.Reader) (code string) {
	s := bufio.NewScanner(reader)
	for s.Scan() {
		code += k.applyLine(s.Text())
	}
	return
}

func (k *keypad) applyLine(instructions string) string {
	for _, instr := range instructions {
		switch instr {
		case 'U':
			if _, ok := k.layout[k.posy-1]; ok {
				if _, ok := k.layout[k.posy-1][k.posx]; ok {
					k.posy--
				}
			}
		case 'R':
			if _, ok := k.layout[k.posy][k.posx+1]; ok {
				k.posx++
			}
		case 'D':
			if _, ok := k.layout[k.posy+1]; ok {
				if _, ok := k.layout[k.posy+1][k.posx]; ok {
					k.posy++
				}
			}
		case 'L':
			if _, ok := k.layout[k.posy][k.posx-1]; ok {
				k.posx--
			}
		default:
			log.Panicf("Unexpected direction: %s", string(instr))
		}
	}
	return string(k.layout[k.posy][k.posx])
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	k1 := newPart1Keypad()
	k2 := newPart2Keypad()
	fmt.Println("Part 1 code: " + k1.applyInput(file))
	file.Seek(0, 0)
	fmt.Println("Part 2 code: " + k2.applyInput(file))
}
