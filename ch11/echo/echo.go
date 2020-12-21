//可以先对36行取消注释，更明显看出换行符效果
//go run echo.go -n=true -s=, 1 2 3
// go run echo.go -n=false -s=, 1 2 3
// Echo 输出它的命令行参数
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newline") // 是否省略尾随换行符
	s = flag.String("s", " ", "separator")
)

var out io.Writer = os.Stdout

func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		// 换行
		fmt.Fprintln(out)
	}
	// fmt.Fprint(out, "这是尾巴")
	return nil
}
