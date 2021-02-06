//结合 gopl/ch7/tempconv 使用
//go run tempflag.go -temp 20C
//go run tempflag.go -temp 20F
package main

import (
	"flag"
	"fmt"

	"gopl/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
