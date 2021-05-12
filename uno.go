package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/paulebose/uno/strextra"
)

func main() {
	fmt.Println("Hello People!")
	fmt.Println(strextra.ReverseRunes("Hello People!"))
	fmt.Println(cmp.Diff("Hello Boy", "Hello People!"))
}
