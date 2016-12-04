//
// Extract a slice from a list.
//
// Given two indices, i and k, the slice is the list containing the elements between the i'th and k'th element of the original list (both limits included). Start counting the elements with 1.
//
// Example:
//
// * (slice '(a b c d e f g h i k) 3 7)
// (C D E F G)
package main

import (
  "fmt"
)

func Slice(arg []string, n int, m int) []string {
  if n < 1 || n > len(arg) || m < n || m > len(arg) {
    return arg
  }
  var ret []string
  for i, x := range arg {
    if (i + 1) >= n && (i + 1) <= m {
      ret = append(ret, x)
    }
  }
  return ret
}

func main() {
  fmt.Println(Slice([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "k"}, 3, 7))
}


