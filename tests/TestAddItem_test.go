package tests

import (
	"cligo/pkg"
	"testing"
)

func TestAddItem(t *testing.T) {
	list := pkg.ShoppingList{}
	list.AddItem("Milk")

	if len(list.Items) != 1 {
		t.Errorf("Expected 1 item, got %d", len(list.Items))
	}

	if list.Items[0] != "Milk" {
		t.Errorf("Expected 'Milk', got '%s'", list.Items[0])
	}
}
