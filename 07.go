# Flatten a nested list structure.
__Difficulty: :white_check_mark::white_check_mark:__

Transform a list, possibly holding lists as elements into a `flat' list by replacing each list with its elements (recursively).

```go
> flatten([]int{5})
[5]
> flatten (List [Elem 1, List [Elem 2, List [Elem 3, Elem 4], Elem 5]])
> flatten([]interface{}{1,[]interface{}{2,[]interface{}{3,4}, 5}})
[1,2,3,4,5]
> flatten ([]interface{}{})
[]
```

<!-- TODO: runtime type inspection with `reflect` -->

```go
func flatten(arr []interface{}) []interface{} {
  if len(arr) == 0 || len(arr) == 1 {
    return arr
  }
  ret := make([]interface, len(arr))

  i := 0
  for {
    if len(arr) > 0 {
      x, a := arr[0], arr[1:]
      arr = a
      ret = append(ret, x)
      i = i + 1
    } else {
      return ret
    }
  }
}
```

