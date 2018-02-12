package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
  num1 := 0
  num2 := 0
  count := 0

  return func() int {
    ret := 0

    switch count {
    case 0:
      num1 = 0
      num2 = 0
      ret = 0
    case 1:
      num1 = 0
      num2 = 1
      ret = 1
    default:
      ret = num1 + num2
      num1 = num2
      num2 = ret
    }

    count += 1

    return ret
  }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

