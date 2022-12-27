package structure

import "fmt"

/*
0 - lookup (GetNearbyLocationsAsStrings)
1 - WalkTo (WalkTo)
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

type Quest struct {
	QuestId      int
	Name         string
	QuestTakenBy *Hero
	Description  string
	Step         []QuestStep
}

var allQuests []Quest = []Quest{}

type QuestFunctions interface {
	Do(hero *Hero, questId *Quest) bool
	Check(questId *Quest) bool
}

func CreateAction(name string, todo int) action {
	act := &action{HeroName: name, Action: todo}
	return *act
}

func CreateStep(h *Hero, a *action) QuestStep {
	s := &QuestStep{Hero: h, Action: a}
	return *s
}

func CreateQuest(name string, description string, steps []QuestStep) Quest {
	q := &Quest{QuestId: len(allQuests), Name: name, Description: description, Step: steps}
	allQuests = append(allQuests, *q)
	return *q
}

func (quest *Quest) Do(hero *Hero, questId int, herolog *HeroLog) bool {
	for i := range allQuests {
		if i == questId {
			quest.QuestTakenBy = hero
			herolog.HLogQ(hero, 2, "Do", quest.QuestId)
			return true
			//fmt.Println("i: ", i)

		}
	}
	return false
}

func (quest *Quest) Check(questId int, data HeroLogs) bool {
	count := 0
	for _, e := range allQuests {
		if e.QuestId == questId {
			for _, elem2 := range e.Step {
				for _, elem := range data.Hlogs {
					if elem.Action == elem2.Action.Action {
						count++

					}
				}
			}
		}
		if count == len(e.Step) {
			fmt.Println("Done!")
			return true
		} else {
			fmt.Println("Count: ", count)
		}
	}
	return false
}
