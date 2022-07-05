package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mattn/go-tty"
)

func getFromPath(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, nil
}

func getFromUrl(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	s := bufio.NewScanner(res.Body)
	lines := []string{}
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines, nil
}

func main() {
	flag.Parse()

	target := flag.Args()[0]

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}

	var lines []string
	if lines, err = getFromPath(target); err != nil {
		if lines, err = getFromUrl(target); err != nil {
			fmt.Printf("Error: Invalid target: %s\n", target)
			os.Exit(0)
		}
	}

	for _, line := range lines {
		fmt.Printf("%s\r", line)
		isFirst := false
		for _, c := range line {
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
