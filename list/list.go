package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")
	fmt.Println(l.Front().Value)
	fmt.Println(l.Back().Value)

	l.InsertAfter(6, l.Front())

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	fmt.Println("")
	l.MoveBefore(l.Front().Next(), l.Front())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")
	l.MoveToFront(l.Back())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	fmt.Println("")
	l.Init()
	fmt.Println(l.Len())
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
