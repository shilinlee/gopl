// go run main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// WaitForServer 尝试连接url对应的服务器
// 在一分钟内使用指数退避策略进行充实
// 所有的尝试失败，返回错误
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // 成功
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries)) // 指数退避策略
	}
	return fmt.Errorf("server %s failed to response after %s", url, timeout)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	//!+main
	// (In function main.)
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
	//!-main
}
/*
$ go run main.go https://golang.org
2020/12/16 00:08:16 server not responding (Head "https://golang.org": dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.); retrying...
2020/12/16 00:08:38 server not responding (Head "https://golang.org": dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.); retrying...
2020/12/16 00:09:01 server not responding (Head "https://golang.org": dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.); retrying...
Site is down: server https://golang.org failed to response after 1m0s
 */