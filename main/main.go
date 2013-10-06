package main

import (
    "github.com/hahnicity/go-discern"
    "github.com/hahnicity/go-discern/data"
)

//XXX Make argument parser

func main() {
    work := make(chan discern.WikiRequest)
    b := discern.MakeBalancer(600)
    go b.Balance(work)
    discern.Requester("2012", data.SP500, 0, work)
}
