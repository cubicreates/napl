package main

import (
	"fmt"
	"strings"
)

// removeByIndex removes an element at a specified index from a slice of strings.
// It returns the updated slice and an error if the index is out of bounds.
func removeByIndex(slice []string, index int) ([]string, error) {
	if index < 0 || index >= len(slice) {
		return slice, fmt.Errorf("index %d is out of range (slice length: %d)", index, len(slice))
	}
	// append elements before index and elements after index
	return append(slice[:index], slice[index+1:]...), nil
}

// printDivider prints a decorative divider line
func printDivider() {
	fmt.Println(strings.Repeat("-", 65))
}

func main() {
	fmt.Println("=================================================================")
	fmt.Println("  EXPERIMENT 3: Collection Management Using Arrays, Slices & Maps ")
	fmt.Println("=================================================================")
	fmt.Println()

	// -------------------------------------------------------------
	// OPTIONAL DEMONSTRATION: Array (Fixed Size)
	// -------------------------------------------------------------
	fmt.Println(">>> DEMO: Fixed-size Array")
	var sampleArray [3]string = [3]string{"Section-A", "Section-B", "Section-C"}
	fmt.Printf("Fixed Array: %v | Type: %T | Length: %d\n\n", sampleArray, sampleArray, len(sampleArray))

	// -------------------------------------------------------------
	// PART 1: Slice Operations (Student Names)
	// Add, Remove by Index, Update, and Print after each step
	// -------------------------------------------------------------
	fmt.Println("=================================================================")
	fmt.Println("  PART 1: Slice Operations (Student Names)")
	fmt.Println("=================================================================")

	// Initial slice
	students := []string{"Alice", "Bob", "Charlie"}
	fmt.Println("[Initial State]")
	fmt.Printf("Slice: %v | Len: %d | Cap: %d\n\n", students, len(students), cap(students))

	// 1. ADD OPERATIONS
	printDivider()
	fmt.Println("1. ADD OPERATION (Appending elements)")
	printDivider()

	// Adding single element
	students = append(students, "Diana")
	fmt.Println("-> Appended 'Diana':")
	fmt.Printf("   Current Slice : %v\n", students)
	fmt.Printf("   Length: %d, Capacity: %d\n\n", len(students), cap(students))

	// Adding multiple elements
	students = append(students, "Evan", "Fiona")
	fmt.Println("-> Appended 'Evan' and 'Fiona':")
	fmt.Printf("   Current Slice : %v\n", students)
	fmt.Printf("   Length: %d, Capacity: %d\n\n", len(students), cap(students))

	// 2. UPDATE OPERATION
	printDivider()
	fmt.Println("2. UPDATE OPERATION (Modifying element at index)")
	printDivider()

	updateIndex := 1
	oldVal := students[updateIndex]
	newVal := "Brian"
	students[updateIndex] = newVal
	fmt.Printf("-> Updated index %d from '%s' to '%s':\n", updateIndex, oldVal, newVal)
	fmt.Printf("   Current Slice : %v\n", students)
	fmt.Printf("   Length: %d, Capacity: %d\n\n", len(students), cap(students))

	// 3. REMOVE-BY-INDEX OPERATION
	printDivider()
	fmt.Println("3. REMOVE-BY-INDEX OPERATION")
	printDivider()

	removeIdx := 2 // Remove "Charlie"
	fmt.Printf("-> Removing element at index %d ('%s')...\n", removeIdx, students[removeIdx])
	var err error
	students, err = removeByIndex(students, removeIdx)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Println("   Removal successful!")
		fmt.Printf("   Current Slice : %v\n", students)
		fmt.Printf("   Length: %d, Capacity: %d\n\n", len(students), cap(students))
	}

	// Boundary check demonstration for remove
	fmt.Println("-> Attempting invalid removal at index 10 (Bounds Checking):")
	_, err = removeByIndex(students, 10)
	if err != nil {
		fmt.Printf("   Caught expected error: %v\n", err)
		fmt.Printf("   Slice remains unchanged: %v\n\n", students)
	}

	// -------------------------------------------------------------
	// PART 2: Map Operations (Subject -> Marks)
	// Insert, Delete, Lookup, and Print after each step
	// -------------------------------------------------------------
	fmt.Println("=================================================================")
	fmt.Println("  PART 2: Map Operations (Subject -> Marks)")
	fmt.Println("=================================================================")

	// Initial map
	marks := map[string]int{
		"Mathematics": 92,
		"Physics":     85,
		"Chemistry":   88,
	}
	fmt.Println("[Initial State]")
	printMap(marks)
	fmt.Println()

	// 1. INSERT OPERATION
	printDivider()
	fmt.Println("1. INSERT OPERATION (Adding new key-value pairs)")
	printDivider()

	fmt.Println("-> Inserting 'Computer Science' with 98 marks:")
	marks["Computer Science"] = 98
	printMap(marks)
	fmt.Println()

	fmt.Println("-> Inserting 'English' with 81 marks:")
	marks["English"] = 81
	printMap(marks)
	fmt.Println()

	// 2. LOOKUP OPERATION
	printDivider()
	fmt.Println("2. LOOKUP OPERATION (Querying keys using comma-ok idiom)")
	printDivider()

	// Existing key lookup
	subjectToLookup := "Mathematics"
	fmt.Printf("-> Looking up '%s':\n", subjectToLookup)
	if val, ok := marks[subjectToLookup]; ok {
		fmt.Printf("   Result: Found! Marks in %s = %d\n", subjectToLookup, val)
	} else {
		fmt.Printf("   Result: '%s' not found in the records.\n", subjectToLookup)
	}
	fmt.Println("   Map state remains:")
	printMap(marks)
	fmt.Println()

	// Non-existing key lookup
	subjectToLookup = "History"
	fmt.Printf("-> Looking up non-existent subject '%s':\n", subjectToLookup)
	if val, ok := marks[subjectToLookup]; ok {
		fmt.Printf("   Result: Found! Marks in %s = %d\n", subjectToLookup, val)
	} else {
		fmt.Printf("   Result: '%s' NOT found in the records.\n", subjectToLookup)
	}
	fmt.Println("   Map state remains:")
	printMap(marks)
	fmt.Println()

	// 3. UPDATE OPERATION (Map)
	printDivider()
	fmt.Println("3. UPDATE OPERATION (Modifying existing key's value)")
	printDivider()
	fmt.Println("-> Updating 'Physics' marks from 85 to 94:")
	marks["Physics"] = 94
	printMap(marks)
	fmt.Println()

	// 4. DELETE OPERATION
	printDivider()
	fmt.Println("4. DELETE OPERATION (Removing key-value pairs)")
	printDivider()

	subjectToDelete := "Chemistry"
	fmt.Printf("-> Deleting '%s' from map...\n", subjectToDelete)
	delete(marks, subjectToDelete)
	fmt.Println("   Deletion completed. Updated map state:")
	printMap(marks)
	fmt.Println()

	// Verify deletion via lookup
	fmt.Printf("-> Verifying deletion of '%s' via lookup:\n", subjectToDelete)
	if _, ok := marks[subjectToDelete]; !ok {
		fmt.Printf("   Confirmed: '%s' is no longer in the map.\n", subjectToDelete)
	}
	printMap(marks)
	fmt.Println()

	fmt.Println("=================================================================")
	fmt.Println("  END OF EXPERIMENT 3 EXECUTION")
	fmt.Println("=================================================================")
}

// printMap displays the contents of a map in a readable format
func printMap(m map[string]int) {
	fmt.Printf("   Map Contents (Count: %d):\n", len(m))
	if len(m) == 0 {
		fmt.Println("     [Empty]")
		return
	}
	for k, v := range m {
		fmt.Printf("     • %-18s : %d\n", k, v)
	}
}
