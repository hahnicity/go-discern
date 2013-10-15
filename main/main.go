package main

import (
    "flag"
    "github.com/hahnicity/go-discern"
    "github.com/hahnicity/go-discern/data"
)

var (
    analyzeMeans   bool
    analyzeTweets  bool
    closeRequests  bool
    maxRequests    int
    meanPercentile float64
    viewPercentile float64
    workers        int
    year           string
)


func parseArgs() {
    flag.IntVar(
        &maxRequests,
        "max",
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
        &viewPercentile,
        "viewp",
        .99,
        "The page view percentile we wish to look for. Must be less than 1",
    )
    flag.BoolVar(
        &closeRequests,
        "closeRequests",
        false,
        "Set to true if you want to close wiki requests after they have been made",
    )
    flag.BoolVar(
        &analyzeMeans,
        "analyzeMeans",
        false,
        "Analyze which companies have the highest means of all the companies surveyed",
    )
    flag.Float64Var(
        &meanPercentile,
        "meanp",
        .75,
        "The mean page view percentile we wish to look for",
    )
    flag.BoolVar(
        &analyzeTweets,
        "analyzeTweets",
        true,
        "Analyze tweets of companies with recent news activity",
    )
    flag.Parse()
    if viewPercentile >= 1.0 || meanPercentile >= 1.0 {
        panic("Your percentile input must be less than 1")
    }
}

func main() {
    parseArgs()
    r := discern.NewRequester(
        analyzeMeans, 
        analyzeTweets, 
        closeRequests, 
        maxRequests, 
        meanPercentile, 
        viewPercentile, 
        year,
    )
    go discern.MakeBalancer(workers).Balance(r.Work)
    r.MakeRequests(data.SP500)
}
