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
		fmt.Println(s.Text())
		for _, c := range s.Text() {
			for {
				r, err := tty.ReadRune()
				if err != nil {
					log.Fatal(err)
				}
				if r == c {
					fmt.Print(string(c))
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
