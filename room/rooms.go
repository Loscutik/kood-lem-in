package room

type AntFarm struct{ 
	Rooms []*Room
	Start,End *Room
}

type Room struct {
	Name  string
	Links []*Room
	x, y  int // coordinates
}

func (f *AntFarm) AddRoom(name string, x, y int) (error) {
	//TODO check doubling rooms

	f.Rooms= append(f.Rooms, &Room{Name: name, x: x, y: y})
	
	return nil
}

func isContains(rooms []*Room, r *Room)bool{
	return false
}

func (r *Room) AddLink(neibor *Room) {}
