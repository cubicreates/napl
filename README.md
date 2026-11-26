# New Age Programming Language

| Details |
| :--- | :--- |
| **Name** | Adit Ghosh |
| **Roll Number** | 2401350016 |
| **Course** | B.Tech CSE (FSD) |
| **Subject** | New Age Programming Language |

## Folder Index

| # | Folder | Section | Overview |
| :-: | :--- | :--- | :--- |
| 1 | [`GO/`](GO/) | [GO](#go) | Go labs (Labs 1–5), Practical Assignment 1 (Inventory Management Module), unit tests, and design/implementation docs. |
| 2 | [`F#/`](F%23/) | [F#](#f) | F# labs (Lab 7, Lab 7.2) on async parallelism and functional list pipelines, plus the SQLite-backed Student Management System assessment. |
| 3 | [`Clojure/`](Clojure/) | [Clojure](#clojure) | Clojure labs (Labs 10.1–11.2) on immutable collections, sequence pipelines, Java interop, and exception handling, plus a Ring/Jetty + SQLite data pipeline assessment. |
| 4 | [`Kotlin/`](Kotlin/) | [Kotlin](#kotlin) | Kotlin labs (Labs 13.1–13.3, Lab 14) covering control flow, functions/lambdas, OOP classes, and a type-safe catalogue builder DSL. |

## File Index & Descriptions

### GO

| # | File Name | Description |
| :-: | :--- | :--- |
| 1 | [`Lab 1/Lab 1/hello.go`](GO/Lab%201/Lab%201/hello.go) | Introductory "Hello, Go!" program demonstrating `package main`, `fmt.Println`, and commented examples of brace-placement and case-sensitivity syntax errors. |
| 2 | [`Lab 1/Lab 2/var.go`](GO/Lab%201/Lab%202/var.go) | Menu-driven console calculator performing arithmetic (`+`, `-`, `*`, `/`) on integers and `float64` values using closures, input validation loops, and division-by-zero checks. |
| 3 | [`Lab 1/Lab 3/strings.go`](GO/Lab%201/Lab%203/strings.go) | Demonstrates Go variable declarations, default zero values, reading user marks via `fmt.Scan`, and `if-else` pass/fail evaluation (`>= 33`). |
| 4 | [`Lab 1/Lab 4/func.go`](GO/Lab%201/Lab%204/func.go) | Demonstrates variadic functions (`sum`, `avg`), multiple return values, and a custom `Celcius` type with a `String()` method and Celsius-to-Fahrenheit conversion. |
| 5 | [`Lab 2/collections.go`](GO/Lab%202/collections.go) | Experiment 3 on Go collections: fixed-size array demo, student-name slice operations (append, update, bounds-checked index removal), and subject-marks map operations (insert, comma-ok lookup, update, delete). |
| 6 | [`Lab 2/Exp 3/collections.go`](GO/Lab%202/Exp%203/collections.go) | Experiment 3 subfolder copy of the array, slice, and map collection management program. |
| 7 | [`Lab 3/main.go`](GO/Lab%203/main.go) | Demonstrates string-to-integer parsing with `strconv.Atoi` and idiomatic `err != nil` error checking. |
| 8 | [`Lab 4/main.go`](GO/Lab%204/main.go) | Experiment 6 executing concurrent goroutines for square, cube, and Fibonacci calculations synchronized with `sync.WaitGroup` and collected via a buffered `TaskResult` channel. |
| 9 | [`Lab 5/go.mod`](GO/Lab%205/go.mod) | Go module file (`module lab5`) configuring Lab 5 for unit testing. |
| 10 | [`Lab 5/main.go`](GO/Lab%205/main.go) | Implements `Divide`, `SquareRoot`, and `ProcessInput` with custom error handling (`fmt.Errorf`) for division by zero, negative inputs, and blank strings. |
| 11 | [`Lab 5/main_test.go`](GO/Lab%205/main_test.go) | Unit tests using `testing.T` subtests (`t.Run`) to verify valid outputs and error cases for `Divide`, `SquareRoot`, and `ProcessInput`. |
| 12 | [`Lab 5/screenshot.png`](GO/Lab%205/screenshot.png) | Screenshot showing execution and unit test output for Lab 5. |
| 13 | [`Assessment 1/go.mod`](GO/Assessment%201/go.mod) | Go module definition (`module assessment1`) for Practical Assignment 1 (Inventory Management Module). |
| 14 | [`Assessment 1/inventory.go`](GO/Assessment%201/inventory.go) | Concurrency-safe `Inventory` module (`sync.RWMutex`) defining `Item` and `SummaryStats` structs, sentinel validation errors, CRUD methods (`Add`, `Update`, `Remove`, `Get`, `GetAll`), and channel-based async summary statistics (`ComputeSummaryAsync`). |
| 15 | [`Assessment 1/main.go`](GO/Assessment%201/main.go) | Automated demonstration runner showcasing item additions, boundary validation errors, updates, removals, and concurrent summary metrics over a channel. |
| 16 | [`Assessment 1/inventory_test.go`](GO/Assessment%201/inventory_test.go) | Unit test suite testing CRUD operations, validation sentinel errors, asynchronous summary channel computation, and concurrent read/write safety. |
| 17 | [`Assessment 1/screenshot.png`](GO/Assessment%201/screenshot.png) | Screenshot capturing test verification output for Assessment 1. |
| 18 | [`Assessment 1/screenshot_run.png`](GO/Assessment%201/screenshot_run.png) | Screenshot capturing the full console demonstration run of Assessment 1. |
| 19 | [`docs/superpowers/specs/2026-10-23-inventory-management-design.md`](GO/docs/superpowers/specs/2026-10-23-inventory-management-design.md) | Design specification outlining models, sentinel errors, CRUD rules, channel concurrency, and testing strategy for Assessment 1. |
| 20 | [`docs/superpowers/plans/2026-10-23-inventory-management.md`](GO/docs/superpowers/plans/2026-10-23-inventory-management.md) | Task-by-task implementation plan for building and testing the Assessment 1 Inventory Management Module. |

### F#

| # | File Name | Description |
| :-: | :--- | :--- |
| 1 | [`lab 7/Lab7.fsproj`](F%23/lab%207/Lab7.fsproj) | .NET (`net11.0`) F# console project file for Lab 7. |
| 2 | [`lab 7/Program.fs`](F%23/lab%207/Program.fs) | Launches simulated asynchronous tasks concurrently with `Async.Sleep` and `Async.Parallel`, logging elapsed execution time via `Stopwatch`. |
| 3 | [`lab 7.2/Lab7_2.fsproj`](F%23/lab%207.2/Lab7_2.fsproj) | .NET (`net11.0`) F# console project file for Lab 7.2. |
| 4 | [`lab 7.2/Program.fs`](F%23/lab%207.2/Program.fs) | Demonstrates functional list processing using pipe (`\|>`) operators with `List.map` (squaring), `List.filter` (evens), `List.fold` (sum and product), and `List.reduce` (maximum). |
| 5 | [`AssmntStudentManagement/AssmntStudentManagement.fsproj`](F%23/AssmntStudentManagement/AssmntStudentManagement.fsproj) | .NET (`net11.0`) F# project file configured with `Microsoft.Data.Sqlite` for the Student Management System assessment. |
| 6 | [`AssmntStudentManagement/Program.fs`](F%23/AssmntStudentManagement/Program.fs) | Menu-driven Student Management System using F# `Student` records, pattern matching, list pipelines (pass/fail status and toppers with marks `>= 75`), and async SQLite CRUD operations. |
| 7 | [`AssmntStudentManagement/student.db`](F%23/AssmntStudentManagement/student.db) | SQLite database file persisting student records (`Id`, `Name`, `Age`, `Marks`) for the F# assessment. |

### Clojure

| # | File Name | Description |
| :-: | :--- | :--- |
| 1 | [`deps.edn`](Clojure/deps.edn) | Root Clojure CLI configuration file for the `Clojure` directory. |
| 2 | [`lab 10.1/exp17.clj`](Clojure/lab%2010.1/exp17.clj) | Experiment 17 demonstrating persistent immutable Lists, Vectors, Maps, and Sets along with core operations (`first`, `rest`, `conj`, `get`, `count`, `assoc`, `dissoc`, `contains?`, `disj`). |
| 3 | [`lab 10.2/exp18.clj`](Clojure/lab%2010.2/exp18.clj) | Experiment 18 applying `map`, `filter`, `reduce`, `take`, `drop`, and thread-last (`->>`) pipelines across vector and list collections. |
| 4 | [`lab 11.1/exp20.clj`](Clojure/lab%2011.1/exp20.clj) | Experiment 20 demonstrating Java interoperability: static method calls (`Math/abs`, `Math/sqrt`, `System/currentTimeMillis`), `ArrayList`/`HashMap` mutation, and bidirectional Clojure–Java collection conversion. |
| 5 | [`lab 11.2/exp23.clj`](Clojure/lab%2011.2/exp23.clj) | Experiment 23 demonstrating `try`/`catch`/`finally` exception handling for `ArithmeticException`, `NullPointerException`, and custom structured errors via `ex-info`, `ex-message`, and `ex-data`. |
| 6 | [`Assesment/deps.edn`](Clojure/Assesment/deps.edn) | Clojure CLI dependency configuration declaring `ring-core`, `ring-jetty-adapter`, and `sqlite-jdbc` for the Clojure assessment. |
| 7 | [`Assesment/data_pipeline.clj`](Clojure/Assesment/data_pipeline.clj) | Data pipeline and Ring/Jetty HTTP server (`localhost:3000/results`) that filters passing students (`marks >= 40`), computes averages, and persists/reads student records in SQLite via JDBC. |
| 8 | [`Assesment/students.db`](Clojure/Assesment/students.db) | SQLite database storing processed student records for the Clojure data pipeline assessment. |

### Kotlin

| # | File Name | Description |
| :-: | :--- | :--- |
| 1 | [`lab 13.1.kt`](Kotlin/lab%2013.1.kt) | Console program implementing an Employee Bonus Calculator (`when` expression across `IT`, `HR`, `SALES` departments and experience) and an Average-of-Digits calculator using a `while` loop. |
| 2 | [`lab 13.2.kt`](Kotlin/lab%2013.2.kt) | Demonstrates Kotlin functions with default parameters, named arguments (`greetUser`), `vararg` parameters (`printNumbers`), and higher-order functions with trailing lambdas (`applyOperation`). |
| 3 | [`lab 13.3.kt`](Kotlin/lab%2013.3.kt) | Demonstrates Kotlin OOP using a `Student` class with private properties (`rollNo`, `name`, `course`, `percentage`) and `acceptData()` / `displayData()` methods. |
| 4 | [`lab 14.kt`](Kotlin/lab%2014.kt) | Builds a type-safe Kotlin DSL (`CatalogueBuilder` / `catalogue { ... }`) using `data class Product`, nullable fields with Elvis operator (`?:`), extension functions (`finalPrice()`, `details()`), and collection pipelines (`filter`, `map`, `reduce`, `groupBy`). |
