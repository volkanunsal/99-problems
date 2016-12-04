// Find the Kth element of a list.
//
// > elementAt([]interface{}{1, 2, 3}, 2)
// 2
// > elementAt([]interface{}{"haskell"}, 5)
// 'e'
//
package main

import (
  "fmt"
  "errors"
)

func elementAt(arr []interface{}, ix int) interface{} {
  if ix < 1 || ix > len(arr) {
    return errors.New(fmt.Sprintf("Invalid index: %d", ix))
  }
  return arr[ix - 1]
}

func main() {
  fmt.Println(elementAt([]interface{}{1, 2, 3}, -1))
  fmt.Println(elementAt([]interface{}{1, 2, 3}, 2))
  old := "haskell"
  new := make([]interface{}, len(old))
  for i, v := range old {
    new[i] = string(v)
  }
  fmt.Println(elementAt(new, 5))
}
// Casting from string to `[]interface{}` automatically is not possible. The workaround is to create an empty interface array and copy the contents of the string into it one by one.
