package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/pitakill/swiss-knife/slice"
)

// Examples of use
func main() {
	// String
	stringsToTransform := []string{"Hello", "World"}

	if err := slice.ForEach(stringsToTransform, func(s string, i int) {
		stringsToTransform[i] = strings.ToUpper(s)
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println(stringsToTransform)

	if err := slice.ForEach(stringsToTransform, func(s string, i int) error {
		fmt.Println(s, i)
		return nil
	}); err != nil {
		log.Println("this won't work because of the firm of the func passed to the StringForEach: ", err)
	}

	// Int
	intsToTranform := []int{23, 458}
	var add int

	if err := slice.ForEach(intsToTranform, func(i int) {
		add += i
	}); err != nil {
		log.Fatal(err)
	}

	fmt.Println(add)
}
