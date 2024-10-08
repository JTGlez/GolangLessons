package main

type Item struct {
	Name  string
	Price float64
}

type StorageItemsSlice struct {
	Items  map[int]Item
	LastId int
}

func (s *StorageItemsSlice) GetItem(id int) Item {
	return s.Items[id]
}

func (s *StorageItemsSlice) AddItem(item *Item) {
	if item.Name == "" {
		return
	}
	s.LastId++
	s.Items[s.LastId] = *item
}
