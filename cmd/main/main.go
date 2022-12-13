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
		arr[3].AddChildToNodeByName(arr, "helsinki", &berlin, 0)
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
	hero.WalkTo(arr, "madrid")
	fmt.Println("\nlocation: ", hero.GetLocationAsString())
	//fmt.Println(arr)
	structure.PrintAll(arr[1].Next, arr[1].Name)
	structure.PrintAll(arr[3].Next, arr[3].Name)
	//fmt.Print("Check: ")
	//hero.GetNearbyLocations(arr)
	/*act := structure.CreateAction("lookup", 0)
	step := structure.CreateStep(&hero, &act)
	steps := make([]structure.QuestStep, 0)
	steps = append(steps, step)
	q := structure.CreateQuest("first", "description 1", steps)
	fmt.Println(q.Step[0].Action.Action)
	q.Do(0)
	inv := structure.CreateInventory(&hero)
	ball := structure.CreateItem("ball")
	inv.PutItem(ball)
	inv.PutItem(ball)
	fmt.Println(" ", inv.Items)
	fmt.Println("-----")
	structure.GetListOfAllItems("inventory_objects/all_objects")*/
}
