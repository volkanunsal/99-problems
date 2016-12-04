# Find out whether a list is a palindrome.
__Difficulty: :white_check_mark:__


## Go

```go
> isPalindromeInt([1, 2, 3])
false
> isPalindromeInt([1, 2, 3, 2, 1])
true
> isPalindromeString("Hello, world!")
false
> isPalindromeString("abcba")
true
```

*Solution:*

```go

// Small cases
//
//   []         => true
//   [1, 1]     => true
//   [1, 1, 2]  => false
//   [1, 1, 2, 1]  => false
//
func isPalindromeInt(arr []int) bool {

  if len(arr) == 0 || len(arr) == 1 {
    return true
  }
  if len(arr) == 2 {
    return arr[0] == arr[1]
  }

  // Last index
  lix := len(arr) - 1

  // Mid point index
  mid := lix / 2

  if len(arr) % 2 != 0 {
    mid = mid + 1
  }

  for i, v := range arr[:mid] {
    if arr[lix - i] != v {
      return false
    }
  }
  return true
}
```


