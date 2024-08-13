package tests

import (
	"bytes"
	"cligo/pkg"
	"io"
	"os"
	"testing"
)

func TestShowItems(t *testing.T) {
	list := pkg.ShoppingList{}
	list.AddItem("Milk")
	list.AddItem("Eggs")

	expected := "1. Milk\n2. Eggs\n"

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	list.ShowItems()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	if buf.String() != expected {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expected, buf.String())
	}
}
