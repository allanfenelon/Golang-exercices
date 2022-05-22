package main

type Queue struct {
	Head *Node
}

type Node struct {
	nameN string
	Next  *Node
}

func (q *Queue) Enqueue(name string) {
	node := Node{nameN: name}
	if q.Head == nil {
		q.Head = &node
	} else {
		aux := q.Head
		for aux.Next != nil {
			aux = aux.Next
		}
		aux.Next = &node
	}
}

func (q *Queue) Dequeue() string {
	if q.Head == nil {
		return ""
	}

	node := q.Head
	q.Head = q.Head.Next
	return node.nameN
}
