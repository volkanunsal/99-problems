// Find the last but one element of a list.
//
// > myButLast([]interface{}{1, 2, 3, 4})
// 3
// > myButLast([]interface{}{"x","y","z"})
// 'y'
// > myButLast([]interface{}{})
// nil
// > myButLast([]interface{}{"x"})
// nil
package main

import (
  "fmt"
)

func main() {
  fmt.Println(myButLast([]interface{}{1, 2, 3, 4}))
  fmt.Println(myButLast([]interface{}{"x","y","z"}))
  fmt.Println(myButLast([]interface{}{}))
  fmt.Println(myButLast([]interface{}{"x"}))
}

func myButLast(arr []interface{}) interface{} {
  if length := len(arr); length == 0 || length == 1 {
    return nil
  }
  return arr[len(arr) - 2]
}
