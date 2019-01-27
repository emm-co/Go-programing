package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ch05/ex03: %v\n", err)
		os.Exit(1)
	}
	visit(os.Stdout, doc)
}

func visit(w io.Writer, n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		visit(w, n.NextSibling)
		return
	} else if n.Type == html.TextNode {
		w.Write([]byte(n.Data))
	}
	visit(w, n.FirstChild)
	visit(w, n.NextSibling)
}
