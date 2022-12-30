package structure

import (
	"fmt"
	"os"
)

/*
0 - lookup (GetNearbyLocationsAsStrings)
1 - WalkTo (WalkTo)
2 - GetAllQuestsinCurrentLocation
*/

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

var AllQuests []Quest = []Quest{}

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

func CreateQuest(qs *[]Quest, name string, description string, steps []QuestStep) Quest {
	q := &Quest{QuestId: len(*qs), Name: name, Description: description, Step: steps}
	*qs = append(*qs, *q)
	fmt.Println(q.Name, " Len of allQuests: ", len(*qs))
	return *q
}

func (quest *Quest) Do(hero *Hero, questId int, questlogs QuestLogs) QuestLogs {
	for i := range AllQuests {
		if i == questId {
			//fmt.Println("HERO: ", hero)
			quest.QuestTakenBy = hero
			hero.TakenQuest = quest.QuestId
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

	for _, e := range AllQuests {
		if e.QuestId != questId {
			continue
		}
		name := e.Name + ".json"
		if _, err := os.Stat(name); os.IsExist(err) {
			fmt.Println("exists")
			data2 = data2.HRead(name)
		} else if os.IsNotExist(err) {
			return *new(QuestLogs), fmt.Errorf("file %q does not exist", name)
		} else {
			data2 = data2.HRead(name)
			for _, elem2 := range e.Step {
				for _, elem := range data2.Hlogs {
					if elem.Action == elem2.Action.Action {
						count++
					}
				}
			}

			if count == len(e.Step) {
				fmt.Println("Count: ", count)
				err := os.Remove(name)
				if err != nil {
					panic(err)
				}
				return data2, nil
			} else {
				fmt.Println("Count: ", count)
			}
			return data2, nil // fmt.Errorf("error checking file %q: %v", name, err)
		}
	}

	return data2, fmt.Errorf("quest %d not found", questId)
}

func FindQuestByName(questName string) *Quest {
	found := new(Quest)

	for _, e := range AllQuests {
		if e.Name == questName {
			found = &e
		}
	}
	return found
}
