package main

type Item struct {
	Name  string
	Price float64
}

type StorageItemsSlice struct {
	Items []Item
}

func (s *StorageItemsSlice) GetItem(id int) Item {
	return s.Items[id]
}

func (s *StorageItemsSlice) AddItem(item *Item) {
	// implicit dereference
	if item.Name == "" {
		return
	}

	s.Items = append(s.Items, *item)
}
