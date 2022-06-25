package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mattn/go-tty"
)

func main() {
	flag.Parse()

	path := flag.Args()[0]
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	s := bufio.NewScanner(f)

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}

	for s.Scan() {
		fmt.Printf("%s\r", s.Text())
		isFirst := false
		for _, c := range s.Text() {
			if !isFirst {
				if c == ' ' || c == '\t' {
					fmt.Printf("%c", c)
					continue
				} else {
					isFirst = true
				}
			}
			for {
				r, err := tty.ReadRune()
				if err != nil {
					log.Fatal(err)
				}
				if r == c {
					fmt.Printf("\x1b[32m%c\x1b[0m", c)
					break
				}
			}
		}
		fmt.Println()
	}

	err = tty.Close()
	if err != nil {
		log.Fatal(err)
	}
}