open System
open Microsoft.Data.Sqlite

type Student = {
    Id: int
    Name: string
    Age: int
    Marks: float
}

let connectionString = "Data Source=student.db"

let createDatabase () =
    use connection = new SqliteConnection(connectionString)
    connection.Open()

    use command = connection.CreateCommand()

    command.CommandText <-
        "CREATE TABLE IF NOT EXISTS Students (
            Id INTEGER PRIMARY KEY,
            Name TEXT NOT NULL,
            Age INTEGER NOT NULL,
            Marks REAL NOT NULL
        )"

    command.ExecuteNonQuery() |> ignore

let addStudents students =
    async {
        use connection = new SqliteConnection(connectionString)
        do! connection.OpenAsync() |> Async.AwaitTask

        let mutable totalAdded = 0

        for student in students do
            use command = connection.CreateCommand()

            command.CommandText <-
                "INSERT INTO Students (Id, Name, Age, Marks)
                 VALUES ($id, $name, $age, $marks)"

            command.Parameters.AddWithValue("$id", student.Id) |> ignore
            command.Parameters.AddWithValue("$name", student.Name) |> ignore
            command.Parameters.AddWithValue("$age", student.Age) |> ignore
            command.Parameters.AddWithValue("$marks", student.Marks) |> ignore

            let! result = command.ExecuteNonQueryAsync() |> Async.AwaitTask
            totalAdded <- totalAdded + result

        return totalAdded
    }

let getStudents () =
    async {
        use connection = new SqliteConnection(connectionString)
        do! connection.OpenAsync() |> Async.AwaitTask

        use command = connection.CreateCommand()
        command.CommandText <- "SELECT * FROM Students"

        let! dataReader =
            command.ExecuteReaderAsync() |> Async.AwaitTask

        use reader = dataReader

        let mutable students = []

        while reader.Read() do
            let student = {
                Id = reader.GetInt32(0)
                Name = reader.GetString(1)
                Age = reader.GetInt32(2)
                Marks = reader.GetDouble(3)
            }

            students <- student :: students

        return List.rev students
    }

let updateStudent student =
    async {
        use connection = new SqliteConnection(connectionString)
        do! connection.OpenAsync() |> Async.AwaitTask

        use command = connection.CreateCommand()

        command.CommandText <-
            "UPDATE Students
             SET Name = $name, Age = $age, Marks = $marks
             WHERE Id = $id"

        command.Parameters.AddWithValue("$id", student.Id) |> ignore
        command.Parameters.AddWithValue("$name", student.Name) |> ignore
        command.Parameters.AddWithValue("$age", student.Age) |> ignore
        command.Parameters.AddWithValue("$marks", student.Marks) |> ignore

        let! result = command.ExecuteNonQueryAsync() |> Async.AwaitTask

        return result
    }

let deleteStudent id =
    async {
        use connection = new SqliteConnection(connectionString)
        do! connection.OpenAsync() |> Async.AwaitTask

        use command = connection.CreateCommand()

        command.CommandText <- "DELETE FROM Students WHERE Id = $id"
        command.Parameters.AddWithValue("$id", id) |> ignore

        let! result = command.ExecuteNonQueryAsync() |> Async.AwaitTask

        return result
    }

let getResult marks =
    match marks with
    | m when m >= 40.0 -> "Pass"
    | _ -> "Fail"

let displayStudents students =
    students
    |> List.iter (fun s ->
        printfn "ID: %d | Name: %s | Age: %d | Marks: %.2f | Result: %s"
            s.Id s.Name s.Age s.Marks (getResult s.Marks)
    )

let main () =
    createDatabase ()

    let mutable choice = 0

    while choice <> 6 do
        printfn "\nSTUDENT MANAGEMENT SYSTEM"
        printfn "1. Add Student"
        printfn "2. View Students"
        printfn "3. Update Student"
        printfn "4. Delete Student"
        printfn "5. View Toppers"
        printfn "6. Exit"

        printf "Enter choice: "
        choice <- int (Console.ReadLine())

        match choice with

        | 1 ->
            printf "How many students do you want to add? "
            let count = int (Console.ReadLine())

            let mutable students = []

            for i in 1 .. count do
                printfn "\nEnter details for Student %d:" i

                printf "Enter ID: "
                let id = int (Console.ReadLine())

                printf "Enter Name: "
                let name = Console.ReadLine()

                printf "Enter Age: "
                let age = int (Console.ReadLine())

                printf "Enter Marks: "
                let marks = float (Console.ReadLine())

                let student = {
                    Id = id
                    Name = name
                    Age = age
                    Marks = marks
                }

                students <- student :: students

            let result =
                List.rev students
                |> addStudents
                |> Async.RunSynchronously

            printfn "%d student(s) added successfully." result

        | 2 ->
            let students =
                getStudents ()
                |> Async.RunSynchronously

            displayStudents students

        | 3 ->
            printf "Enter Student ID: "
            let id = int (Console.ReadLine())

            printf "Enter New Name: "
            let name = Console.ReadLine()

            printf "Enter New Age: "
            let age = int (Console.ReadLine())

            printf "Enter New Marks: "
            let marks = float (Console.ReadLine())

            let student = {
                Id = id
                Name = name
                Age = age
                Marks = marks
            }

            let result =
                updateStudent student
                |> Async.RunSynchronously

            match result with
            | 1 -> printfn "Student updated successfully."
            | _ -> printfn "Student not found."

        | 4 ->
            printf "Enter Student ID: "
            let id = int (Console.ReadLine())

            let result =
                deleteStudent id
                |> Async.RunSynchronously

            match result with
            | 1 -> printfn "Student deleted successfully."
            | _ -> printfn "Student not found."

        | 5 ->
            let students =
                getStudents ()
                |> Async.RunSynchronously

            let toppers =
                students
                |> List.filter (fun s -> s.Marks >= 75.0)
                |> List.map (fun s -> s.Name)

            printfn "Topper Students: %A" toppers

        | 6 ->
            printfn "Exiting application."

        | _ ->
            printfn "Invalid choice."

main()
