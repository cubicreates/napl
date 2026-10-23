# Inventory Management Module Design Document

**Date:** 2026-10-23  
**Status:** Approved  
**Author:** Antigravity  

## 1. Context & Objectives
This project satisfies the practical assignment for Go Sessions 1–4 under `Assessment 1/`.

### Requirements:
1. Define an `Item` struct with fields (`ID`, `Name`, `Quantity`, `Price`) and use a slice to store items.
2. Implement `Add`, `Update`, and `Remove` operations with rigorous input validation and error handling.
3. Add a concurrent goroutine using channels to compute summary statistics (`TotalItems`, `TotalQuantity`, `TotalValue`, `AveragePrice`).
4. Write at least three unit tests in a `_test.go` file covering core functions and error cases.

## 2. Directory Layout
```text
Assessment 1/
├── go.mod                # module assessment1
├── inventory.go          # Core models, validation, CRUD, and concurrent stats
├── inventory_test.go     # Unit tests verifying CRUD, validation errors, and concurrency
└── main.go               # Automated end-to-end demonstration runner
```

## 3. Data Models & Errors
```go
type Item struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    Quantity int     `json:"quantity"`
    Price    float64 `json:"price"`
}

type SummaryStats struct {
    TotalItems    int     `json:"total_items"`
    TotalQuantity int     `json:"total_quantity"`
    TotalValue    float64 `json:"total_value"`
    AveragePrice  float64 `json:"average_price"`
}

type Inventory struct {
    mu    sync.RWMutex
    items []Item
}
```

### Sentinel Errors
- `ErrInvalidID = errors.New("item ID must be greater than zero")`
- `ErrDuplicateID = errors.New("item with this ID already exists")`
- `ErrEmptyName = errors.New("item name cannot be empty")`
- `ErrNegativeQuantity = errors.New("quantity cannot be negative")`
- `ErrNegativePrice = errors.New("price cannot be negative")`
- `ErrItemNotFound = errors.New("item not found")`

## 4. CRUD Operations & Validation
- **`NewInventory() *Inventory`**: Allocates an empty slice.
- **`Add(item Item) error`**:
  - Validates `ID > 0`, non-empty trimmed `Name`, `Quantity >= 0`, `Price >= 0`.
  - Checks duplicate ID.
  - Appends to slice under write lock.
- **`Update(id int, name string, quantity int, price float64) error`**:
  - Validates inputs.
  - Finds item by ID and updates fields under write lock; returns `ErrItemNotFound` if missing.
- **`Remove(id int) error`**:
  - Finds item index and removes it using `append(items[:i], items[i+1:]...)` under write lock.
- **`Get(id int) (Item, error)`**: Returns a copy of the item or `ErrItemNotFound` under read lock.
- **`GetAll() []Item`**: Returns a clone of the internal slice under read lock.

## 5. Concurrency & Channels
- **`ComputeSummaryAsync() <-chan SummaryStats`**:
  - Captures a safe snapshot of `items` under read lock (`RLock`).
  - Spawns a goroutine.
  - Computes `TotalItems`, `TotalQuantity`, `TotalValue`, and `AveragePrice`.
  - Sends the `SummaryStats` result over a dedicated channel and closes it.
- **`ComputeSummary() SummaryStats`**:
  - Blocks on `<-ComputeSummaryAsync()` for synchronous consumption.

## 6. Testing Strategy
- `TestAddItem`: Tests successful additions, duplicate ID errors, invalid ID errors, empty name errors, and negative value errors.
- `TestUpdateAndRemoveItem`: Tests updating item details, removing items, verifying slice reduction, and non-existent item errors.
- `TestComputeSummaryAsync`: Tests concurrent summary computation against expected counts and monetary totals, including on empty inventories.
