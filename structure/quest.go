package structure

import (
	"fmt"
	"os"
)

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
	Do(hero *Hero, questId int, questlogs QuestLogs) QuestLogs
	Check(questId int) (QuestLogs, error)
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

func (quest *Quest) Do(hero *Hero, questId int, questlogs QuestLogs) QuestLogs {
	for i := range allQuests {
		if i == questId {
			quest.QuestTakenBy = hero
			questlogs.Hlogs = make([]QuestLog, 0)

			for _, st := range quest.Step {
				ql := new(QuestLog)
				ql.Action = st.Action.Action
				ql.Location = hero.GetLocationAsString()
				questlogs.Hlogs = append(questlogs.Hlogs, *ql)

			}
			//questlog.Hlogs = append(questlog.Hlogs,quest)
			questlogs.HWrite(questlogs, quest.Name)
			//herolog.HLogQ(hero, 2, "Do", quest.QuestId)
			return questlogs
			//fmt.Println("i: ", i)

		}
	}
	return questlogs
}

func (quest *Quest) Check(questId int) (QuestLogs, error) {
	count := 0
	data := new(QuestLogs)
	data2 := *data

	for _, e := range allQuests {
		if e.QuestId != questId {
			continue
		}

		if _, err := os.Stat(e.Name); os.IsExist(err) {
			data2 = data2.HRead(e.Name)
		} else if os.IsNotExist(err) {
			return *new(QuestLogs), fmt.Errorf("file %q does not exist", e.Name)
		} else {
			return *new(QuestLogs), fmt.Errorf("error checking file %q: %v", e.Name, err)
		}

		for _, elem2 := range e.Step {
			for _, elem := range data2.Hlogs {
				if elem.Action == elem2.Action.Action {
					count++
				}
			}
		}

		if count == len(e.Step) {
			return data2, nil
		} else {
			fmt.Println("Count: ", count)
		}
	}

	return data2, fmt.Errorf("quest %d not found", questId)
}
