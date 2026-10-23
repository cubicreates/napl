# Inventory Management Module Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a Go inventory management module integrating Sessions 1–4 concepts (structs, slices, CRUD with validation, concurrent summary calculation using goroutines & channels, and unit tests) under `Assessment 1`.

**Architecture:** Encapsulated `Inventory` struct managing `[]Item` with mutex synchronization, input validation with custom sentinel errors, asynchronous channel-based summary statistics computation, unit tests, and an automated demo runner.

**Tech Stack:** Go 1.26 (Standard Library: `sync`, `errors`, `strings`, `fmt`, `testing`).

**Spec:** `docs/superpowers/specs/2026-10-23-inventory-management-design.md`

## Global Constraints
- Target directory: `Assessment 1/`
- Module name: `assessment1`
- Only standard Go library packages (`sync`, `errors`, `strings`, `fmt`, `testing`).

---

### Task 1: Module Scaffolding and Core Models (`go.mod` and `inventory.go` declarations)

**Files:**
- Create: `Assessment 1/go.mod`
- Create: `Assessment 1/inventory.go`
- Test: `Assessment 1/inventory_test.go`

**Interfaces:**
- Produces: `Item`, `SummaryStats`, `Inventory`, `NewInventory() *Inventory`, sentinel errors (`ErrInvalidID`, `ErrDuplicateID`, `ErrEmptyName`, `ErrNegativeQuantity`, `ErrNegativePrice`, `ErrItemNotFound`).

- [ ] **Step 1: Create go.mod**
- [ ] **Step 2: Write failing unit test for NewInventory and Add**
- [ ] **Step 3: Run test to verify it fails**
- [ ] **Step 4: Implement Item, Inventory, NewInventory, and Add with validation**
- [ ] **Step 5: Run test to verify it passes**

---

### Task 2: Implement Update, Remove, and Query Operations

**Files:**
- Modify: `Assessment 1/inventory.go`
- Test: `Assessment 1/inventory_test.go`

**Interfaces:**
- Consumes: `Inventory`, `Item`, sentinel errors.
- Produces: `Update(id int, name string, quantity int, price float64) error`, `Remove(id int) error`, `Get(id int) (Item, error)`, `GetAll() []Item`.

- [ ] **Step 1: Write unit tests for Update and Remove (including not-found errors and slice shrink)**
- [ ] **Step 2: Run test to verify failure**
- [ ] **Step 3: Implement Update, Remove, Get, and GetAll in inventory.go**
- [ ] **Step 4: Run tests to verify they pass**

---

### Task 3: Implement Concurrent Summary Statistics with Channels

**Files:**
- Modify: `Assessment 1/inventory.go`
- Test: `Assessment 1/inventory_test.go`

**Interfaces:**
- Consumes: `Inventory`, `Item`.
- Produces: `ComputeSummaryAsync() <-chan SummaryStats`, `ComputeSummary() SummaryStats`.

- [ ] **Step 1: Write unit tests for ComputeSummaryAsync testing channel reception, calculation accuracy, and empty inventory handling**
- [ ] **Step 2: Run tests to verify failure**
- [ ] **Step 3: Implement ComputeSummaryAsync and ComputeSummary**
- [ ] **Step 4: Run tests to verify all tests pass (`go test -v ./...`)**

---

### Task 4: Automated Demo Runner (`main.go`) and Complete Verification

**Files:**
- Create: `Assessment 1/main.go`

**Interfaces:**
- Consumes: `Inventory`, `Item`, `NewInventory`, `ComputeSummaryAsync`, error checking.

- [ ] **Step 1: Create main.go demonstrating addition, validation errors, updates, deletions, and channel-based summary statistics**
- [ ] **Step 2: Run `go run .` in Assessment 1 to verify execution output**
- [ ] **Step 3: Run `go test -v ./...` in Assessment 1 to verify full test suite**
- [ ] **Step 4: Final verification and commit**
