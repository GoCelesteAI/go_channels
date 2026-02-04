package main

import (
  "fmt"
  "time"
)

func main() {
  // Create an unbuffered channel
  messages := make(chan string)

  // Send a message in a goroutine
  go func() {
    messages <- "Hello from goroutine!"
  }()

  // Receive the message
  msg := <-messages
  fmt.Println("Received:", msg)

  // Channel with goroutine communication
  done := make(chan bool)

  go func() {
    fmt.Println("Working...")
    time.Sleep(1 * time.Second)
    fmt.Println("Done!")
    done <- true
  }()

  // Wait for completion
  <-done
  fmt.Println("Main: received done signal")
}

