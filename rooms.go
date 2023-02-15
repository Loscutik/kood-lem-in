package main

type room struct {
	name  string
	links map[string]*room
	x, y  int // coordinates
}

func New(name string, x, y int) (*room, error) {
	//TODO check doubling rooms
	r := room{name: name, x: x, y: y}
	
	return &r,nil
}

func (r *room) AddLink(neibor *room) {}
