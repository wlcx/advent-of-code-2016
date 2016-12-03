package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type keypad struct {
	posx, posy int
}

func (k *keypad) applyLine(instructions string) string {
	for _, instr := range instructions {
		switch instr {
		case 'U':
			if k.posy > 0 {
				k.posy--
			}
		case 'R':
			if k.posx < 2 {
				k.posx++
			}
		case 'D':
			if k.posy < 2 {
				k.posy++
			}
		case 'L':
			if k.posx > 0 {
				k.posx--
			}
		default:
			log.Panicf("Unexpected direction: %s", string(instr))
		}
	}
	return strconv.Itoa((k.posx + 1) + (k.posy * 3))
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	k := keypad{1, 1}
	var code string
	for s.Scan() {
		code += k.applyLine(s.Text())
	}
	fmt.Println("The code: " + code)
}
