package tests

import (
	"cligo/pkg"
	"testing"
)

func TestMarkItem(t *testing.T) {
	list := pkg.ShoppingList{}
	list.AddItem("Milk")
	list.MarkItemAndUnmarkItem(0)

	if !list.Marked[0] {
		t.Errorf("Expected item 0 to be marked")
	}
}
