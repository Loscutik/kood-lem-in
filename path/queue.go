package path

import "lemin/room"

/*
adds a room at the tail of queue. If queue's tail is nil, creates a new queue that contains only the given room
*/
func (q *queue) add(r *room.Room) {
	node := queuesNode{room: r, next: nil}
	if q.tail == nil {
		q.head = &node
	} else {
		q.tail.next = &node
	}
	q.tail = &node
}

func (q *queue) isEmpty() bool {
	if q.head == nil {
		return true
	}
	return false
}
