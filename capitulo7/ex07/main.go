package main

import (
	"flag"
	"fmt"

	"github.com//emm-co/go/ch07/ex07/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}