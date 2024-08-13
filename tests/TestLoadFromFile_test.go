package tests

import (
	"cligo/pkg"
	"os"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	list := pkg.ShoppingList{}
	list.AddItem("Milk")
	list.AddItem("Eggs")
	pkg.SaveToFile(&list)
	defer os.Remove("shopping_list.json")

	var loadedList pkg.ShoppingList
	pkg.LoadFromFile(&loadedList)

	if len(loadedList.Items) != 2 || loadedList.Items[0] != "Milk" || loadedList.Items[1] != "Eggs" {
		t.Errorf("Expected 'Milk', 'Eggs', got '%v'", loadedList.Items)
	}
}
