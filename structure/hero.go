package structure

import "fmt"

type Hero struct {
	Id         int
	Name       string
	location   *Node
	TakenQuest *quest
}

type HeroFunctions interface {
	GetName() string
	SetLocation(array []Node, location string)
	GetLocation() Node
	Move(array []Node, where string) Node
	WalkTo(array []Node, where string, herolog *HeroLog) bool
}

func CreateHero(name string) Hero {
	new_hero := &Hero{Id: 0, Name: name}
	return *new_hero
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) GetLocation() Node {
	return *hero.location
}

func (hero *Hero) GetLocationAsString() string {
	return hero.location.Name
}

func (hero *Hero) GetNearbyLocations(array []Node) []*Node {
	loc := hero.GetLocation()
	nearby := make([]*Node, 0)
	for i, e := range array {
		if loc.Id == e.Id {
			if len(array[i].Next) > 0 {
				for _, e2 := range array[i].Next {
					nearby = append(nearby, e2)
				}
			}
		}

	}
	return nearby
}

func (hero *Hero) GetNearbyLocationsAsStringsOf(array []Node, of string) []string {
	node := FindElementByName(array, of)
	nearby := make([]string, 0)
	for i, e := range array {
		if node.Id == e.Id {
			if len(array[i].Next) > 0 {
				for _, e2 := range array[i].Next {
					nearby = append(nearby, e2.Name)
				}
			}
		}

	}
	return nearby
}

func (hero *Hero) GetNearbyLocationsAsStrings(array []Node, heroLog *HeroLog) []string {
	loc := hero.GetLocation()
	nearby := make([]string, 0)
	for i, e := range array {
		if loc.Id == e.Id {
			if len(array[i].Next) > 0 {
				for _, e2 := range array[i].Next {
					nearby = append(nearby, e2.Name)
				}

			}
		}
	}
	heroLog.HLog(hero, 0, "GetNearbyLocationsAsStrings", hero.GetLocationAsString())
	fmt.Println("Log: ", heroLog.Function, "++")
	return nearby
}

func (hero *Hero) SetLocation(array []Node, location string) Node {
	loc := FindElementByName(array, location)
	hero.location = &loc
	return array[loc.Id]
}

func (hero *Hero) Move(array []Node, where string) Node {
	ret := FindElementByName(array, where)
	if ret.Id >= 0 {
		hero.SetLocation(array, ret.Name)
	}
	return ret
}

func PrintAll(array []*Node, location string) {
	fmt.Println("Oroginal location: ", location)
	fmt.Println("Nearby locations: ")
	for i, e := range array {
		fmt.Println("		index:", i, " -> ", e.Name)
	}
	fmt.Println("...............")
}

func (hero *Hero) WalkTo(array []Node, where string, herolog *HeroLog) bool {
	current := ""
	tmp := hero.GetNearbyLocations(array)
	for i := range tmp {
		current = tmp[i].Name
		if tmp[i].Name != where {
			tmp2 := hero.GetNearbyLocationsAsStringsOf(array, tmp[i].Name)
			for _, el := range tmp2 {
				if el == where {
					hero.SetLocation(array, where)
					return true
				}
			}
			fmt.Println("RR ", current, " ", i, " ", hero.GetLocationAsString(), hero.GetNearbyLocationsAsStrings(array, herolog))
		}
	}
	if current == where {
		hero.SetLocation(array, where)
		return true
	} else {
		fmt.Println("No road ", hero.GetLocation())
		return false
	}
}
