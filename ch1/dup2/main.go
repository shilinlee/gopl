// go run main.go
// 1
// 1
// bye

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countsLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countsLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

func countsLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if input.Text() == "bye" {
			break
		}
		counts[input.Text()]++
	}
	// 注意忽略input.Err()中可能存在的错误
}
