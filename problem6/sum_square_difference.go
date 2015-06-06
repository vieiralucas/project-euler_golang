// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import (
  "flag"
  "strconv"
  "math"
  "fmt"
)

func main() {
  flag.Parse()
  arg := flag.Arg(0)
  number, err := strconv.Atoi(arg)
  if err != nil {
    panic(err)
  }

  sumOfSquares, squareOfSum := 0, 0

  for i := 1; i <= number; i++ {
      sumOfSquares += int(math.Pow(float64(i), 2))
      squareOfSum += i
  }

  squareOfSum = int(math.Pow(float64(squareOfSum), 2))

  fmt.Println(int(math.Abs(float64(sumOfSquares - squareOfSum))))
}
