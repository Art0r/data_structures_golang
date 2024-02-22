package data

import "fmt"

type Queue struct {
	arr []*node
}

func (queue *Queue) Head() *node {
	return queue.arr[len(queue.arr)-1]
}

func (queue *Queue) Rear() *node {
	return queue.arr[0]
}

func (queue *Queue) Push(value interface{}) {
	i := &node{Data: value, Next: nil}

	queue.arr = append(queue.arr, i)

	if len(queue.arr) > 1 {
		queue.arr[len(queue.arr)-2].Next = queue.arr[len(queue.arr)-1]
	}
}

func (queue *Queue) Pop() {
	queue.arr = queue.arr[:len(queue.arr)-1]
	queue.arr[len(queue.arr)-1].Next = nil
}

func (queue *Queue) Display() {
	for i := 0; i < len(queue.arr); i++ {
		fmt.Println("CURRENT: ", queue.arr[i].Data)

		if queue.arr[i].Next != nil{
			fmt.Println("NEXT: ", queue.arr[i].Next.Data)
		}

		fmt.Println("-----------------------")

	}
}
