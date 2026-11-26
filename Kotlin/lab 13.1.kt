fun main() {
    println("Hello Kotlin!!!")

    println("\n--- Employee Bonus Calculator ---")
    print("Enter department code (IT, HR, SALES): ")
    val deptCode = readLine()!!.uppercase()

    print("Enter years of experience: ")
    val experience = readLine()!!.toInt()

    print("Enter salary: ")
    val salary = readLine()!!.toDouble()

    val bonusPercent = when (deptCode) {
        "IT" -> if (experience > 5) 20 else 10
        "HR" -> if (experience > 3) 15 else 8
        "SALES" -> if (experience > 4) 18 else 10
        else -> 10
    }

    val bonusAmount = salary * bonusPercent / 100.0
    println("Department: $deptCode")
    println("Bonus Percentage: $bonusPercent%")
    println("Computed Bonus Amount: $bonusAmount")

    println("\n--- Average of Digits ---")
    print("Enter an integer: ")
    val num = readLine()!!.toInt()
    var temp = if (num < 0) -num else num
    var sum = 0
    var count = 0

    if (temp == 0) {
        count = 1
    } else {
        while (temp > 0) {
            val digit = temp % 10
            sum += digit
            count++
            temp /= 10
        }
    }

    val average = sum.toDouble() / count
    println("Average of digits: $average")
}
