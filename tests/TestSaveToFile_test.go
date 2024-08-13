package tests

import (
	"cligo/pkg"
	"encoding/json"
	"os"
	"testing"
)

func TestSaveToFile(t *testing.T) {
	list := pkg.ShoppingList{}
	list.AddItem("Milk")
	list.AddItem("Eggs")

	pkg.SaveToFile(&list)

	defer os.Remove("shopping_list.json")

	file, err := os.Open("shopping_list.json")
	if err != nil {
		t.Fatalf("Expected no error opening file, got %v", err)
	}
	defer file.Close()

	var loadedList pkg.ShoppingList
	json.NewDecoder(file).Decode(&loadedList)

	if len(loadedList.Items) != 2 || loadedList.Items[0] != "Milk" || loadedList.Items[1] != "Eggs" {
		t.Errorf("Expected 'Milk', 'Eggs', got '%v'", loadedList.Items)
	}
}
