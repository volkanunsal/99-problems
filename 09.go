// Pack consecutive duplicates of list elements into sublists.

// ```go
// pack([]interface{}{5, 5, 5, 4, 3, 3})
// [[5 5 5] [4] [3 3]]
// ```

func pack(arg []interface{}) [][]interface{} {
  var ret [][]interface{}
  // Index of the return slice
  j := 0
  // Index of the arg slice
  i := 1

  // Insert the first element of the argument slice into the return slice
  ret = append(ret, []interface{}{arg[0]})

  // Walk through the rest of the slice and append its members into `ret`
  for i < len(arg) and len(arg) > 1 {
    item := arg[i]

    // When the current item is different from the last one, append an empty interface
    // into ret –– meaning move to the next element.
    if arg[i-1] != item {
      var empty []interface{}
      ret = append(ret, empty)
      j += 1
    }

    // Append current item at the jth position of the ret slice.
    ret[j] = append(ret[j], item)
    i += 1
  }

  return ret
}

