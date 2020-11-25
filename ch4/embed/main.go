// go run main.go
package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	// %#v 表示用和GO语言类似的语法打印值
	fmt.Printf("%#v\n", w)

	w2 := Wheel{
		Circle: Circle{
			Point: Point{
				X: 8,
				Y: 8,
			},
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", w2)

	w2.X = 42
	fmt.Printf("%#v\n", w2)
}
