// Remove the K'th element from a list.
//
// Example in Lisp:
//
// * (remove-at '(a b c d) 2)
// (A C D)
//
package main

import (
  "fmt"
)

func RemoveAt(arg []string, n int) []string {
  if n < 0 || n > len(arg) {
    return arg
  }
  var ret []string
  ret = append(ret, arg[0:(n-1)]...)
  ret = append(ret, arg[(n):]...)
  return ret
}

func main() {
  fmt.Println(RemoveAt([]string{"a", "b", "c", "d"}, 2))
}
