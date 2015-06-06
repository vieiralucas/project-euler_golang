// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to x?

package main

import (
  "flag"
  "strconv"
  "fmt"
)

func main() {
  flag.Parse()
  arg := flag.Arg(0)
  x, err := strconv.Atoi(arg)
  if err != nil {
    panic(err)
  }

  candidate := x

  for !isEvenlyDivisableByAll(candidate, x) {
    candidate += x
  }

  fmt.Printf("The smallest positive number that is evenly divisible by all of the numbers from 1 to %d is %d\n", x, candidate)
}

func isEvenlyDivisableByAll(n int, all int) bool {
  for i := all; i > 0; i-- {
    if n % i != 0 {
      return false
    }
  }

  return true
}
