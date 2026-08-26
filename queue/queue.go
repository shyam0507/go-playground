package main

import "fmt"

type queue struct {
	items []int
}

type Queue interface {
	Enqueue(int)
	Dequeue() int
}

func (q *queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *queue) Dequeue() int {
	if len(q.items) == 0 {
		return 0
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func main() {

	q := &queue{}
	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	fmt.Println(q)

	///
	fmt.Println(q.Dequeue())
	fmt.Println(q)

}
