// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
  "strconv"
  "fmt"
  "sort"
)

func main() {
  palindromes := []int{}

  for left := 999; left >= 0; left-- {
    for right := 999; right >= 0; right-- {
      if isPalidrome(left * right) {
        palindromes = append(palindromes, left * right)
      }
    }
  }

  sort.Ints(palindromes)
  fmt.Println(palindromes[len(palindromes) - 1])
}

func isPalidrome(n int) bool {
  return n == revertInt(n)
}

func revertInt(n int) int {
  n, err := strconv.Atoi(revertString(strconv.Itoa(n)))
  if (err != nil) {
    panic(err)
  }

  return n
}

func revertString(s string) string {
  r := ""
  for _, c := range s {
    r = string(c) + r
  }

  return r
}
