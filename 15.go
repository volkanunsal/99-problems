// Duplicate the elements of a list a given number of times.
// --------------------------------------------------
//
// Example:
//
// * (repli '(a b c) 3)
// (A A A B B B C C C)
//
package main

import (
	"fmt"
)

func Repli(arg []string, n int) []string {
    var ret []string
    temp := n
    if n < 1 {
        return arg
    }
    for _, x := range arg {
        for n > 0 {
            ret = append(ret, x)
            n -= 1
        }    
        n = temp
    }
	return ret
}

func main() {
	fmt.Println(Repli([]string{"a", "b", "c", "d"}, 9))
}
