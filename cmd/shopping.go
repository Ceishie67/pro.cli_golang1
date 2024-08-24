package cmd

import (
	"cligo/pkg"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func Shopping() {
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add item")
		fmt.Println("2. Show items")
		fmt.Println("3. Mark item")
		fmt.Println("4. Modify quantity owner")
		fmt.Println("5. Modify required quantity")
		fmt.Println("6. Delete item")
		fmt.Println("7. Reorder item")
		fmt.Println("8. Exit and Save")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter item to add: ")
			var item string
			fmt.Scanln(&item)
			if err := pkg.AddItem(item); err != nil {
				fmt.Println("Error adding item:", err)
			} else {
				fmt.Println("Item added successfully!")
				err := pkg.ReassignIDs()
				if err != nil {
					return
				}
				pkg.ShowItems()
			}

		case 2:
			fmt.Println("Current Shopping List:")
			pkg.ShowItems()

		case 3:
			fmt.Print("Enter the item ID to mark/unmark: ")
			var id int
			fmt.Scanln(&id)
			if err := pkg.MarkItem(id); err != nil {
				fmt.Println("Error marking item:", err)
			} else {
				fmt.Println("Item marked successfully!")
				pkg.ShowItems()
			}
		case 4:
			fmt.Print("Enter the item ID to change owned quantity: ")
			var id int
			fmt.Scanln(&id)
			fmt.Print("Enter owned quantity: ")
			var val int
			fmt.Scanln(&val)
			if err := pkg.ModifyOwned(id, val); err != nil {
				fmt.Println("Error invalid value:", err)
			} else {
				fmt.Println("value added successfully!")
				pkg.ShowItems()
			}
		case 5:
			fmt.Print("Enter the item ID to change required quantity: ")
			var id int
			fmt.Scanln(&id)
			fmt.Print("Enter required quantity: ")
			var val int
			fmt.Scanln(&val)
			if err := pkg.ModifyRequired(id, val); err != nil {
				fmt.Println("Error invalid value:", err)
			} else {
				fmt.Println("value added successfully!")
				pkg.ShowItems()
			}

		case 6:
			fmt.Print("Enter the item ID to delete: ")
			var id int
			fmt.Scanln(&id)
			if err := pkg.RemoveItem(id); err != nil {
				fmt.Println("Error deleting item:", err)
			} else {
				fmt.Println("Item deleted successfully!")
				err := pkg.ReassignIDs()
				if err != nil {
					return
				}
				pkg.ShowItems()
			}

		case 7:
			fmt.Print("Element to replace (ID): ")
			var fromIndex int
			fmt.Scanln(&fromIndex)
			fmt.Print("New place (ID): ")
			var toIndex int
			fmt.Scanln(&toIndex)
			if err := pkg.ReorderList(fromIndex, toIndex); err != nil {
				fmt.Println("Error reordering items:", err)
			} else {
				fmt.Println("List reordered successfully!")
				err := pkg.ReassignIDs()
				if err != nil {
					return
				}
				pkg.ShowItems()
			}

		case 8:
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Invalid choice, Please try again.")
		}
	}
}
