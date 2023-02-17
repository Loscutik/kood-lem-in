package main

type antFarm struct{ 
	rooms map[string]*room
	start,end *room
}

type room struct {
	name  string
	links map[string]*room
	x, y  int // coordinates
}

func (f *antFarm) AddRoom(name string, x, y int) (error) {
	//TODO check doubling rooms

	f.rooms[name]= &room{name: name, x: x, y: y}
	
	return nil
}

func isContains(rooms map[string]*room, r *room)bool{
	return false
}

func (r *room) AddLink(neibor *room) {}
