package main

import "fmt"

func main() {
  // Range over channel
  ch := make(chan int)

  go func() {
    for i := 1; i <= 5; i++ {
      ch <- i * 10
    }
    close(ch)
  }()

  // Range automatically stops when channel closes
  fmt.Println("Values from channel:")
  for value := range ch {
    fmt.Println(value)
  }

  // Pipeline pattern
  numbers := generateNumbers(1, 5)
  squared := squareNumbers(numbers)

  fmt.Println("\nSquared values:")
  for result := range squared {
    fmt.Println(result)
  }
}

func generateNumbers(start, end int) chan int {
  out := make(chan int)
  go func() {
    for i := start; i <= end; i++ {
      out <- i
    }
    close(out)
  }()
  return out
}

func squareNumbers(in chan int) chan int {
  out := make(chan int)
  go func() {
    for n := range in {
      out <- n * n
    }
    close(out)
  }()
  return out
}

