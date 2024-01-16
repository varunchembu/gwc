package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

type Stats struct {
	Lines int
	Words int
	Bytes int
	Runes int
}

func main() {
	var useC, useL, useW, useM bool

	flag.BoolVar(&useC, "c", false, "Print byte count")
	flag.BoolVar(&useL, "l", false, "Print line count")
	flag.BoolVar(&useW, "w", false, "Print word count")
	flag.BoolVar(&useM, "m", false, "Print rune count")

	flag.Parse()

	var filename string
	if nonFlagArgs := flag.Args(); len(nonFlagArgs) > 0 {
		filename = nonFlagArgs[0]
	}
	lineCount, wordCount, runeCount, byteCount := countStats(filename)

	stats := Stats{
		Lines: lineCount,
		Words: wordCount,
		Bytes: byteCount,
		Runes: runeCount,
	}

	printResults(stats, filename, useL, useW, useC, useM)

}

func countStats(filename string) (lineCount, wordCount, runeCount, byteCount int) {

	var reader io.Reader

	if filename == "" {
		reader = os.Stdin
	} else {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader = file
	}

	bufReader := bufio.NewReader(reader)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "error reading file: %v\n", err)
			os.Exit(1)
		}
		if line != "" {
			lineCount++
			wordCount += len(strings.Fields(line))
			runeCount += utf8.RuneCountInString(line)
			byteCount += len(line)
		}
		if err == io.EOF {
			break
		}
	}

	return
}

func printResults(stats Stats, filename string, useL, useW, useC, useM bool) {

	lineCount, wordCount, runeCount, byteCount := stats.Lines, stats.Words, stats.Runes, stats.Bytes
	if !(useC || useL || useW || useM) {
		fmt.Printf("%d %d %d %s\n", lineCount, wordCount, byteCount, filename)
		return
	}
	if useL {
		fmt.Printf("%d ", lineCount)
	}
	if useW {
		fmt.Printf("%d ", wordCount)
	}
	if useC {
		fmt.Printf("%d ", byteCount)
	}
	if useM {
		fmt.Printf("%d ", runeCount)
	}
	fmt.Println(filename)

}
