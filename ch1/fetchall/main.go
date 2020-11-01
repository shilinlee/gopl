// go run main.go http://gopl.io https://godoc.org http://shilinlee.com

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		// 启动一个goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// 从通道ch接收
		fmt.Println(<-ch)
	}
	// elapsed 过去
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 发送到通道ch
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 不要泄露资源
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
