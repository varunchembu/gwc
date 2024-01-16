package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var useC bool

	flag.BoolVar(&useC, "c", false, "Enable option C")

	flag.Parse()
	nonFlagArgs := flag.Args()
	var filename string

	// fmt.Println("Options:", options)
	if len(nonFlagArgs) > 0 {
		filename = nonFlagArgs[0]
	} else {
		fmt.Println("No filename provided.")
	}

	file, _ := os.Open(filename)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	var byteCount int
	for scanner.Scan() {
		byteCount++
	}

	fmt.Println(byteCount, filename)
}
