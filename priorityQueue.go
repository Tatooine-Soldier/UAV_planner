package main

// import (
// 	"fmt"
// )

// //node struct
// type myNode struct {
// 	value    string
// 	edges    map[string]*myNode
// 	priority int
// }

// type Queue struct {
// 	queue []*myNode
// }

// //##PRIORITY QUEUE
// func (q *Queue) insert(n *myNode) {
// 	if n != nil {
// 		q.queue = append(q.queue, n)
// 	}
// 	fmt.Printf("%v", q.queue[len(q.queue)-1].value)
// }

// func (q *Queue) pop() *myNode {
// 	max := 0
// 	for index, val := range q.queue {
// 		if val.priority < q.queue[max].priority {
// 			max = index
// 		}
// 	}
// 	item := q.queue[max]
// 	q.queue[max] = nil
// 	return item
// }

// func main() {
// 	a := &myNode{
// 		value:    "a",
// 		edges:    map[string]*myNode{},
// 		priority: 4,
// 	}

// 	b := &myNode{
// 		value:    "b",
// 		edges:    map[string]*myNode{},
// 		priority: 3,
// 	}

// 	c := &myNode{
// 		value:    "c",
// 		edges:    map[string]*myNode{},
// 		priority: 2,
// 	}

// 	d := &myNode{
// 		value:    "d",
// 		edges:    map[string]*myNode{},
// 		priority: 1,
// 	}

// 	open := Queue{
// 		queue: []*myNode{},
// 	}
// 	open.insert(a)
// 	open.insert(b)
// 	open.insert(c)
// 	open.insert(d)

// 	myval := open.pop()
// 	fmt.Printf(" POPPED==>%v", myval.value)
// }
