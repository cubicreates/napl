class Student(
    private var rollNo: Int,
    private var name: String,
    private var course: String,
    private var percentage: Double
) {
    public fun acceptData() {
        print("Enter Roll No: ")
        rollNo = readLine()!!.toInt()
        print("Enter Name: ")
        name = readLine()!!
        print("Enter Course: ")
        course = readLine()!!
        print("Enter Percentage: ")
        percentage = readLine()!!.toDouble()
    }

    public fun displayData() {
        println("-------------------------")
        println("Roll No    : $rollNo")
        println("Name       : $name")
        println("Course     : $course")
        println("Percentage : $percentage%")
        println("-------------------------")
    }
}

fun main() {
    val student1 = Student(101, "Aman", "BCA", 85.5)
    val student2 = Student(102, "Neha", "MCA", 90.0)

    println("Initial Student Details:")
    student1.displayData()
    student2.displayData()

    println("\nEnter new details for Student 1:")
    student1.acceptData()

    println("\nEnter new details for Student 2:")
    student2.acceptData()

    println("\nUpdated Student Details:")
    student1.displayData()
    student2.displayData()
}
