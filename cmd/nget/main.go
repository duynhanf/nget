package main

import (
	"fmt"

	"github.com/duynhanf/nget"
)

func main() {
	fmt.Println("nget start")

	h := nget.NewHTTPDownloader("https://dl.google.com/go/go1.13.1.src.tar.gz")

	fmt.Println(h)

	h.Do()
}
