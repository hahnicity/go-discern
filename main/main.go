package main

import (
    "flag"
    "github.com/hahnicity/go-discern"
    "github.com/hahnicity/go-discern/data"
)

var (
    cl        bool
    f         float64
    processes int
    workers   int
    year      string
)

func parseArgs() {
    flag.IntVar(
        &processes,
        "p",
        1,
        "The number of parallel processes we want operating",
    )
    flag.IntVar(
        &workers,
        "workers",
        1,
        "The number of workers we want to have in our pool",
    )
    flag.StringVar(
        &year,
        "year",
        "2013",
        "The year we wish to look for stats in",
    )
    flag.Float64Var(
        &f,
        "percentile",
        .99,
        "The page view percentile we wish to look for. Must be less than 1",
    )
    flag.BoolVar(
        &cl,
        "closeRequests",
        false,
        "Set to true if you want to close wiki requests after they have been made",
    )
    flag.Parse()
}

func main() {
    parseArgs()
    work := make(chan discern.WikiRequest)
    go discern.MakeBalancer(workers).Balance(work)
    if f >= 1.0 {panic("The percentile input must be less than 1")}
    discern.Requester(year, data.SP500, processes - 1, f, cl, work)
}
