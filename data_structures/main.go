package main

import "fmt"

func main() {
	queue := Queue{}
	queue.Enqueue("Pedro")
	queue.Enqueue("Lucas")
	queue.Enqueue("Fernando")
	queue.Enqueue("Joaquim")
	queue.Enqueue("Tantas")

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

}
