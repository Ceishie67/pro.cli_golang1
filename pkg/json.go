package pkg

//
//import (
//	"encoding/json"
//	"fmt"
//	"os"
//)
//
//const FILENAME = "./storage/shopping_list.json"
//
//func LoadFromFile(shoppingList *ShoppingList) {
//	file, err := os.Open(FILENAME)
//	if err != nil {
//		if os.IsNotExist(err) {
//			fmt.Println("No existing shopping list found, starting a new one.")
//			return
//		}
//		fmt.Printf("Error opening file: %v\n", err)
//		return
//	}
//	defer file.Close()
//
//	err = json.NewDecoder(file).Decode(shoppingList)
//	if err != nil {
//		fmt.Printf("Error decoding JSON: %v\n", err)
//	} else {
//		fmt.Println("Shopping list loaded from file.")
//	}
//}
//
//func SaveToFile(shoppingList *ShoppingList) {
//	file, err := os.Create(FILENAME)
//	if err != nil {
//		fmt.Printf("Error creating file: %v\n", err)
//		return
//	}
//	defer file.Close()
//
//	err = json.NewEncoder(file).Encode(shoppingList)
//	if err != nil {
//		fmt.Printf("Error encoding JSON: %v\n", err)
//	} else {
//		fmt.Println("Shopping list saved to file.")
//	}
//}
