package structure

type item struct {
	Name string
}

type Inventory struct {
	Hero *Hero
	Items []item
}

type InventoryFunctions interface {
	PutItem(name string) Inventory
}

func CreateInventory(h *Hero) Inventory {
	inventory := &Inventory{Hero: h, Items: nil}
	return *inventory
}