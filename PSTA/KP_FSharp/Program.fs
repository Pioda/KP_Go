// Learn more about F# at http://fsharp.org

open System


let rec test n =
    match n with
    | 0 -> 1
    | x -> n + test (n-1)

let closure n =
    let mutable c = 0
    fun () -> (
        c <- c + n
        c
    )

let curried n = 
    fun c -> c + n 

let pow n = n * n
let negate n = n * -1
let pn = pow >> negate
let np = negate >> pow


let f x = x + 2

let g f x = 2 * f (x)

let h x =
    let mutable f = 2
    f <- 3
    printfn "%i" (f * x)

    
let fib n = 
    let mutable a = 0
    let mutable b = 1
    for i in 0 .. n-1 do
        let temp = a
        a <- b
        b <- temp + a
    printfn "%i" a

let rec fib_rec n =
    match n with
    | 0 | 1 -> n
    | x -> fib_rec (x-1) + fib_rec (x-2)

let rec map nums =
    match nums with
    | [] -> []
    | head::tail -> (head * 2)::map tail

[<EntryPoint>]
let main argv =
    printfn "Hello World from F#!"
    //let x = test (5)
    let a = closure (2)
    let b = curried (4) (5)
    //fib 5
    //printfn "%i" (fib_rec 5)
    //printfn "%A" (map [1;2;3])
    printfn "%i" (a ())
    printfn "%i" (a ())
    printfn "%i" b
    //printfn "%i" (pn (3))
    //printfn "%i" (np (3))
    0 // return an integer exit code