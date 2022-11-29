package main

import (
	//"log"
	"fmt"

	"github.com/ullf/text-based-engine/structure"
)

func main() {
	arr := structure.CreateArrOfNodes()
	arr = structure.AppendElemToArr(arr, "10")
	arr = structure.AppendElemToArr(arr, "20")
	arr = structure.AppendElemToArr(arr, "30")
	arr = structure.AppendElemToArr(arr, "40")

	structure.FindElementById(arr, 0)
	structure.FindElementById(arr, 1)
	elem1 := structure.FindElementById(arr, 3)
	structure.AddChildToNodeByName(arr, "20", &elem1)
	structure.AddChildToNodeByName(arr, "20", &elem1)
	hero := structure.CreateHero("mark")
	hero.SetLocation(arr, "30")
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
}
