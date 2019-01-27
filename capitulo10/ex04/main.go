package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kdama/gopl/ch10/ex04/golist"
)

func main() {
	dependees, err := golist.Dependees(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}
	for _, dependee := range dependees {
		fmt.Println(dependee.ImportPath)
	}
}
