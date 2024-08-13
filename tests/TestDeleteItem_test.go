package tests

import (
	"cligo/pkg"
	"testing"
)

func TestDeleteItem(t *testing.T) {
	list := pkg.ShoppingList{}
	list.AddItem("Milk")
	list.AddItem("Eggs")
	list.RemoveItem(0)

	if len(list.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(list.Items))
	}

	if list.Items[0] != "Eggs" {
		t.Errorf("Expected 'Eggs', got '%s'", list.Items[0])
	}
}
