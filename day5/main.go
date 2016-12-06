package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

type door struct {
	ID    string
	index int
}

func newDoor(doorID string) *door {
	return &door{doorID, 0}
}

func (d *door) nextChar() rune {
	var h hash.Hash
	for {
		h = md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", d.ID, d.index))
		d.index++
		hexHash := hex.EncodeToString(h.Sum(nil))
		if hexHash[0:5] == "00000" {
			return rune(hexHash[5])
		}
		switch {
		case d.index%500000 == 0:
			fmt.Printf("%d", d.index)
		case d.index%100000 == 0:
			fmt.Print(".")
		}
	}
}

func main() {
	d := newDoor(os.Args[1])
	var pw string
	for i := 0; i < 8; i++ {
		theChar := string(d.nextChar())
		fmt.Println("Found: " + theChar)
		pw += theChar
	}
	fmt.Println("pw: " + pw)

}
