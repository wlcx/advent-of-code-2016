package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
	"time"
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
	}
}

var bar = []rune{'|', '/', '-', '\\'}

type result struct {
	char rune
	pos  int
}

func main() {
	fmt.Println("Beginning super-sophisticated decryption...")
	pwlength := 8
	d := newDoor(os.Args[1])
	pw := []rune(strings.Repeat("*", pwlength))
	t := time.NewTicker(100 * time.Millisecond)
	barpos := 0
	resCh := make(chan result, 1)
	go func() {
		for i := 0; i < pwlength; i++ {
			resCh <- result{d.nextChar(), i}
		}
		close(resCh)
	}()
MainLoop:
	for {
		select {
		case <-t.C:
			barpos++
			barpos %= 4
		case result, open := <-resCh:
			if !open {
				fmt.Println("\nComplete!")
				break MainLoop
			}
			pw[result.pos] = result.char
		}
		fmt.Print(string(bar[barpos]) + " Decrypting: " + string(pw) + "\r")
	}

}
