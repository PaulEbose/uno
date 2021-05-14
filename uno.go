package main

import (
	"errors"

	"fmt"

	"github.com/google/go-cmp/cmp"

	"github.com/paulebose/uno/strextra"
)

func main() {
	fmt.Println("Hello People!")
	fmt.Println(strextra.ReverseRunes("Hello People!"))
	fmt.Println(cmp.Diff("Hello Boy", "Hello People!"))

	// Error codes returned by failures to parse an expression.
	var (
		ErrInternal      = errors.New("regexp: internal error")
		ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
		ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
	)

	fmt.Println(ErrInternal)
	fmt.Println(ErrUnmatchedLpar)
	fmt.Println(ErrUnmatchedRpar)

}
