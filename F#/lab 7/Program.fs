open System
open System.Diagnostics

let stopwatch = Stopwatch.StartNew()

let log message =
    Console.WriteLine(sprintf "[%4dms] %s" stopwatch.ElapsedMilliseconds message)

let runTask taskName delayMs resultValue =
    async {
        log (sprintf "Started %s (simulated delay: %dms)" taskName delayMs)
        do! Async.Sleep delayMs
        log (sprintf "Completed %s" taskName)
        return sprintf "%s -> %s" taskName resultValue
    }

let main () =
    log "Launching tasks concurrently using Async.Parallel..."

    let task1 = runTask "Task 1 (Fetch User Data)" 3000 "User profile loaded"
    let task2 = runTask "Task 2 (Query Orders)" 1000 "15 orders retrieved"
    let task3 = runTask "Task 3 (Generate Report)" 2000 "Monthly report ready"

    let results =
        [ task1; task2; task3 ]
        |> Async.Parallel
        |> Async.RunSynchronously

    log "All parallel tasks finished."
    printfn "\nCollected Results:"

    results
    |> Array.iter (fun res -> printfn " - %s" res)

    printfn "\nTotal elapsed time: %dms (sequential would take ~6000ms)" stopwatch.ElapsedMilliseconds

main ()
