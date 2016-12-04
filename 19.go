// Rotate a list N places to the left.
//
// Examples:
//
// * (rotate '(a b c d e f g h) 3)
// (D E F G H A B C)
//
// * (rotate '(a b c d e f g h) -2)
// (G H A B C D E F)
package main

import (
	"fmt"
)

func Rotate(arg []string, n int) []string {
    var ret []string
    if n < 1 {
        n = len(arg) + n
    }
    ret = append(ret, arg[n:]...)
    ret = append(ret, arg[0:n]...)
	return ret
}

func main() {
	fmt.Println(Rotate([]string{"a", "b", "c", "d", "e", "f", "g"}, 3))
	fmt.Println(Rotate([]string{"a", "b", "c", "d", "e", "f", "g"}, -2))
}
