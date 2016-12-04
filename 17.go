// Split a list into two parts; the length of the first part is given.
//
// Example:
//
// * (split '(a b c d e f g h i k) 3)
// ( (A B C) (D E F G H I K))
package main

import (
	"fmt"
)

func Split(arg []string, n int) [][]string  {
  if n < 1 || n > len(arg) {
    return [][]string{arg, []string{}}
  }
  var uno []string
  var duo []string
  
  for i, x := range arg {
    if i < n {
      uno = append(uno, x)
    } else {
      duo = append(duo, x)
    }
  }
  return [][]string{uno, duo}
}

func main() {
  fmt.Println(Split([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "k"}, 3))
}

