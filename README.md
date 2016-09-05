# Find the last element of a list.
__Difficulty: :white_check_mark:__

## Go

```Go
> myLast([]interface{}{1, 2, 3, 4})
4
> myLast([]interface{}{"x","y","z"})
'z'
> myLast([]interface{}{})
nil
```

*Solution:*

```go
func myLast(arr []interface{}) interface{} {
  if length := len(arr); length == 0 {
    return nil
  }
  return arr[len(arr) - 1]
}
```

[Playground](https://play.golang.org/p/Zb0Xsvm718)

# Find the last but one element of a list.
__Difficulty: :white_check_mark:__


## Go

```go
> myButLast([]interface{}{1, 2, 3, 4})
3
> myButLast([]interface{}{"x","y","z"})
'y'
> myButLast([]interface{}{})
nil
> myButLast([]interface{}{"x"})
nil
```
*Solution:*

```go
func myButLast(arr []interface{}) interface{} {
  if length := len(arr); length == 0 || length == 1 {
    return nil
  }
  return arr[len(arr) - 2]
}
```

[Playground](https://play.golang.org/p/PYXaCAI-XR)

# Find the Kth element of a list.
__Difficulty: :white_check_mark:__

## Go

```go
> elementAt([]interface{}{1, 2, 3}, 2)
2
> elementAt([]interface{}{"haskell"}, 5)
'e'
```

*Solution:*

```go
func elementAt(arr []interface{}, ix int) interface{} {
  if ix < 1 || ix > len(arr) {
    return errors.New(fmt.Sprintf("Invalid index: %d", ix))
  }
  return arr[ix - 1]
}

func main() {
  fmt.Println(elementAt([]interface{}{1, 2, 3}, -1))
  fmt.Println(elementAt([]interface{}{1, 2, 3}, 2))
  old := "haskell"
  new := make([]interface{}, len(old))
  for i, v := range old {
    new[i] = string(v)
  }
  fmt.Println(elementAt(new, 5))
}
```

Casting from string to `[]interface{}` automatically is not possible. The workaround is to create an empty interface array and copy the contents of the string into it one by one.


[Playground](https://play.golang.org/p/73849rRz7W)


# Find the number of elements of a list.
__Difficulty: :white_check_mark:__

## Go

```go
> myLength([]interface{}{123, 456, 789})
3
> myLength([]interface{}{"Hello, world!"})
13
```

*Solution:*

```go
func myLength1(arr []int) interface{} {
  return len(arr)
}

func myLength2(arr string) interface{} {
  return len(arr)
}

func main() {
  fmt.Println(myLength1([]int{123, 456, 789}))
  fmt.Println(myLength2("haskell"))
}
```

# Reverse a list.
__Difficulty: :white_check_mark:__
# Find out whether a list is a palindrome.
__Difficulty: :white_check_mark:__
# Flatten a nested list structure.
__Difficulty: :white_check_mark::white_check_mark:__
# Eliminate consecutive duplicates of list elements.
__Difficulty: :white_check_mark::white_check_mark:__
# Pack consecutive duplicates of list elements into sublists.
__Difficulty: :white_check_mark::white_check_mark:__
# Run-length encoding of a list.
__Difficulty: :white_check_mark:__
# Modified run-length encoding.
__Difficulty: :white_check_mark:__
# Decode a run-length encoded list.
__Difficulty: :white_check_mark::white_check_mark:__
# Run-length encoding of a list (direct solution).
__Difficulty: :white_check_mark::white_check_mark:__
# Duplicate the elements of a list.
__Difficulty: :white_check_mark:__
# Duplicate the elements of a list a given number of times.
__Difficulty: :white_check_mark::white_check_mark:__
# Drop every Nth element from a list.
__Difficulty: :white_check_mark::white_check_mark:__
# Split a list into two parts.
__Difficulty: :white_check_mark:__
# Extract a slice from a list.
__Difficulty: :white_check_mark::white_check_mark:__
# Rotate a list N places to the left.
__Difficulty: :white_check_mark::white_check_mark:__
# Remove the Kth element from a list.
__Difficulty: :white_check_mark:__
# Insert an element at a given position into a list.
__Difficulty: :white_check_mark:__
# Create a list containing all integers within a given range.
__Difficulty: :white_check_mark:__
# Extract a given number of randomly selected elements from a list.
__Difficulty: :white_check_mark::white_check_mark:__
# Lotto: Draw N different random numbers from the set 1..M.
__Difficulty: :white_check_mark:__
# Generate a random permutation of the elements of a list.
__Difficulty: :white_check_mark:__
# Generate the combinations of K distinct objects chosen from the N elements of a list.
__Difficulty: :white_check_mark::white_check_mark:__
# Group the elements of a set into disjoint subsets.
__Difficulty: :white_check_mark::white_check_mark:__
# Sorting a list of lists according to length of sublists.
__Difficulty: :white_check_mark::white_check_mark:__
# Determine whether a given integer number is prime.
__Difficulty: :white_check_mark::white_check_mark:__
# Determine the greatest common divisor of two positive integer numbers.
__Difficulty: :white_check_mark::white_check_mark:__
# Determine whether two positive integer numbers are coprime.
__Difficulty: :white_check_mark:__
# Calculate Euler's totient function phi(m).
__Difficulty: :white_check_mark::white_check_mark:__
# Determine the prime factors of a given positive integer.
__Difficulty: :white_check_mark::white_check_mark:__
# Determine the prime factors of a given positive integer (2).
__Difficulty: :white_check_mark::white_check_mark:__
# Calculate Euler's totient function phi(m) (improved).
__Difficulty: :white_check_mark::white_check_mark:__
# Compare the two methods of calculating Euler's totient function.
__Difficulty: :white_check_mark:__
# A list of prime numbers.
__Difficulty: :white_check_mark:__
# Goldbach's conjecture.
__Difficulty: :white_check_mark::white_check_mark:__
# A list of Goldbach compositions.
__Difficulty: :white_check_mark::white_check_mark:__
# Truth tables for logical expressions.
__Difficulty: :white_check_mark::white_check_mark:__
# Truth tables for logical expressions (2).
__Difficulty: :white_check_mark:__
# Truth tables for logical expressions (3).
__Difficulty: :white_check_mark::white_check_mark:__
# Gray code.
__Difficulty: :white_check_mark::white_check_mark:__
# Huffman code.
__Difficulty: :white_check_mark::white_check_mark::white_check_mark:__
# Construct completely balanced binary trees.
__Difficulty: :white_check_mark::white_check_mark:__
# Symmetric binary trees.
__Difficulty: :white_check_mark::white_check_mark:__
# Binary search trees (dictionaries).
__Difficulty: :white_check_mark::white_check_mark:__
# Generate-and-test paradigm.
__Difficulty: :white_check_mark::white_check_mark:__
# Construct height-balanced binary trees.
__Difficulty: :white_check_mark::white_check_mark:__
# Construct height-balanced binary trees with a given number of nodes.
__Difficulty: :white_check_mark::white_check_mark:__
# Count the leaves of a binary tree.
__Difficulty: :white_check_mark:__
# Collect the leaves of a binary tree in a list.
__Difficulty: :white_check_mark:__
# Collect the internal nodes of a binary tree in a list.
__Difficulty: :white_check_mark:__
# Collect the nodes at a given level in a list.
__Difficulty: :white_check_mark:__
# Construct a complete binary tree.
__Difficulty: :white_check_mark::white_check_mark:__
# Layout a binary tree (1).
__Difficulty: :white_check_mark::white_check_mark:__
# Layout a binary tree (2).
__Difficulty: :white_check_mark::white_check_mark:__
# Layout a binary tree (3).
__Difficulty: :white_check_mark::white_check_mark::white_check_mark:__
# A string representation of binary trees.
__Difficulty: :white_check_mark::white_check_mark:__
# Preorder and inorder sequences of binary trees.
__Difficulty: :white_check_mark::white_check_mark:__
# Dotstring representation of binary trees.
__Difficulty: :white_check_mark::white_check_mark:__
# Count the nodes of a multiway tree.
__Difficulty: :white_check_mark:__
# Tree construction from a node string.
__Difficulty: :white_check_mark::white_check_mark:__
# Determine the internal path length of a tree.
__Difficulty: :white_check_mark:__
# Construct the postorder sequence of the tree nodes.
__Difficulty: :white_check_mark:__
# Lisp-like tree representation.
__Difficulty: :white_check_mark::white_check_mark:__
# Conversions.
__Difficulty: :white_check_mark::white_check_mark::white_check_mark:__
# Path from one node to another one.
__Difficulty: :white_check_mark::white_check_mark:__
# Cycle from a given node.
__Difficulty: :white_check_mark:__
# Construct all spanning trees.
__Difficulty: :white_check_mark::white_check_mark:__
# Construct the minimal spanning tree.
__Difficulty: :white_check_mark::white_check_mark:__
# Graph isomorphism.
__Difficulty: :white_check_mark::white_check_mark:__
# Node degree and graph coloration.
__Difficulty: :white_check_mark::white_check_mark:__
# Depth-first order graph traversal.
__Difficulty: :white_check_mark::white_check_mark:__
# Connected components.
__Difficulty: :white_check_mark::white_check_mark:__
# Bipartite graphs.
__Difficulty: :white_check_mark::white_check_mark:__
# Eight queens problem
__Difficulty: :white_check_mark::white_check_mark:__
# Knight's tour.
__Difficulty: :white_check_mark::white_check_mark:__
# Von Koch's conjecture.
__Difficulty: :white_check_mark::white_check_mark::white_check_mark:__
# An arithmetic puzzle.
__Difficulty: :white_check_mark::white_check_mark::white_check_mark:__
# Generate K-regular simple graphs with N nodes.
__Difficulty: :white_check_mark::white_check_mark::white_check_mark:__
# English number words.
__Difficulty: :white_check_mark::white_check_mark:__
# Sudoku. (alternate solution)
__Difficulty: :white_check_mark::white_check_mark:__
# Nonograms.
__Difficulty: :white_check_mark::white_check_mark::white_check_mark:__
# Crossword puzzle.
__Difficulty: :white_check_mark::white_check_mark::white_check_mark:__
