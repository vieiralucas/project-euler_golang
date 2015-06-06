package main

import "fmt"

var (
  term1 = 0
  term2 = 1
)

func main() {
  sum := 0

  for currentFib := 1; currentFib <= 4000000; currentFib = term1 + term2 {
      term1 = term2
      term2 = currentFib
      if currentFib % 2 == 0 {
        sum += currentFib
      }
  }

  fmt.Printf("The sum is: %d\n", sum)
}
