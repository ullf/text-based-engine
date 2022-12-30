package structure

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Hero struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Location   *Node  `json:"location"`
	TakenQuest int    `json:"takenQuest"`
}

type HeroFunctions interface {
	GetName() string
	SetLocation(array []Node, location string)
	GetLocation() Node
	Move(array []Node, where string) Node
	WalkTo(array []Node, where string, questlog *QuestLog) bool
	Stat(array []Node) Hero
	ReStat() (*Hero, error)
}

func CreateHero(name string) Hero {
	new_hero := &Hero{Id: 0, Name: name}
	return *new_hero
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) GetLocation() Node {
	return *hero.Location
}

func (hero *Hero) GetLocationAsString() string {
	return hero.Location.Name
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

func (hero *Hero) GetNearbyLocationsAsStrings(array []Node) []string {
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

	//heroLog.HLog(hero, 0, "GetNearbyLocationsAsStrings", hero.GetLocationAsString())
	return nearby
}

func (hero *Hero) SetLocation(array []Node, location string) Node {
	loc := FindElementByName(array, location)
	hero.Location = &loc
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

func (hero *Hero) WalkTo(array []Node, where string) bool {
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
			//fmt.Println("RR ", current, " ", i, " ", hero.GetLocationAsString(), hero.GetNearbyLocationsAsStrings(array, herolog))
		}
	}
	if current == where {
		hero.SetLocation(array, where)
		//herolog.HLog(hero, 1, "WalkTo", hero.GetLocationAsString())
		return true
	} else {
		fmt.Println("No road ", hero.GetLocation())
		//herolog.HLog(hero, 1, "WalkTo", hero.GetLocationAsString())
		return false
	}
}

func (hero *Hero) Stat(array []Node) Hero {
	// Marshal the Hero struct into a JSON string
	jsonData, err := json.Marshal(hero)
	if err != nil {
		panic(err)
		//return *hero
	}

	// Write the JSON string to a file
	file, err := os.Create(hero.Name + ".json")
	if err != nil {
		// handle error
		return *hero
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		// handle error
		return *hero
	}

	return *hero
}

func (hero *Hero) ReStat() (*Hero, error) {
	// Open the file
	file, err := os.Open("mark.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the file into a byte slice
	data := make([]byte, 10000)
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}

	// Convert the byte slice to a string and remove any null characters
	jsonString := strings.Replace(string(data), "\x00", "", -1)

	// Unmarshal the string into a Hero struct
	var h Hero
	err = json.Unmarshal([]byte(jsonString), &h)
	fmt.Println("h: ", h)
	if err != nil {
		fmt.Println("h err: ", err)
		return nil, err
	}
	return &h, nil
}

func (hero *Hero) GetAllQuestsInCurrentLocationAsStrings(arr []Node) []string {
	myLocation := hero.GetLocation()
	fmt.Println("len: ", len(arr[myLocation.Id].Quests))
	localQuests := make([]string, 0)
	for _, e := range arr[myLocation.Id].Quests {
		fmt.Println("second: ", e.Name)
		localQuests = append(localQuests, e.Name)
	}
	return localQuests
}
