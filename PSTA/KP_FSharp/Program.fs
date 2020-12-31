// Learn more about F# at http://fsharp.org

open System


let rec test n =
    match n with
    | 0 -> 1
    | x -> n + test (n-1)


