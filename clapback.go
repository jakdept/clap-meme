package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input := os.Args[1:]
	var clean []string
	for _, each := range input {
		if each != "👏" {
			clean = append(clean, each)
		}
	}
	fmt.Println(strings.Join(strings.Fields(strings.Join(clean, " ")), " 👏 "))
}
