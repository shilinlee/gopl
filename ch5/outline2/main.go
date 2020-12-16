// go run main.go https://github.com
package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// forEachNode 调用 pre(x) and post(x) 遍历以n为根的树中的每个节点x
// 两个函数是可选的
// pre在子叶子被访问前（前序）调用
// post在访问后（后序）调用

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	forEachNode(doc, startElement, endElement)
	return nil
}

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}
