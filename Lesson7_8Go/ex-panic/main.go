package main

func main() {

	store := &StorageItemsSlice{
		Items: []Item{
			{
				Name:  "Item 1",
				Price: 9.99,
			},
			{
				Name:  "Item 2",
				Price: 19.99,
			},
		},
	}

	/* 	item := store.GetItem(3)
	   	fmt.Printf("%v\n", item) */

	store.AddItem(nil)

}
