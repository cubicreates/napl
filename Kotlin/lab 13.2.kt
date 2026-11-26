fun greetUser(name: String = "Guest", role: String = "Student", city: String = "Delhi") {
    println("Hello $name, Role: $role, City: $city")
}

fun printNumbers(vararg numbers: Int) {
    print("Numbers: ")
    for (num in numbers) {
        print("$num ")
    }
    println()
}

fun applyOperation(numbers: List<Int>, operation: (Int) -> Int): List<Int> {
    val result = mutableListOf<Int>()
    for (num in numbers) {
        result.add(operation(num))
    }
    return result
}

fun main() {
    println("--- Default Parameters ---")
    greetUser()
    greetUser("Rahul", "Developer", "Mumbai")

    println("\n--- Named Arguments ---")
    greetUser(city = "Bangalore", name = "Priya", role = "Designer")

    println("\n--- Vararg Function ---")
    printNumbers(10, 20, 30)
    printNumbers(1, 2, 3, 4, 5)

    println("\n--- Higher-Order Function ---")
    val list = listOf(1, 2, 3, 4, 5)
    val doubled = applyOperation(list) { x -> x * 2 }
    println("Original list: $list")
    println("Doubled list: $doubled")
}
