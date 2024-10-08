package main

func main() {
	println("aaa im trying to add an item to a nil map!")
	store := &StorageItemsSlice{
		// Items: make(map[int]Item, 0),
	}

	store.AddItem(&Item{Name: "Item 1", Price: 9.99})
	store.GetItem(0)
}
