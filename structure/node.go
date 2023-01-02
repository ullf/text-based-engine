package structure

// Structure of an abstract location
// Node represents some node in a graph
type Node struct {
	Id     int
	Name   string
	Next   []*Node
	Quests []*Quest
}

type NodeOps interface {
	AddChildToNodeByName(array []Node, name string, new_node *Node, z int)
	AddQuestToNode(questId int)
}

func CreateArrOfNodes() []Node {
	var array = make([]Node, 0)
	return array
}

func AppendElemToArr(array []Node, name string) []Node {
	elem := new(Node)
	elem.Next = make([]*Node, 0)
	elem = &Node{Id: len(array), Name: name, Next: elem.Next}
	array = append(array, *elem)
	return array
}

func Lookup(array []Node, element_to_find Node) Node {
	for _, elem := range array {
		if elem.Id == element_to_find.Id {
			return element_to_find
		}
	}
	return element_to_find
}

func FindElementById(array []Node, id int) Node {
	if id >= 0 && id < len(array) {
		return array[id]
	}
	elem := new(Node)
	return *elem
}

func FindElementByName(array []Node, name string) Node {
	elem := 0
	for i, e := range array {
		if e.Name == name {
			elem = i
			return array[elem]
		}
	}
	//fmt.Println("elem: ", array[elem])
	return array[elem]
}

func AddChildToNodeById(array []Node, id int, new_node *Node) {
	for _, e := range array {
		if e.Id == id {
			array[id].Next = append(e.Next, new_node)
			array[new_node.Id].Next = append(new_node.Next, &array[e.Id])
			break
		}
	}
}

func (node *Node) AddChildToNodeByName(array []Node, name string, new_node *Node, z int) {
	if len(node.Next) == 0 {
		node.Next = make([]*Node, 0)
	} //else {
	node.Next = append(node.Next, new_node)
	z++
	if z < 2 {
		array[new_node.Id].AddChildToNodeByName(array, new_node.Name, node, z)
	}
}

func (node *Node) AddQuestToNode(questId int) {
	for _, e := range AllQuests {
		if questId == e.QuestId {
			node.Quests = append(node.Quests, &e)
			break
		}
	}
}
