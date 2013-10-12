package discern

import (
    "fmt"
    "github.com/hahnicity/go-discern/config"
    "github.com/hahnicity/go-stringit"
)


func Requester(conf *config.Config, companies map[string]string, work chan <-WikiRequest) {
    items := 0
    c := make(chan *WikiResponse)
    ar :=  make([]*WikiResponse, 0)
    for symbol, page := range companies {
        items++
        req := makeWikiRequest(conf.Year, page, symbol, conf.CloseReq, c)
        work <- req
        if items > conf.Processes { // statsgrok.se is the problem here
            resp := <- c
            ar = append(ar, resp)
            items--
        }
    }
    Analyze(ar, conf)
}

func Analyze(ar []*WikiResponse, conf *config.Config) {
    analyzeMeans(ar, conf.MeanPercentile)
    analyzePercentiles(ar, conf.ViewPercentile)
}

func analyzePercentiles(ar []*WikiResponse, viewPercentile float64) {
    for _, resp := range ar {
        dates := FindRecentDates(resp, viewPercentile)
        if len(*dates) == 0 {
            return
        }
        fmt.Println(
            stringit.Format(
                "Analyzed {} and found following dates within {} percentile", 
                resp.Symbol, 
                viewPercentile,
            ),
        )
        for _, date := range *dates {
            fmt.Println(stringit.Format("\t{}:{}", date, resp.Yearly[date]))    
        }
    }
}

func analyzeMeans(ar []*WikiResponse, meanPercentile float64) {
    means := make(map[string]int)
    for _, resp := range ar {
        means[resp.Symbol] = FindMeanViews(resp)
    }
    fmt.Println("Companies with the Mean Views within the %f percentile were:", meanPercentile)
    for symbol, views := range FindHighestMeans(means, meanPercentile) {
        fmt.Println(stringit.Format("\t{}:{}", symbol, views))      
    }
}
