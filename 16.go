// Drop every Nth element from a list.
//
// Example:
//
// * (drop '(a b c d e f g h i k) 3)
// (A B D E G H K)
//
package main

import (
	"fmt"
)

func Drop(arg []string, n int) []string {
  if n < 1 {
    return arg
  }
  var ret []string
  for i, x := range arg {
    if ((i + 1) % n) != 0 {
      ret = append(ret, x)
    }
  } 
  return ret
}

func main() {
  fmt.Println(Drop([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "k"}, 2))
}

