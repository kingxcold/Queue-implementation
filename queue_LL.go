package main

type QueueLL struct {
	List
}

func (q *QueueLL) Enqueue(data any) {
	q.Insert(data)
}

func (q *QueueLL) Dequeue() any {
	// remove element from list
	if q.IsEmpty() {
		return ""
	}
	val := q.Head.Value
	q.RemoveHead()
	return val
}

func (q *QueueLL) Peek() any {
	if q.IsEmpty() {
		return ""
	}
	return q.Head.Value
}

func (q *QueueLL) IsEmpty() bool {
	return q.Length <= 0
}

func (q *QueueLL) Display() {
	q.DisplayList()
}
