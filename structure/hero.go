package structure

import "fmt"

type hero struct {
	Id       int
	Name     string
	location *Node
	//...
}

type hero_functions interface {
	GetName() string
	SetLocation(array []Node, location string)
	GetLocation() Node
	Move(array []Node, where string) Node
	GoForward(array []Node, where string)
}

func CreateHero(name string) hero {
	new_hero := &hero{Id: 0, Name: name}
	return *new_hero
}

func (hero *hero) GetName() string {
	return hero.Name
}

func (hero *hero) GetLocation() Node {
	return *hero.location
}

func (hero *hero) GetNearbyLocations(array []Node) {
	loc := hero.GetLocation()
	for i, e := range array {
		if loc.Id == e.Id {
			fmt.Println("Nearby locs: ", array[i].Next)
		}
	}

}

func (hero *hero) SetLocation(array []Node, location string) Node {
	loc := FindElementByName(array, location)
	hero.location = &loc
	return array[loc.Id]
}

func (hero *hero) Move(array []Node, where string) Node {
	ret := FindElementByName(array, where)
	if ret.Id >= 0 {
		hero.SetLocation(array, ret.Name)
	}
	return ret
}

func (hero *hero) WalkTo(array []Node, where string) {
	ret := FindElementByName(array, where)
	id := hero.GetLocation().Id
	for ret.Id >= 0 {
		if array[ret.Id].Id > hero.GetLocation().Id {
			id++
			hero.SetLocation(array, array[id].Name)
		} else {
			id--
			hero.SetLocation(array, array[id].Name)
		}
		if hero.GetLocation().Name == where {
			break
		}
	}
}
