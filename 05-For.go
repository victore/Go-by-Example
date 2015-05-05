package main

import "fmt"

func main() {
  i := 1
  for i <= 3 {
    fmt.Println(i)
    i = i + 1
  }

  // A classic initial/condition/after for loop.
  for j := 7; j <= 9; j++ {
    fmt.Println(j)
  }

  for {
    fmt.Println("loope!")
    break
  }
}
