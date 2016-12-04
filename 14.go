// Duplicate the elements of a list.
// Example:
//
// * (dupli '(a b c c d))
// (A A B B C C C C D D)
package main

import (
	"fmt"
)

func Dupli(arg []string) []string {
    var ret []string
    for _, x := range arg {
        ret = append(ret, x)
        ret = append(ret, x)
    }
    return ret
}

func main() {
  fmt.Println(Dupli([]string{"a", "a", "b", "c", "c", "d"}))
}
