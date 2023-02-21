package path

import "lemin/room"

type queuesNode struct {
	room *room.Room
	next *queuesNode
}

type queue struct {
	head, tail *queuesNode
}

/*
adds a room at the tail of queue. If queue's tail is nil, creates a new queue that contains only the given room
*/
func (q *queue) pushToBack(r *room.Room) {
	node := queuesNode{room: r, next: nil}
	if q.tail == nil {
		q.head = &node
	} else {
		q.tail.next = &node
	}
	q.tail = &node
}

/*
adds a room on the head of queue. If queue is empty, creates a new queue that contains only the given room
*/
func (q *queue) pushToFront(r *room.Room) {
	node := queuesNode{room: r, next: q.head}
	q.head = &node
	if q.tail == nil {
		q.tail = &node
	}
}

/* 
deletes a node from the front of the queue and returns a pointer to the room kept in it
*/
func (q *queue) popFromFront() *room.Room {
	node := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	node.next = nil
	return node.room
}

/*
return true if there is no thing in the queue
*/
func (q *queue) isEmpty() bool {
	return q.head == nil 
}
/*
deletes all elements from the queue
*/
func (q *queue) clear() {
	if q.head == nil {
		q.tail = nil
		return
	}
	for q.head.next != nil {
		node := q.head.next
		q.head.next = nil
		q.head = node
	}
	q.head = nil
	q.tail = nil
}
