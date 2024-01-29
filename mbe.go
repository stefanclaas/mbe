package main

import (
	"bufio"
	"fmt"
	"mime"
	"os"
)

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("Usage: $ echo -n 'UTF-8 chars' | mbe > outfile")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		encoded := mime.BEncoding.Encode("UTF-8", scanner.Text())
		fmt.Println(encoded)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}

