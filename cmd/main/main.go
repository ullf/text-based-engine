package main

import (
	"fmt"

	//"text-based-engine/structure"

	"github.com/ullf/text-based-engine/structure"
)

func main() {
	game := structure.NewGame("mark")
	mapg := game.GetMap()
	hero := game.GetHero()
	mapg = structure.AppendElemToArr(mapg, "paris")
	mapg = structure.AppendElemToArr(mapg, "madrid")
	mapg = structure.AppendElemToArr(mapg, "berlin")
	mapg = structure.AppendElemToArr(mapg, "helsinki")

	berlin := structure.FindElementByName(mapg, "berlin")
	paris := structure.FindElementByName(mapg, "paris")

	//elem1 := structure.FindElementById(arr, 3)
	if berlin.Id >= 0 {
		mapg[1].AddChildToNodeByName(mapg, "madrid", &berlin, 0)
		//arr[3].AddChildToNodeByName(arr, "helsinki", &berlin, 0)
		mapg[3].AddChildToNodeByName(mapg, "helsinki", &paris, 0)
		//fmt.Println("Check : ", berlin.Id)
	} else {
		panic("error")
	}

	//structure.AddChildToNodeByName(arr, "20", &elem1)
	/*hero2, err := hero.ReStat()
	if err != nil {
		hero2 = structure.CreateHero("mark")
	} else {
		hero2.SetLocation(arr, hero2.GetLocationAsString())
	}*/
	hero.SetLocation(mapg, "helsinki")
	//hero.GetNearbyLocations(arr)
	fmt.Println("\nlocation: ", hero.GetLocationAsString())
	//bb := hero.WalkTo(mapg, "paris")
	//fmt.Println("bool: ", bb)
	fmt.Println("\nlocation: ", hero.GetNearbyLocationsAsStrings(mapg))
	//fmt.Println(arr)
	//structure.PrintAll(.Next, arr[1].Name)
	//structure.PrintAll(arr[0].Next, arr[0].Name)

	act := structure.CreateAction("lookup", 0)
	act2 := structure.CreateAction("walkto", 1)
	step := structure.CreateStep(hero, &act)
	step2 := structure.CreateStep(hero, &act2)
	steps := make([]structure.QuestStep, 0)

	act3 := structure.CreateAction("lookup", 0)
	act4 := structure.CreateAction("get all quests", 2)
	step3 := structure.CreateStep(hero, &act3)
	step4 := structure.CreateStep(hero, &act4)
	steps1 := make([]structure.QuestStep, 0)

	steps1 = append(steps1, step3)
	steps1 = append(steps1, step4)

	//hero.GetNearbyLocationsAsStrings(arr)
	steps = append(steps, step)
	steps = append(steps, step2)
	//steps = append(steps, step)
	//q := structure.CreateQuest(&structure.AllQuests, "first", "description 1", steps)
	q2 := structure.CreateQuest(&structure.AllQuests, "second", "description 2", steps1)
	//fmt.Println("LEEEEEN: ", len(structure.AllQuests), structure.AllQuests[0].Name)

	game.SetMap(mapg)
	game.SetHero(hero)
	game.SetQuests(structure.AllQuests)

	hlogs := new(structure.QuestLogs)
	hh := *hlogs
	//hlogs.Hlogs[0].Action = act.Action
	hh = q2.Do(game.GetHero(), q2.QuestId, hh)
	fmt.Println("hh: ", hh)
	hh, err := q2.Check(q2.QuestId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("hh: ", hh)
	mapg[3].AddQuestToNode(0)
	mapg[3].AddQuestToNode(0)

	mapg[2].AddQuestToNode(0)
	fmt.Println("CHECK: ", game.GetQuests()[0].Description)
	//fmt.Println("gg: ", arr[0].Quests[0].Name)
	//fmt.Println("gg: ", arr[2].Quests[0].Name)

	//fmt.Println("Quests of ", hero.GetLocationAsString(), " quests are: ", hero.GetAllQuestsInCurrentLocationAsStrings(arr))
	//	tt := hero.Stat(arr)
	//
	// fmt.Println(tt.GetLocation())
	// tmp, _ := hero.ReStat()
	// fmt.Println("Tmp: ", tmp.Location.Next[0])
	// fmt.Println(q.Step[0].Action.Action)
	structure.GameLoop(game)

}
