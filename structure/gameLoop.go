package structure

import "fmt"

type Game struct {
	gameMap    *[]Node //pointer to game map
	gameHero   *Hero   //pointer to hero
	gameQuests *[]Quest
}

type GameInt interface {
	NewGame(heroName string) *Game
	GetMap() *[]Node
	SetMap(arr []Node)
	GetHero() *Hero
}

func NewGame(heroName string) *Game {
	game := new(Game)
	arr := CreateArrOfNodes()
	hero := CreateHero(heroName)
	game.gameMap = &arr
	game.gameHero = &hero
	game.gameQuests = &AllQuests
	return game
}

func (game *Game) SetMap(arr []Node) {
	game.gameMap = &arr
}

func (game *Game) SetHero(hero *Hero) {
	game.gameHero = hero
}

func (game *Game) SetQuests(quests []Quest) {
	game.gameQuests = &quests
}

func (game *Game) GetMap() []Node {
	return *game.gameMap
}

func (game *Game) GetHero() *Hero {
	return game.gameHero
}

func (game *Game) GetQuests() []Quest {
	return *game.gameQuests
}

func GameLoop(game *Game) {
	cmd := ""
	for {
		fmt.Print("Type command: ")
		fmt.Scanln(&cmd)
		if cmd == "quit" {
			break
		}
		switch cmd {
		case "look":
			fmt.Println("Current location: ", game.gameHero.GetLocationAsString())
			fmt.Println("Nearby locations: ")
			fmt.Println(game.gameHero.GetNearbyLocationsAsStrings(*game.gameMap))
		case "quests":
			fmt.Println(game.gameHero.GetAllQuestsInCurrentLocationAsStrings(*game.gameMap))
		case "go":
			fmt.Print("to where?: ")
			location := ""
			fmt.Scanln(&location)
			found := FindElementByName(*game.gameMap, location)
			game.gameHero.WalkTo(*game.gameMap, found.Name)
			fmt.Println("Current location: ", game.gameHero.GetLocationAsString())
		default:
			continue
		}
	}
}
