package main

import "fmt"

func main() {
  // Unbuffered channel - blocks until received
  unbuffered := make(chan int)

  go func() {
    unbuffered <- 42
    fmt.Println("Sent to unbuffered")
  }()

  value := <-unbuffered
  fmt.Println("Received from unbuffered:", value)

  // Buffered channel - can hold values
  buffered := make(chan string, 3)

  // Send without blocking (up to capacity)
  buffered <- "first"
  buffered <- "second"
  buffered <- "third"
  fmt.Println("Sent 3 messages to buffered channel")

  // Receive values
  fmt.Println(<-buffered)
  fmt.Println(<-buffered)
  fmt.Println(<-buffered)

  // Check buffer capacity and length
  ch := make(chan int, 5)
  ch <- 1
  ch <- 2

  fmt.Println("Capacity:", cap(ch))
  fmt.Println("Length:", len(ch))
}

