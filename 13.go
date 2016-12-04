// Run-length encoding of a list (direct solution).
//
// Implement the so-called run-length encoding data compression method directly.
// I.e. don't explicitly create the sublists containing the duplicates, as in
// problem 9, but only count them. As in problem P11, simplify the result list
// by replacing the singleton lists (1 X) by X.
//
// Example:
//
// * (encode-direct '(a a a a b c c a a d e e e e))
// ((4 A) B (2 C) (2 A) D (4 E))
//
// https://wiki.haskell.org/99_questions/11_to_20
//
package main

import (
	"fmt"
)

func EncodeDirect(arg []string) []interface{} {
	ret := make([]interface{}, 0)
	i := 0
	for i <= len(arg)-1 {
		samesies := 0
		for (i+samesies+1) <= len(arg)-1 && arg[i] == arg[i+samesies+1] {
			samesies += 1
		}
		
		if samesies > 0 {
			ret = append(ret, []interface{}{samesies + 1, arg[i]})
			i += samesies + 1
		} else {
			ret = append(ret, arg[i])
			i += 1
		}
	}
	return ret
}

func main() {
	fmt.Println(EncodeDirect([]string{"A", "A", "B", "C", "A", "A", "A", "C", "C", "D", "E"}))
}
