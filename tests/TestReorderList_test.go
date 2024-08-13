package tests

import (
	"cligo/pkg"
	"testing"
)

func TestReorderList(t *testing.T) {
	list := pkg.ShoppingList{}
	list.AddItem("Milk")
	list.AddItem("Eggs")
	list.AddItem("Bread")

	list.ReorderList(2, 0)

	if list.Items[0] != "Bread" || list.Items[1] != "Milk" {
		t.Errorf("Expected order 'Bread', 'Milk', 'Eggs', got '%s', '%s'", list.Items[0], list.Items[1])
	}
}
