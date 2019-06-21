package main

import (
  "fmt"
)

type fn func() int {}

func do() (int) {
  return 1
}

func dodo(f func() int) (int) {
  return f()
}

func dododo(f func(func() int) int, g func() int) (int) {
  // disgusting
  return f(g)
}


func main(){
  dododo(dodo, do)
  fmt.Println("ok")
}
