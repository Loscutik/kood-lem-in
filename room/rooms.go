package room

type AntFarm struct{ 
	Rooms map[string]*Room
	Start,End *Room
}

type Room struct {
	Name  string
	Links map[string]*Room
	x, y  int // coordinates
}

func (f *AntFarm) AddRoom(name string, x, y int) (error) {
	//TODO check doubling rooms

	f.Rooms[name]= &Room{Name: name, x: x, y: y}
	
	return nil
}

func isContains(rooms map[string]*Room, r *Room)bool{
	return false
}

func (r *Room) AddLink(neibor *Room) {}
