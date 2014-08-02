package main

import (
  "fmt"
  "github.com/gillesdemey/hn/lib"
)

func main() {

  posts, err := lib.LoadFrontPage()

  if err != nil {
    panic(err)
  }

  for i, post := range posts {
    i++
    fmt.Printf("\n%02d: %s", i, post.Title)
  }

}