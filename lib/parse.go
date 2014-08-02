package lib

import (
  "fmt"
  "encoding/json"
)

type Post struct {
  Id            int    `json:"id"`
  CreatedAt     int    `json:"created_at_i"`
  Type          string `json:"type"`
  Author        string `json:"author"`
  Title         string `json:"title"`
  Url           string `json:"url"`
  Text          string `json:"text"`
  Points        int    `json:"points"`
}

func ParsePosts(data []byte) ([]Post, error) {

  fmt.Printf("parsing posts...")

  var posts []Post

  // data is a byte array containing an array of posts
  err := json.Unmarshal(data, &posts)

  if err != nil {
    panic(err)
  }

  return posts, err
}