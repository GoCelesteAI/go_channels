package main

import "fmt"

func main() {
  ch := make(chan int, 3)

  // Send values
  ch <- 10
  ch <- 20
  ch <- 30

  // Close the channel
  close(ch)

  // Receive with ok idiom
  value, ok := <-ch
  fmt.Printf("Value: %d, Open: %v\n", value, ok)

  value, ok = <-ch
  fmt.Printf("Value: %d, Open: %v\n", value, ok)

  value, ok = <-ch
  fmt.Printf("Value: %d, Open: %v\n", value, ok)

  // After all values consumed
  value, ok = <-ch
  fmt.Printf("Value: %d, Open: %v\n", value, ok)

  // Producer pattern
  numbers := make(chan int)

  go func() {
    for i := 1; i <= 5; i++ {
      numbers <- i
    }
    close(numbers)
  }()

  // Consumer receives until closed
  for {
    n, ok := <-numbers
    if !ok {
      fmt.Println("Channel closed!")
      break
    }
    fmt.Println("Got:", n)
  }
}

