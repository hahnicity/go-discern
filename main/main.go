package main

import (
    "flag"
    "github.com/hahnicity/go-discern"
    "github.com/hahnicity/go-discern/data"
)

var (
    processes int
    workers   int
    year      string
)

func parseArgs() {
    flag.IntVar(
        &processes,
        "p",
        20,
        "The number of parallel processes we want operating",
    )
    flag.IntVar(
        &workers,
        "workers",
        50,
        "The number of workers we want to have in our pool",
    )
    flag.StringVar(
        &year,
        "year",
        "2013",
        "The year we wish to look for stats in",
    )
    flag.Parse()
}

func main() {
    parseArgs()
    work := make(chan discern.WikiRequest)
    go discern.MakeBalancer(workers).Balance(work)
    discern.Requester(year, data.SP500, processes, work)
}
