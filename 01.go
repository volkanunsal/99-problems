// Find the last element of a list.
//
// > myLast([]interface{}{1, 2, 3, 4})
// 4
// > myLast([]interface{}{"x","y","z"})
// 'z'
// > myLast([]interface{}{})
// nil
//
package main

import (
  "fmt"
)

func myLast(arr []interface{}) interface{} {
  if length := len(arr); length == 0 {
    return nil
  }
  return arr[len(arr) - 1]
}

func main() {
  fmt.Println(myLast([]interface{}{1, 2, 3, 4}))
  fmt.Println(myLast([]interface{}{"x","y","z"}))
  fmt.Println(myLast([]interface{}{}))
}
