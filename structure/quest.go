package structure

import (
	"fmt"
)

/*
0 - lookup
1 - WalkTo
*/
var listOfActions []int = []int{0, 1, 2, 3}

type action struct {
	HeroName string
	Action   int
}

type QuestStep struct {
	QuestId int
	Hero    *Hero
	Action  *action
}

type quest struct {
	QuestId     int
	Name        string
	Description string
	Step        []QuestStep
}

var allQuests []quest = []quest{}

type questFunctions interface {
	Do(questId *quest) bool
	Check(questId *quest) bool
}

func CreateAction(name string, todo int) action {
	act := &action{HeroName: name, Action: todo}
	return *act
}

func CreateStep(h *Hero, a *action) QuestStep {
	s := &QuestStep{Hero: h, Action: a}
	return *s
}

func CreateQuest(name string, description string, steps []QuestStep) quest {
	q := &quest{QuestId: len(allQuests), Name: name, Description: description, Step: steps}
	allQuests = append(allQuests, *q)
	return *q
}

func (quest *quest) Do(questId int) bool {
	for i, _ := range allQuests {
		if i == questId {
			fmt.Println("i: ", i)
		}
	}
	return false
}

func (quest *quest) Check(questId int) bool {
	return false
}
