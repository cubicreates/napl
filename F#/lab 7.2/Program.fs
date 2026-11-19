let list1 = [ 1; 2; 3; 4; 5; 6 ]
let list2 = [ -3; 0; 7; 10; 15; 22 ]

printfn "=== 1. List.map (Square each element) ==="
let squared1 = list1 |> List.map (fun x -> x * x)
let squared2 = list2 |> List.map (fun x -> x * x)

printfn "List 1: %A -> Squared: %A" list1 squared1
printfn "List 2: %A -> Squared: %A" list2 squared2

printfn "\n=== 2. List.filter (Retain only even numbers) ==="
let evens1 = list1 |> List.filter (fun x -> x % 2 = 0)
let evens2 = list2 |> List.filter (fun x -> x % 2 = 0)

printfn "List 1: %A -> Evens: %A" list1 evens1
printfn "List 2: %A -> Evens: %A" list2 evens2

printfn "\n=== 3. List.fold (Compute sum and product) ==="
let sum1 = list1 |> List.fold (fun acc x -> acc + x) 0
let sum2 = list2 |> List.fold (fun acc x -> acc + x) 0

let product1 = list1 |> List.fold (fun acc x -> acc * x) 1
let max2 = list2 |> List.reduce (fun acc x -> if x > acc then x else acc)

printfn "List 1: %A -> Sum: %d | Product: %d" list1 sum1 product1
printfn "List 2: %A -> Sum: %d | Maximum: %d" list2 sum2 max2
