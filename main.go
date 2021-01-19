package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"fileSuffixArray/suffix_array"
)

func main() {
	fileName := os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	array, err := suffix_array.NewSuffixArray(file)
	if err != nil {
		log.Fatal(err)
	}

	defer array.Terminate()

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		txt := sc.Text()
		fmt.Println(array.FindSubstring(txt))
	}
}
