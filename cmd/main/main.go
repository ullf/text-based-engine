package main

import (
	//"log"
	"fmt"
	"text-based-engine/structure"
)

func main() {
	fmt.Println("Hello")
	arr := structure.CreateArrOfNodes()
	arr = structure.AppendElemToArr(arr, "mark")
	arr = structure.AppendElemToArr(arr, "alex")
	arr = structure.AppendElemToArr(arr, "arslan")
	arr = structure.AppendElemToArr(arr, "jay")

	structure.FindElementById(arr, 0)
	structure.FindElementById(arr, 1)
	elem2 := structure.FindElementById(arr, 2)
	jay := structure.FindElementById(arr, 3)
	structure.AddChildToNodeById(arr, 0, &elem2)
	structure.AddChildToNodeByName(arr, "mark", &jay)
	fmt.Println("found by id: ", structure.FindElementById(arr, 0))
	hero := structure.CreateHero("me")
	hero.SetLocation(arr, "jay")
	fmt.Println("Hero's Name: ", hero.GetName(), "location: ", hero.GetLocation())
	hero.GetNearbyLocations(arr)
	hero.Move(arr, "mark")
	hero.GetNearbyLocations(arr)
	hero.WalkTo(arr, "arslan")
	fmt.Println("Current loc: ", hero.GetLocation())
	hero.GetNearbyLocations(arr)
	//hero.GoForward(arr, "mark"
	//q := structure.CreateQuest("a","description",nil)
	//ac := structure.CreateAction("mark",0)
	//q.CreateSteps(hero,ac)
	//fmt.Println(q.Name)
	//fmt.Println(structure.FindElementById(arr, 0).Next)
	//log.Println(http.ListenAndServe(":8080",nil))

}
