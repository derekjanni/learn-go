package main

import (
  "fmt"
)

type LinkedList struct {
  Head Node
}

func NewLinkedListFromNode(node Node) *LinkedList {
    return &LinkedList{Head: node}
}

type Node struct {
  Data string
  Next *Node
}

func ( link *LinkedList ) append ( node Node ) {
    if link.Head.Next == (&Node{}) {
      link.Head.Next = &node

    } else {
      last := &link.Head
  		for last.Next != nil {
  			last = last.Next
  		}
    		last.Next = &node
      }
    }

func ( link *LinkedList ) show() {
  for node := &link.Head; node != nil; node = node.Next {
    fmt.Println(node)
  }
}

func main() {
  first_node := Node{
      Data: "i'm the head",
  }
  list := NewLinkedListFromNode(first_node)
  list.show()

  next_node := Node{
      Data: "i'm the second one",
  }
  list.append(next_node)
  list.show()

  third_node := Node{
      Data: "i'm the third",
  }
  list.append(third_node)
  list.show()

}
