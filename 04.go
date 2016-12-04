// Find the number of elements of a list.
// > myLength1([123, 456, 789])
// 3
// > myLength("Hello, world!")
// 13

func myLength1(arr []int) int {
  return len(arr)
}

func myLength2(arr string) int {
  return len(arr)
}

func main() {
  fmt.Println(myLength1([]int{123, 456, 789}))
  fmt.Println(myLength2("haskell"))
}

