// Run-length encoding of a list.
// --------------------------------
//
// Run-length encoding of a list. Use the result of problem P09 to implement the so-called run-length encoding data compression method. Consecutive duplicates of elements are encoded as lists (N E) where N is the number of duplicates of the element E.
//
// ```go
// encode([]interface{}{5, 5, 5, 4, 3, 3})
// [[3 5] [1 4] [2 3]]
// ```
func encode(arg []interface{}) [][]interface{} {
  // Create a slice of interfaces containing different items
  packed := pack(arg)
  // Eliminate duplicates
  compressed := compress(arg)

  var ret [][]interface{}

  // Walk through the sequentially unique elements and build an interface
  // of length of the duplicates of the item, and the item, and insert
  // it into the ret slice.
  for i, item := range compressed {
    ret = append(ret, []interface{}{len(packed[i]), item})
  }
  return ret
}
