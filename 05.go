# Reverse a list.
__Difficulty: :white_check_mark:__


## Go

```go
> myReverseInt([1, 2, 3])
[3, 2, 1]
> myReverseString("Hello, world!")
"!dlrow ,olleH"
```
*Solution:*

```go
func myReverseInt(s []int) []int {
  ret := make([]int, len(s))
  for i := range s {
    ret[len(ret) - 1 - i] = s[i]
  }
  return ret
}
func myReverseString(s string) string {
  var ret = make([]byte, len(s))
  for i := range s {
    ret[len(ret) - 1 - i] = s[i]
  }
  return string(ret)
}
```

