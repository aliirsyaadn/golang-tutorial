package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Product struct represents a product in the inventory
type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Category string  `json:"category"`
}

// Inventory represents the collection of products
type Inventory struct {
	Products map[int]Product `json:"products"`
}

func main() {
	inv := &Inventory{Products: make(map[int]Product)}
	loadInventory("inventory.json", inv)

	// Example usage:
	inv.addProduct(1, "Laptop", 999.99, 10, "Electronics")
	inv.addProduct(2, "Shirt", 29.99, 50, "Clothing")
	inv.addProduct(3, "Book", 14.99, 100, "Books")

	inv.listProducts()

	inv.updateProduct(2, "T-Shirt", 19.99, 40, "Clothing")

	inv.viewProductDetails(1)

	inv.removeProduct(3)

	inv.listProducts()

	saveInventory("inventory.json", inv)
}

func (inv *Inventory) addProduct(id int, name string, price float64, quantity int, category string) {
	prod := Product{ID: id, Name: name, Price: price, Quantity: quantity, Category: category}
	inv.Products[id] = prod
	fmt.Println("Product added:", prod)
}

func (inv *Inventory) updateProduct(id int, name string, price float64, quantity int, category string) {
	if prod, ok := inv.Products[id]; ok {
		prod.Name = name
		prod.Price = price
		prod.Quantity = quantity
		prod.Category = category
		inv.Products[id] = prod
		fmt.Println("Product updated:", prod)
	} else {
		fmt.Println("Product not found")
	}
}

func (inv *Inventory) removeProduct(id int) {
	if _, ok := inv.Products[id]; ok {
		delete(inv.Products, id)
		fmt.Println("Product removed:", id)
	} else {
		fmt.Println("Product not found")
	}
}

func (inv *Inventory) viewProductDetails(id int) {
	if prod, ok := inv.Products[id]; ok {
		fmt.Println("Product details:")
		fmt.Printf("ID: %d\nName: %s\nPrice: $%.2f\nQuantity: %d\nCategory: %s\n", prod.ID, prod.Name, prod.Price, prod.Quantity, prod.Category)
	} else {
		fmt.Println("Product not found")
	}
}

func (inv *Inventory) listProducts() {
	fmt.Println("List of Products:")
	for _, prod := range inv.Products {
		fmt.Printf("ID: %d, Name: %s, Price: $%.2f, Quantity: %d, Category: %s\n", prod.ID, prod.Name, prod.Price, prod.Quantity, prod.Category)
	}
}

func loadInventory(filename string, inv *Inventory) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("No existing inventory found.")
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading inventory file:", err)
		return
	}

	err = json.Unmarshal(data, inv)
	if err != nil {
		fmt.Println("Error unmarshalling inventory data:", err)
		return
	}
	fmt.Println("Inventory loaded successfully.")
}

func saveInventory(filename string, inv *Inventory) {
	data, err := json.MarshalIndent(inv, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling inventory data:", err)
		return
	}

	err = ioutil.WriteFile(filename, data, 0o644)
	if err != nil {
		fmt.Println("Error writing inventory file:", err)
		return
	}
	fmt.Println("Inventory saved successfully.")
}
