package queue

import "fmt"

type Queue struct {
	Head        *Node
	Tail        *Node
	currentSize int
	size        int
}

type Node struct {
	Data int
	Next *Node
}

func (q *Queue) Add(num int) {
	newNode := Node{num, &Node{}}
	if q.Head == nil {
		fmt.Println("initialize queue")
		q.Head = &newNode
		q.Tail = &newNode
		q.Head.Next = &newNode
	} else if q.Head == q.Tail {
		q.Head.Next = &newNode
		q.Tail = &newNode
	} else {
		q.Tail.Next = &newNode
		q.Tail = &newNode
		q.Tail.Next = nil
	}
	q.currentSize++
}

func (q *Queue) Delete() int {
	deletedNode := q.Head
	q.Head = deletedNode.Next
	if q.Head.Next == nil {
		q.Head = nil
		q.Tail = nil
	}
	q.currentSize--
	return deletedNode.Data
}

func InitQueue(size int) Queue {
	newQueue := Queue{size: size, currentSize: 0}
	return newQueue
}

func (q *Queue) FindNode(index int) *Node {
	currentNode := q.Head
	for i := 0; i < index; i++ {
		currentNode = currentNode.Next
	}
	return currentNode
}

func (q *Queue) IsFull() bool {
	return q.currentSize == q.size
}
func (q *Queue) IsEmpty() bool {
	return q.Head == nil && q.Tail == nil
}

func (q *Queue) ListAllElement(node *Node) *Node {
	if node.Next == nil {
		fmt.Println(node.Data)
		return node
	}
	fmt.Println(node.Data)
	return q.ListAllElement(node.Next)
}

func (q *Queue) ListReverseElement(node *Node) {
	if node.Next == nil {
		fmt.Println("Done")
	} else {
		q.ListReverseElement(node.Next)
		fmt.Println(node.Data)
	}
}
func (q Queue) String() string {
	return fmt.Sprintf("%v", q.Tail)
}
