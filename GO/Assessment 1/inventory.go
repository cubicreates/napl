package main

import (
	"errors"
	"strings"
	"sync"
)

// Sentinel error definitions for input validation and operational errors
var (
	ErrInvalidID        = errors.New("item ID must be greater than zero")
	ErrDuplicateID      = errors.New("item with this ID already exists")
	ErrEmptyName        = errors.New("item name cannot be empty")
	ErrNegativeQuantity = errors.New("quantity cannot be negative")
	ErrNegativePrice    = errors.New("price cannot be negative")
	ErrItemNotFound     = errors.New("item not found")
)

// Item represents a product record in the inventory.
type Item struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// SummaryStats holds aggregated summary statistics for the inventory.
type SummaryStats struct {
	TotalItems    int     `json:"total_items"`
	TotalQuantity int     `json:"total_quantity"`
	TotalValue    float64 `json:"total_value"`
	AveragePrice  float64 `json:"average_price"`
}

// Inventory encapsulates a slice of items with concurrency-safe access.
type Inventory struct {
	mu    sync.RWMutex
	items []Item
}

// NewInventory initializes and returns an empty Inventory instance.
func NewInventory() *Inventory {
	return &Inventory{
		items: make([]Item, 0),
	}
}

// validateItem checks that an item's fields satisfy all validation rules.
func validateItem(item Item) error {
	if item.ID <= 0 {
		return ErrInvalidID
	}
	if strings.TrimSpace(item.Name) == "" {
		return ErrEmptyName
	}
	if item.Quantity < 0 {
		return ErrNegativeQuantity
	}
	if item.Price < 0 {
		return ErrNegativePrice
	}
	return nil
}

// Add inserts a new item into the inventory after validation and uniqueness check.
func (inv *Inventory) Add(item Item) error {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	if err := validateItem(item); err != nil {
		return err
	}

	for _, existing := range inv.items {
		if existing.ID == item.ID {
			return ErrDuplicateID
		}
	}

	// Clean name string
	item.Name = strings.TrimSpace(item.Name)
	inv.items = append(inv.items, item)
	return nil
}

// Update modifies an existing item's attributes after validating the new values.
func (inv *Inventory) Update(id int, name string, quantity int, price float64) error {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return ErrEmptyName
	}
	if quantity < 0 {
		return ErrNegativeQuantity
	}
	if price < 0 {
		return ErrNegativePrice
	}

	for i, existing := range inv.items {
		if existing.ID == id {
			inv.items[i].Name = trimmedName
			inv.items[i].Quantity = quantity
			inv.items[i].Price = price
			return nil
		}
	}

	return ErrItemNotFound
}

// Remove deletes an item by its ID using slice re-slicing.
func (inv *Inventory) Remove(id int) error {
	inv.mu.Lock()
	defer inv.mu.Unlock()

	for i, existing := range inv.items {
		if existing.ID == id {
			inv.items = append(inv.items[:i], inv.items[i+1:]...)
			return nil
		}
	}

	return ErrItemNotFound
}

// Get retrieves a copy of an item by its ID.
func (inv *Inventory) Get(id int) (Item, error) {
	inv.mu.RLock()
	defer inv.mu.RUnlock()

	for _, item := range inv.items {
		if item.ID == id {
			return item, nil
		}
	}

	return Item{}, ErrItemNotFound
}

// GetAll returns a defensive copy of all items in the inventory.
func (inv *Inventory) GetAll() []Item {
	inv.mu.RLock()
	defer inv.mu.RUnlock()

	result := make([]Item, len(inv.items))
	copy(result, inv.items)
	return result
}

// ComputeSummaryAsync launches a concurrent goroutine to compute inventory statistics
// and returns a receive-only channel through which the result is delivered.
func (inv *Inventory) ComputeSummaryAsync() <-chan SummaryStats {
	statsChan := make(chan SummaryStats, 1)

	// Take a defensive snapshot of items under read lock so the computation can
	// proceed asynchronously without holding the lock or causing race conditions.
	inv.mu.RLock()
	snapshot := make([]Item, len(inv.items))
	copy(snapshot, inv.items)
	inv.mu.RUnlock()

	go func(items []Item) {
		defer close(statsChan)

		var totalQty int
		var totalVal float64
		var totalPriceSum float64

		for _, item := range items {
			totalQty += item.Quantity
			totalVal += float64(item.Quantity) * item.Price
			totalPriceSum += item.Price
		}

		avgPrice := 0.0
		if len(items) > 0 {
			avgPrice = totalPriceSum / float64(len(items))
		}

		statsChan <- SummaryStats{
			TotalItems:    len(items),
			TotalQuantity: totalQty,
			TotalValue:    totalVal,
			AveragePrice:  avgPrice,
		}
	}(snapshot)

	return statsChan
}

// ComputeSummary blocks until the concurrent calculation via ComputeSummaryAsync completes.
func (inv *Inventory) ComputeSummary() SummaryStats {
	return <-inv.ComputeSummaryAsync()
}
