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
	inputs := bufio.NewScanner(os.Stdin)
	for inputs.Scan() {
		if inputs.Text() == "bye" {
			break
		}
		counts[inputs.Text()]++
	}

	// 注意忽略input.Err()中可能存在的错误
	for line, n := range counts {

		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
