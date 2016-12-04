# Eliminate consecutive duplicates of list elements.
__Difficulty: :white_check_mark::white_check_mark:__

```go
> compress([]interface{}{5, 5, 5, 4, 3, 3})
[5 4 5]
```

```go
func compress(arg []interface{}) []interface{} {
  var ret []interface{}
  // Index of arg slice.
  i := 1

  // Insert the first element of arg into ret
  ret = append(ret, arg[0])

  // Walk through the rest of the arg slice and insert the sequentially unique elements
  // into ret.
  for i < len(arg) && len(arg) > 1 {
      item := arg[i]
      if arg[i-1] != item {
        ret = append(ret, item)
      }
      i = i + 1
    }
  return ret
}
```


