package main

import (
  "os"
	"net/http"
  "encoding/json"
)

func getJson(url string, target interface{}) error {
    r, err := http.Get(url)
    if err != nil {
        return err
    }
    defer r.Body.Close()
    return json.NewDecoder(r.Body).Decode(target)
}

type Response struct {
    thing string
}

func main() {
	url := os.Args[1]
  resp := &Response{} // or &Foo{}
  getJson(url, resp)
  println(resp.thing)

}
