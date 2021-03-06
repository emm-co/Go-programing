package main

import (
	"fmt"
	"os"

	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("ch05/ex09: %v", err)
	}
	fmt.Fprintf(os.Stdout, expand(string(b), double))
}

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$[^\s]+`)
	return re.ReplaceAllStringFunc(s, func(x string) string {
		return f(x[1:])
	})
}

func double(s string) string {
	return s + s
}
