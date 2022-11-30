package main

import (
	"fmt"

	"github.com/ullf/text-based-engine/structure"
)

func main() {
	arr := structure.CreateArrOfNodes()
	arr = structure.AppendElemToArr(arr, "moscow")
	arr = structure.AppendElemToArr(arr, "petersburg")
	arr = structure.AppendElemToArr(arr, "novgorod")
	arr = structure.AppendElemToArr(arr, "tula")

	peters := structure.FindElementByName(arr, "petersburg")
	//elem1 := structure.FindElementById(arr, 3)
	structure.AddChildToNodeByName(arr, "moscow", &peters)
	//structure.AddChildToNodeByName(arr, "20", &elem1)
	hero := structure.CreateHero("mark")
	hero.SetLocation(arr, "moscow")
	fmt.Println(arr, "\nlocation: ", hero.GetLocation())
	hero.WalkTo(arr,"petersburg")
	fmt.Println(arr, "\nlocation: ", hero.GetLocation())
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
