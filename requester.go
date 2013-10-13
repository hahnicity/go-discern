package discern

import (
    "fmt"
    "github.com/hahnicity/go-discern/config"
    "github.com/hahnicity/go-stringit"
)


func Requester(conf *config.Config, companies map[string]string, work chan <-WikiRequest) {
    activeRequests := 0
    c := make(chan *WikiResponse)
    ar :=  make([]*WikiResponse, 0)
    for symbol, page := range companies {
        activeRequests++
        req := makeWikiRequest(conf.Year, page, symbol, conf.CloseReq, c)
        work <- req
        ar = manageActiveProc(&activeRequests, conf.Processes, ar, c)
    }
    // Wait for all requests to finish
    for len(ar) < len(companies) {
        resp := <- c
        ar = append(ar, resp)
    }
    Analyze(ar, conf)
}

// Throttle number of active requests
// statsgrok.se is the problem here
func manageActiveProc(activeRequests *int, 
                      maxRequests int, 
                      ar []*WikiResponse, 
                      c chan *WikiResponse) []*WikiResponse {
    if *activeRequests == maxRequests {
        resp := <- c
        ar = append(ar, resp)
    }
    *(activeRequests)--
    return ar
}

func Analyze(ar []*WikiResponse, conf *config.Config) {
    analyzeMeans(ar, conf.MeanPercentile)
    analyzePercentiles(ar, conf.ViewPercentile)
}

func analyzePercentiles(ar []*WikiResponse, viewPercentile float64) {
    for _, resp := range ar {
        dates := FindRecentDates(resp, viewPercentile)
        if len(dates) == 0 {
            return
        }
        fmt.Println(
            stringit.Format(
                "Analyzed {} and found following dates within {} percentile", 
                resp.Symbol, 
                viewPercentile,
            ),
        )
        for date, views := range dates {
            fmt.Println(stringit.Format("\t{}:{}", date, views))    
        }
    }
}

func analyzeMeans(ar []*WikiResponse, meanPercentile float64) {
    means := make(map[string]int)
    for _, resp := range ar {
        means[resp.Symbol] = FindMeanViews(resp)
    }
    fmt.Printf("Companies with mean views within the %f percentile were:\n", meanPercentile)
    for symbol, views := range FindHighestMeans(means, meanPercentile) {
        fmt.Println(stringit.Format("\t{}:{}", symbol, views))      
    }
}
