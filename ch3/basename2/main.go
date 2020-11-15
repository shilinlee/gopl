// go run main.go

package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	slash := strings.LastIndex(s, "/") // 如果没找到"/", slash = -1
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func main() {
	fmt.Println(basename("a"))
	fmt.Println(basename("a/b/c.d.go"))
}
