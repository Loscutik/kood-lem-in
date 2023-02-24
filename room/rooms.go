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
	room:=Room{Name: name, x: x, y: y}
	//TODO check doubling rooms if isContains(f.Rooms,&room) return fmt.Errorf("double room...")

	f.Rooms= append(f.Rooms, &room)
	
	return nil
}

func isContains(rooms []*Room, room *Room)bool{
	for _,r:= range rooms{
		if r==room{
			return true
		}
	}
	return false
}

func (r *Room) AddLink(neibor *Room) (error) {
	//TODO check doubling rooms

	r.Links= append(r.Links, neibor)
	return nil
}

