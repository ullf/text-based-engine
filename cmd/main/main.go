package main

import (
	"fmt"

	"github.com/ullf/text-based-engine/structure"
)

func main() {
	arr := structure.CreateArrOfNodes()
	arr = structure.AppendElemToArr(arr, "paris")
	arr = structure.AppendElemToArr(arr, "madrid")
	arr = structure.AppendElemToArr(arr, "berlin")
	arr = structure.AppendElemToArr(arr, "helsinki")

	berlin := structure.FindElementByName(arr, "berlin")
	paris := structure.FindElementByName(arr, "paris")

	hlog := structure.NewHeroLog(-1, "", "")
	hls := structure.NewHeroLogs(nil)
	hh := *hls
	hh = hlog.HRead("data")

	//elem1 := structure.FindElementById(arr, 3)
	if berlin.Id >= 0 {
		arr[1].AddChildToNodeByName(arr, "madrid", &berlin, 0)
		//arr[3].AddChildToNodeByName(arr, "helsinki", &berlin, 0)
		arr[3].AddChildToNodeByName(arr, "helsinki", &paris, 0)
		//fmt.Println("Check : ", berlin.Id)
	} else {
		panic("error")
	}

	//structure.AddChildToNodeByName(arr, "20", &elem1)
	hero := structure.CreateHero("mark")
	hero.SetLocation(arr, "helsinki")
	//hero.GetNearbyLocations(arr)
	fmt.Println("\nlocation: ", hero.GetLocationAsString())
	bb := hero.WalkTo(arr, "paris", hlog)
	hh.AppendHeroLog(hlog)
	hlog.HWrite(hh, "data")
	fmt.Println("bool: ", bb)
	fmt.Println("\nlocation: ", hero.GetLocationAsString())
	//fmt.Println(arr)
	structure.PrintAll(arr[1].Next, arr[1].Name)
	structure.PrintAll(arr[0].Next, arr[0].Name)

	act := structure.CreateAction("lookup", 0)
	step := structure.CreateStep(&hero, &act)
	steps := make([]structure.QuestStep, 0)
	hero.GetNearbyLocationsAsStrings(arr, hlog)
	hh.AppendHeroLog(hlog)
	hlog.HWrite(hh, "data")
	steps = append(steps, step)
	q := structure.CreateQuest("first", "description 1", steps)
	q.Do(&hero, 0, hlog)
	q.Check(0, hh)
	hh.AppendHeroLog(hlog)
	hlog.HWrite(hh, "data")
	//fmt.Println(q.Step[0].Action.Action)
}
