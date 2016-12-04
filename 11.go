// Modified run-length encoding.
// -----------------------------
// https://wiki.haskell.org/99_questions/11_to_20
//
// Modify the result of problem 10 in such a way that if an element has no duplicates it is simply copied into the result list. Only elements with duplicates are transferred as (N E) lists.
//
// Example:
//
//  encodeModified([]interface{}{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"})
//  > [[4 A] B [2 C] [2 A] D [4 E]]
//
// My strategy: encode the input, search for 1-length slices, extract them out of their slice, as I insert them into the ret slice. The length > 1 slices will be inserted as they are.
package main
func encodeModified(args []interface{}) []interface{} {
  var ret []interface{}

  // Encoded slice will be of following form
  //
  //    [[n el]]
  //
  // ...where `n` equals the number of `el`s
  //
  encoded := encode(args)

  for n, el := range encoded {
    if n > 1 {
      ret = append(ret, el)
    } else {
      ret = append(ret, el[1])
    }
  }
  return ret
}

// func main() {
//   fmt.Println(encodeModified([]interface{}{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "e", "e", "e", "e"}))
// }
