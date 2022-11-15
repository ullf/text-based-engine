package structure

/*
0 - lookup
1 - WalkTo
*/
var listOfActions []int = []int{0,1,2,3}

type action struct {
	HeroName string
	Action int
}

type questStep struct {
	QuestId int
	Hero *hero
	Action *action

}

type quest struct {
	QuestId int
	Name string
	Description string
	Step []questStep
}

type questFunctions interface {
	Do(questId int) bool
	Check(questId int) bool
}

func CreateAction(name string, todo int) action {
	act := &action{HeroName: name, Action: todo}
	return *act
}

func CreateSteps(h *hero,a *action) questStep {
	s := &questStep{Hero: h,Action: a}

	return *s
}

func CreateQuest(name string, description string,steps []questStep) quest {
	q := &quest{ Name: name, Description: description, Step: steps }
	
	return *q
}

