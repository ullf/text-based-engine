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
	//elem1 := structure.FindElementById(arr, 3)
	structure.AddChildToNodeByName(arr, "madrid", &berlin)
	structure.AddChildToNodeByName(arr, "helsinki", &berlin)
	//structure.AddChildToNodeByName(arr, "20", &elem1)
	hero := structure.CreateHero("mark")
	hero.SetLocation(arr, "helsinki")
	fmt.Println(arr, "\nlocation: ", hero.GetLocation())
	hero.WalkTo(arr,"madrid")
	fmt.Println(arr, "\nlocation: ", hero.GetLocation())
	hero.GetNearbyLocations(arr)
	act := structure.CreateAction("lookup", 0)
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
	structure.GetListOfAllItems("inventory_objects/all_objects")
}
