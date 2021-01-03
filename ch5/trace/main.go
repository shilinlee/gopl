// go run main.go
package main

import (
	"log"
	"time"
)

//!+main
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // 别忘记这对括号
	// ...lots of work...
	time.Sleep(10 * time.Second) // 通过休眠仿真慢操作
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() { log.Printf("exit %s (%s)", msg, time.Since(start)) }
}

//!-main

func main() {
	bigSlowOperation()
}

/*
!+output
$ go run main.go
2015/11/18 09:53:26 enter bigSlowOperation
2015/11/18 09:53:36 exit bigSlowOperation (10.000589217s)
!-output
*/