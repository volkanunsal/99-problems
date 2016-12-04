// Decode a run-length encoded list.
// -----------------------------------
//
// Given a run-length code list generated as specified in problem 11. Construct its uncompressed version.
//
// > decodeModified([]interface{}{[]interface{}{4, "a"}, "b", []interface{}{2, "c"}, []interface{}{2, "a"}, "d", []interface{}{4, "e"}})
// "aaaabccaadeeee"
//
// My strategy: iterate over args, checking the type of each element. If it's a string, append it into ret slice, and if it's a slice, generate a string of length equal to the int at the first position of the slice, and append that. Finally, join the ret slice into a string, and return it.
//
package main
import "strconv"
import "strings"
import "fmt"

func decodeModified(args []interface{}) string {
  var ret string

  // Walk through the args slice
  for _, item := range args {
    var s []string

    switch v := item.(type) {
    default:
      fmt.Printf("unexpected type %T", v)
    case int:
      s = strconv.Itoa(v)
    case string:
      s = v
    }
    ret = append(ret, s)
  }
  return strings.Join(ret, ",")
}

func main() {
  fmt.Println(decodeModified([]interface{}{[]interface{}{4, "a"}, "b", []interface{}{2, "c"}, []interface{}{2, "a"}, "d", []interface{}{4, "e"}}))
}
