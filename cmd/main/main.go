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
	bb := hero.WalkTo(arr, "paris")
	fmt.Println("bool: ", bb)
	fmt.Println("\nlocation: ", hero.GetLocationAsString())
	//fmt.Println(arr)
	structure.PrintAll(arr[1].Next, arr[1].Name)
	structure.PrintAll(arr[0].Next, arr[0].Name)

	act := structure.CreateAction("lookup", 0)
	act2 := structure.CreateAction("walkto", 1)
	step := structure.CreateStep(&hero, &act)
	step2 := structure.CreateStep(&hero, &act2)
	steps := make([]structure.QuestStep, 0)

	act3 := structure.CreateAction("lookup", 0)
	act4 := structure.CreateAction("lookup", 0)
	step3 := structure.CreateStep(&hero, &act3)
	step4 := structure.CreateStep(&hero, &act4)
	steps1 := make([]structure.QuestStep, 0)

	steps1 = append(steps1, step3)
	steps1 = append(steps1, step4)

	hero.GetNearbyLocationsAsStrings(arr)
	steps = append(steps, step)
	steps = append(steps, step2)
	//steps = append(steps, step)
	q := structure.CreateQuest("first", "description 1", steps)
	q2 := structure.CreateQuest("second", "description 2", steps1)
	hlogs := new(structure.QuestLogs)
	hh := *hlogs
	//hlogs.Hlogs[0].Action = act.Action
	hh = q.Do(&hero, 0, hh)
	qq, err := q2.Check(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(qq)
	hero.Stat(arr)
	tmp, _ := hero.ReStat()
	fmt.Println(tmp)
	//fmt.Println(q.Step[0].Action.Action)
}
