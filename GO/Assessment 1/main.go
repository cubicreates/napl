package main

import (
	"fmt"
	"strings"
	"time"
)

func printDivider() {
	fmt.Println(strings.Repeat("-", 75))
}

func printHeader(title string) {
	fmt.Println()
	printDivider()
	fmt.Printf("  %s\n", title)
	printDivider()
}

func displayInventory(items []Item) {
	if len(items) == 0 {
		fmt.Println("  [Inventory is empty]")
		return
	}
	fmt.Printf("  %-6s | %-35s | %-8s | %-12s | %-12s\n", "ID", "Item Name", "Quantity", "Price ($)", "Value ($)")
	fmt.Println("  " + strings.Repeat("-", 73))
	for _, it := range items {
		val := float64(it.Quantity) * it.Price
		fmt.Printf("  %-6d | %-35s | %-8d | %12.2f | %12.2f\n", it.ID, it.Name, it.Quantity, it.Price, val)
	}
}

func main() {
	fmt.Println("===========================================================================")
	fmt.Println("  GO PRACTICAL ASSIGNMENT 1: INVENTORY MANAGEMENT MODULE                   ")
	fmt.Println("  Integrating Sessions 1-4: Structs, Slices, Validation, Errors & Channels ")
	fmt.Println("===========================================================================")

	inv := NewInventory()

	// -------------------------------------------------------------------------
	// 1. ADD OPERATIONS (Success Cases)
	// -------------------------------------------------------------------------
	printHeader("STEP 1: Adding Items to Inventory (Slices & Structs)")

	initialProducts := []Item{
		{ID: 101, Name: "Ultrabook Pro 15\"", Quantity: 12, Price: 1199.99},
		{ID: 102, Name: "Noise-Cancelling Headphones", Quantity: 25, Price: 199.50},
		{ID: 103, Name: "Mechanical Keyboard", Quantity: 40, Price: 89.95},
		{ID: 104, Name: "Ergonomic Optical Mouse", Quantity: 60, Price: 45.00},
		{ID: 105, Name: "4K UHD IPS Monitor 27\"", Quantity: 15, Price: 349.99},
	}

	for _, product := range initialProducts {
		if err := inv.Add(product); err != nil {
			fmt.Printf("  [FAILED] Could not add ID %d: %v\n", product.ID, err)
		} else {
			fmt.Printf("  [ADDED]  ID: %-4d | Name: %-30s | Qty: %-3d | Price: $%.2f\n",
				product.ID, product.Name, product.Quantity, product.Price)
		}
	}

	fmt.Println()
	fmt.Println("  Current Inventory State:")
	displayInventory(inv.GetAll())

	// -------------------------------------------------------------------------
	// 2. INPUT VALIDATION & ERROR HANDLING
	// -------------------------------------------------------------------------
	printHeader("STEP 2: Input Validation & Error Handling (Boundary Checks)")

	testCases := []struct {
		description string
		item        Item
	}{
		{
			description: "Duplicate ID Check (ID: 101 already exists)",
			item:        Item{ID: 101, Name: "Another Laptop", Quantity: 5, Price: 999.00},
		},
		{
			description: "Negative Quantity Check (Qty: -10)",
			item:        Item{ID: 106, Name: "USB-C Hub", Quantity: -10, Price: 29.99},
		},
		{
			description: "Negative Price Check (Price: -$15.00)",
			item:        Item{ID: 107, Name: "HDMI Cable", Quantity: 20, Price: -15.00},
		},
		{
			description: "Empty Item Name Check",
			item:        Item{ID: 108, Name: "   ", Quantity: 10, Price: 19.99},
		},
		{
			description: "Invalid ID Check (ID <= 0)",
			item:        Item{ID: 0, Name: "Webcam HD", Quantity: 10, Price: 59.99},
		},
	}

	for _, tc := range testCases {
		err := inv.Add(tc.item)
		if err != nil {
			fmt.Printf("  [EXPECTED ERROR] %-42s -> Error: %v\n", tc.description, err)
		} else {
			fmt.Printf("  [UNEXPECTED SUCCESS] %s added unexpectedly!\n", tc.description)
		}
	}

	// -------------------------------------------------------------------------
	// 3. UPDATE OPERATION
	// -------------------------------------------------------------------------
	printHeader("STEP 3: Update Operation with Field Validation")

	fmt.Println("  Updating ID 103 (Mechanical Keyboard) with new pricing and stock...")
	err := inv.Update(103, "RGB Mechanical Gaming Keyboard", 35, 94.99)
	if err != nil {
		fmt.Printf("  [UPDATE ERROR] %v\n", err)
	} else {
		updatedItem, _ := inv.Get(103)
		fmt.Printf("  [UPDATED] ID: %d | New Name: %s | New Qty: %d | New Price: $%.2f\n",
			updatedItem.ID, updatedItem.Name, updatedItem.Quantity, updatedItem.Price)
	}

	fmt.Println("  Attempting to update non-existent item (ID: 999)...")
	err = inv.Update(999, "Ghost Item", 5, 10.00)
	if err != nil {
		fmt.Printf("  [EXPECTED ERROR] Update ID 999 -> Error: %v\n", err)
	}

	// -------------------------------------------------------------------------
	// 4. REMOVE OPERATION
	// -------------------------------------------------------------------------
	printHeader("STEP 4: Remove Operation (Slice Reslicing)")

	fmt.Println("  Removing ID 104 (Ergonomic Optical Mouse) from inventory...")
	err = inv.Remove(104)
	if err != nil {
		fmt.Printf("  [REMOVE ERROR] %v\n", err)
	} else {
		fmt.Println("  [REMOVED] ID 104 successfully removed from inventory slice.")
	}

	fmt.Println("  Attempting to remove non-existent item (ID: 888)...")
	err = inv.Remove(888)
	if err != nil {
		fmt.Printf("  [EXPECTED ERROR] Remove ID 888 -> Error: %v\n", err)
	}

	fmt.Println()
	fmt.Println("  Updated Inventory State after Update and Remove:")
	displayInventory(inv.GetAll())

	// -------------------------------------------------------------------------
	// 5. CONCURRENT SUMMARY COMPUTATION USING GOROUTINE & CHANNEL
	// -------------------------------------------------------------------------
	printHeader("STEP 5: Concurrent Summary Statistics via Goroutine & Channel")

	fmt.Println("  >>> Launching concurrent goroutine to compute statistics...")
	startTime := time.Now()
	statsChannel := inv.ComputeSummaryAsync()

	fmt.Println("  >>> Waiting on channel for summary results...")
	summary, ok := <-statsChannel
	if !ok {
		fmt.Println("  [ERROR] Statistics channel was closed unexpectedly.")
		return
	}
	elapsed := time.Since(startTime)

	fmt.Printf("  >>> Summary statistics received from channel in %v!\n\n", elapsed)
	fmt.Println("  +-------------------------------------------------------+")
	fmt.Println("  |            INVENTORY SUMMARY METRICS                  |")
	fmt.Println("  +-------------------------------------------------------+")
	fmt.Printf("  |  Distinct Items Count  : %-28d |\n", summary.TotalItems)
	fmt.Printf("  |  Total Units in Stock  : %-28d |\n", summary.TotalQuantity)
	fmt.Printf("  |  Total Inventory Value : $%-27.2f |\n", summary.TotalValue)
	fmt.Printf("  |  Average Unit Price    : $%-27.2f |\n", summary.AveragePrice)
	fmt.Println("  +-------------------------------------------------------+")

	printHeader("EXECUTION COMPLETED SUCCESSFULLY")
}
