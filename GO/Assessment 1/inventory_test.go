package main

import (
	"errors"
	"math"
	"sync"
	"testing"
)

func TestAddItem(t *testing.T) {
	inv := NewInventory()

	// 1. Success case
	item1 := Item{ID: 1, Name: "Laptop", Quantity: 10, Price: 1200.50}
	if err := inv.Add(item1); err != nil {
		t.Fatalf("expected item to be added successfully, got error: %v", err)
	}

	// Verify slice length and item retrieval
	items := inv.GetAll()
	if len(items) != 1 {
		t.Fatalf("expected 1 item in inventory, got %d", len(items))
	}
	if items[0].Name != "Laptop" || items[0].Quantity != 10 || items[0].Price != 1200.50 {
		t.Errorf("item stored does not match added item: %+v", items[0])
	}

	// 2. Duplicate ID validation
	duplicateItem := Item{ID: 1, Name: "Gaming Laptop", Quantity: 5, Price: 1500.00}
	if err := inv.Add(duplicateItem); !errors.Is(err, ErrDuplicateID) {
		t.Errorf("expected ErrDuplicateID, got: %v", err)
	}

	// 3. Invalid ID (<= 0)
	invalidIDItem := Item{ID: 0, Name: "Monitor", Quantity: 2, Price: 300.00}
	if err := inv.Add(invalidIDItem); !errors.Is(err, ErrInvalidID) {
		t.Errorf("expected ErrInvalidID, got: %v", err)
	}

	// 4. Empty Name
	emptyNameItem := Item{ID: 2, Name: "   ", Quantity: 5, Price: 50.00}
	if err := inv.Add(emptyNameItem); !errors.Is(err, ErrEmptyName) {
		t.Errorf("expected ErrEmptyName, got: %v", err)
	}

	// 5. Negative Quantity
	negQtyItem := Item{ID: 3, Name: "Keyboard", Quantity: -1, Price: 75.00}
	if err := inv.Add(negQtyItem); !errors.Is(err, ErrNegativeQuantity) {
		t.Errorf("expected ErrNegativeQuantity, got: %v", err)
	}

	// 6. Negative Price
	negPriceItem := Item{ID: 4, Name: "Mouse", Quantity: 10, Price: -20.00}
	if err := inv.Add(negPriceItem); !errors.Is(err, ErrNegativePrice) {
		t.Errorf("expected ErrNegativePrice, got: %v", err)
	}
}

func TestUpdateItem(t *testing.T) {
	inv := NewInventory()

	_ = inv.Add(Item{ID: 1, Name: "Mechanical Keyboard", Quantity: 15, Price: 80.00})
	_ = inv.Add(Item{ID: 2, Name: "Wireless Mouse", Quantity: 25, Price: 35.00})

	// 1. Successful update
	err := inv.Update(1, "RGB Mechanical Keyboard", 12, 95.50)
	if err != nil {
		t.Fatalf("unexpected error updating item: %v", err)
	}

	updated, err := inv.Get(1)
	if err != nil {
		t.Fatalf("unexpected error fetching item 1: %v", err)
	}
	if updated.Name != "RGB Mechanical Keyboard" || updated.Quantity != 12 || updated.Price != 95.50 {
		t.Errorf("updated fields do not match expected: %+v", updated)
	}

	// 2. Non-existent item update
	err = inv.Update(99, "Unknown", 5, 10.00)
	if !errors.Is(err, ErrItemNotFound) {
		t.Errorf("expected ErrItemNotFound, got: %v", err)
	}

	// 3. Update with empty name
	err = inv.Update(1, "   ", 10, 50.00)
	if !errors.Is(err, ErrEmptyName) {
		t.Errorf("expected ErrEmptyName, got: %v", err)
	}

	// 4. Update with negative quantity
	err = inv.Update(1, "Valid Name", -5, 50.00)
	if !errors.Is(err, ErrNegativeQuantity) {
		t.Errorf("expected ErrNegativeQuantity, got: %v", err)
	}

	// 5. Update with negative price
	err = inv.Update(1, "Valid Name", 5, -10.00)
	if !errors.Is(err, ErrNegativePrice) {
		t.Errorf("expected ErrNegativePrice, got: %v", err)
	}
}

func TestRemoveItem(t *testing.T) {
	inv := NewInventory()

	_ = inv.Add(Item{ID: 1, Name: "Item A", Quantity: 10, Price: 10.00})
	_ = inv.Add(Item{ID: 2, Name: "Item B", Quantity: 20, Price: 20.00})
	_ = inv.Add(Item{ID: 3, Name: "Item C", Quantity: 30, Price: 30.00})

	// 1. Remove middle element
	if err := inv.Remove(2); err != nil {
		t.Fatalf("unexpected error removing item 2: %v", err)
	}

	items := inv.GetAll()
	if len(items) != 2 {
		t.Fatalf("expected 2 items after removal, got %d", len(items))
	}
	if items[0].ID != 1 || items[1].ID != 3 {
		t.Errorf("remaining items mismatch: %+v", items)
	}

	// 2. Verify Get(2) returns ErrItemNotFound
	_, err := inv.Get(2)
	if !errors.Is(err, ErrItemNotFound) {
		t.Errorf("expected ErrItemNotFound for removed item, got: %v", err)
	}

	// 3. Remove non-existent ID
	err = inv.Remove(999)
	if !errors.Is(err, ErrItemNotFound) {
		t.Errorf("expected ErrItemNotFound for non-existent item, got: %v", err)
	}

	// 4. Remove remaining items
	if err := inv.Remove(1); err != nil {
		t.Fatalf("error removing item 1: %v", err)
	}
	if err := inv.Remove(3); err != nil {
		t.Fatalf("error removing item 3: %v", err)
	}

	if len(inv.GetAll()) != 0 {
		t.Errorf("expected empty inventory, got len %d", len(inv.GetAll()))
	}

	// 5. Remove from empty inventory
	err = inv.Remove(1)
	if !errors.Is(err, ErrItemNotFound) {
		t.Errorf("expected ErrItemNotFound when removing from empty inventory, got: %v", err)
	}
}

func TestConcurrentSummaryStats(t *testing.T) {
	inv := NewInventory()

	// 1. Test summary on empty inventory
	emptyStatsChan := inv.ComputeSummaryAsync()
	emptyStats, ok := <-emptyStatsChan
	if !ok {
		t.Fatal("expected value on emptyStatsChan, got closed channel")
	}
	if emptyStats.TotalItems != 0 || emptyStats.TotalQuantity != 0 || emptyStats.TotalValue != 0.0 || emptyStats.AveragePrice != 0.0 {
		t.Errorf("expected zero stats on empty inventory, got: %+v", emptyStats)
	}

	// Channel should be closed now
	if _, ok := <-emptyStatsChan; ok {
		t.Error("expected emptyStatsChan to be closed")
	}

	// 2. Populate inventory
	_ = inv.Add(Item{ID: 1, Name: "Item 1", Quantity: 5, Price: 100.0})
	_ = inv.Add(Item{ID: 2, Name: "Item 2", Quantity: 10, Price: 50.0})
	_ = inv.Add(Item{ID: 3, Name: "Item 3", Quantity: 2, Price: 250.0})

	// 3. Compute summary asynchronously through channel
	statsChan := inv.ComputeSummaryAsync()
	stats, ok := <-statsChan
	if !ok {
		t.Fatal("expected value on statsChan, channel was already closed")
	}

	expectedTotalItems := 3
	expectedTotalQty := 17
	expectedTotalValue := 1500.0
	expectedAvgPrice := 400.0 / 3.0 // 133.33333333333334

	if stats.TotalItems != expectedTotalItems {
		t.Errorf("expected TotalItems %d, got %d", expectedTotalItems, stats.TotalItems)
	}
	if stats.TotalQuantity != expectedTotalQty {
		t.Errorf("expected TotalQuantity %d, got %d", expectedTotalQty, stats.TotalQuantity)
	}
	if math.Abs(stats.TotalValue-expectedTotalValue) > 0.0001 {
		t.Errorf("expected TotalValue %v, got %v", expectedTotalValue, stats.TotalValue)
	}
	if math.Abs(stats.AveragePrice-expectedAvgPrice) > 0.0001 {
		t.Errorf("expected AveragePrice %v, got %v", expectedAvgPrice, stats.AveragePrice)
	}

	// 4. Test synchronous wrapper
	syncStats := inv.ComputeSummary()
	if syncStats != stats {
		t.Errorf("expected ComputeSummary() to match ComputeSummaryAsync(), got %+v vs %+v", syncStats, stats)
	}
}

func TestConcurrentAccessSafety(t *testing.T) {
	inv := NewInventory()
	var wg sync.WaitGroup

	// Launch multiple concurrent writers and readers
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_ = inv.Add(Item{
				ID:       id,
				Name:     "Concurrent Item",
				Quantity: id * 2,
				Price:    float64(id * 10),
			})
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stats := inv.ComputeSummary()
			if stats.TotalItems < 0 || stats.TotalQuantity < 0 {
				t.Errorf("invalid stats observed: %+v", stats)
			}
		}()
	}

	wg.Wait()

	finalStats := inv.ComputeSummary()
	if finalStats.TotalItems != 20 {
		t.Errorf("expected 20 items after concurrent additions, got %d", finalStats.TotalItems)
	}
}
