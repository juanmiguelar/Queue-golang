package main

import "fmt"

// My first Queue
type Queue struct{
	items []int
}

// Enqueue this func append the item to the end of the queue
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue this func removes one item from the queue and returns it
func (q *Queue) Dequeue() int {
	// Identify which element do we need to remove
	itemToRemove := q.items[0]
	// [0 , 1 , 2] => [1 , 2] => slice[1:]
	q.items = q.items[1:]

	// Return the item that we just removed
	return itemToRemove
}

func main(){
	queue := Queue{}
	PrintWithCoolFormat(queue, "The very beggining")
	queue.Enqueue(100)
	PrintWithCoolFormat(queue, "Enqueue x1")
	queue.Enqueue(200)
	PrintWithCoolFormat(queue, "Enqueue x2")
	item := queue.Dequeue()
	PrintWithCoolFormat(queue, "Dequeue item x1 = " + fmt.Sprint(item))
	item = queue.Dequeue()
	PrintWithCoolFormat(queue, "Dequeue item x2 = " + fmt.Sprint(item))
}

func PrintWithCoolFormat(q Queue, detail string){
	fmt.Println("========================")
	fmt.Println(detail)
	fmt.Println("------------------------")
	fmt.Println(q)
	fmt.Println("------------------------")
	fmt.Println("========================")



}