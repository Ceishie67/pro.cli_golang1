package main

import (
	"cligo/cmd"
	"cligo/pkg"
)

func main() {
	shoppingList := pkg.ShoppingList{}

	pkg.LoadFromFile(&shoppingList)
	cmd.Shopping(shoppingList)
}
