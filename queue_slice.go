package main

import "fmt"

type Queue []any

func (q *Queue) Enqueue(data any) {
	*q = append(*q, data)
}

func (q *Queue) Dequeue() any {
	var val any = (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue) Peek() any {
	if q.IsEmpty() {
		return ""
	}
	return (*q)[0]
}

func (q *Queue) IsEmpty() bool {
	return len(*q) <= 0
}

func (q *Queue) Display() {
	for i := 0; i < len(*q); i++ {
		fmt.Println((*q)[i])
	}
}
