package queue

import (
	"fmt"
	"log"
	"os"
)

func Test() {
	var newQueue = InitQueue(8)

	var input int
	for {
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case input == 0:
			if !newQueue.IsEmpty() {
				deletedData := newQueue.Delete()
				fmt.Println("Head:", newQueue.Head)
				fmt.Println("Tail:", newQueue.Tail)
				fmt.Println("Deleted Data:", deletedData)
			} else {
				fmt.Println("Queue is empty")
			}
		case input == -1:
			newQueue.ListReverseElement(newQueue.Head)
			os.Exit(0)
		default:
			if !newQueue.IsFull() {
				newQueue.Add(input)
			} else {
				fmt.Println("Queue is Full.")
			}
		}
	}

}
