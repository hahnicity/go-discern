package main

import (
    "flag"
    "github.com/hahnicity/go-discern"
    "github.com/hahnicity/go-discern/config"
    "github.com/hahnicity/go-discern/data"
)

func parseArgs(conf *config.Config) {
    flag.IntVar(
        &conf.Processes,
        "p",
        1,
        "The number of parallel processes we want operating",
    )
    flag.IntVar(
        &conf.Workers,
        "workers",
        1,
        "The number of workers we want to have in our pool",
    )
    flag.StringVar(
        &conf.Year,
        "year",
        "2013",
        "The year we wish to look for stats in",
    )
    flag.Float64Var(
        &conf.ViewPercentile,
        "viewp",
        .99,
        "The page view percentile we wish to look for. Must be less than 1",
    )
    flag.Float64Var(
        &conf.MeanPercentile,
        "meanp",
        .75,
        "The mean page view percentile we wish to look for",
    )
    flag.BoolVar(
        &conf.CloseReq,
        "closeRequests",
        false,
        "Set to true if you want to close wiki requests after they have been made",
    )
    flag.Parse()
    if conf.ViewPercentile >= 1.0 || conf.MeanPercentile >= 1.0 {
        panic("Your percentile input must be less than 1")
    }
}

func main() {
    conf := new(config.Config)
    parseArgs(conf)
    work := make(chan discern.WikiRequest)
    go discern.MakeBalancer(conf.Workers).Balance(work)
    discern.Requester(conf, data.SP500, work)
}
