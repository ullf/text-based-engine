package structure

type item struct {
	Name string
}

type Inventory struct {
	Hero *Hero
	Items []item
}

type InventoryFunctions interface {
	PutItem(it item)
}

func CreateItem(name string) item {
	it := &item{Name: name}
	return *it
}

func CreateInventory(h *Hero) Inventory {
	inventory := &Inventory{Hero: h, Items: nil}
	return *inventory
}

func (i *Inventory) PutItem(it item) {
	i.Items = append(i.Items,it)
}