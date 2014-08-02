package lib

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func LoadFrontPage() ([]Post, error) {

  fmt.Println("loading...")

  buff, err := makeRequest("http://hn.algolia.com/rss.json")

  posts, err := ParsePosts(buff)

  return posts, err
}

func makeRequest(url string) ([]byte, error) {

  resp, err := http.Get(url)

  if err != nil {
    panic(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  return body, err
}