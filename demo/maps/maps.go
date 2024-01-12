package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 1
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("Need", shoppingList["eggs"], "eggs")

	cereals, found := shoppingList["cereals"]
	fmt.Println("Need cereals?")
	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("Yup", cereals, "boxes")
	}

	totalItems := 0
	for _, item := range shoppingList {
		totalItems += item
	}
	fmt.Println("There are are", totalItems, "items in the shopping list")
}
