// go run main.go

package main

import "fmt"

func basename(s string) string {
	// 移除路径部分和.后缀
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}

	}
	// 保留最后一个.之前的所有内容
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func main() {
	fmt.Println(basename("a"))
	fmt.Println(basename("a/b/c.d.go"))
}
