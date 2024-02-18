# Inventory Management System

## Problem Description:
You are tasked with developing an inventory management system for a small retail store. The system should allow the store to keep track of its products, including their names, prices, quantities, and categories. You are required to implement this system in Go, incorporating various composite types and concepts such as arrays, slices, maps, structs, and string manipulation.

### Functionalities to Implement:
1. **Add Product:**
   - Implement a function to add a new product to the inventory.
   - Each product should have a unique identifier, name, price, quantity, and category.

2. **Update Product:**
   - Create a function to update the details of an existing product.
   - Allow updating product name, price, quantity, and category based on the product's identifier.

3. **Remove Product:**
   - Implement a function to remove a product from the inventory based on its identifier.

4. **View Product Details:**
   - Develop a function to display the details of a specific product based on its identifier.

5. **List All Products:**
   - Create a function to list all products currently available in the inventory, including their details.

6. **Search Products by Category:**
   - Implement a function to search for products by category and display their details.

7. **Restock Product:**
   - Create a function to restock a specific product by adding a specified quantity to its existing stock.

8. **Sell Product:**
   - Implement a function to sell a specific quantity of a product from the inventory.
   - Ensure that the product quantity is updated accordingly, and if the requested quantity exceeds the available stock, display an appropriate message.

9. **Calculate Revenue:**
   - Develop a function to calculate the total revenue generated from product sales.
   - Ensure that the revenue calculation considers the quantity sold and the price of each product.

10. **Data Persistence:**
    - Implement file handling mechanisms to store and load the inventory data from a file. Use JSON or any suitable format for serialization.

11. **User Interface (Optional):**
    - Design a basic user interface (command-line interface) to interact with the inventory management system, enabling users to perform the above operations conveniently.
