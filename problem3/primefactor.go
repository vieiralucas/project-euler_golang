package main

import (
  "flag"
  "strconv"
  "fmt"
  "math"
)

func main() {
  flag.Parse()
  arg := flag.Arg(0)
  number, err := strconv.Atoi(arg)
  if err != nil {
    panic(err)
  }

  var left, right int = findFactors(number)
  primes := []int{left}

  for (!isPrime(right)) {
    left, right = findFactors(right)
    primes = append(primes, left)
  }

  primes = append(primes, right)

  fmt.Printf("Prime factors: %+v\nThe largest one: %d\n", primes, primes[len(primes) - 1])
}

func findFactors(n int) (first int, second int) {
  for i := 2; i <= n / 2; i++ {
    if n % i == 0 && isPrime(i) {
      return i, n / i
    }
  }

  return -1, -1
}

func isPrime(n int) bool {
  if (n <= 0) { return false }
  if (n == 2) { return true }
  var sqrt int = int(math.Ceil(math.Sqrt(float64(n))))

  for candidate := 2; candidate <= sqrt; candidate++ {
    if n % candidate == 0 {
      return false
    }
  }

  return true
}
