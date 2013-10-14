package discern

import (
    "fmt"
    "github.com/hahnicity/go-stringit"
)

type Requester struct {
    activeRequests int
    allResponses   []*WikiResponse
    closeRequests  bool
    maxRequests    int
    meanPercentile float64
    Work           chan WikiRequest
    viewPercentile float64
    year           string
}

func NewRequester(closeRequests bool, maxRequests int, meanPercentile, viewPercentile float64, year string) (r *Requester){
    r = new(Requester)    
    r.closeRequests = closeRequests
    r.maxRequests = maxRequests
    r.meanPercentile = meanPercentile
    r.viewPercentile = viewPercentile
    r.Work = make(chan WikiRequest)
    r.year = year
    return
}

func (r *Requester) MakeRequests(companies map[string]string) {
    activeRequests := 0
    c := make(chan *WikiResponse)
    for symbol, page := range companies {
        activeRequests++
        r.Work <- makeWikiRequest(r.year, page, symbol, r.closeRequests, c)
        r.manageActiveProc(c)
    }
    // Wait for all requests to finish
    for len(r.allResponses) < len(companies) {
        resp := <- c
        r.allResponses = append(r.allResponses, resp)
    }
    r.Analyze()
}

// Throttle number of active requests
// statsgrok.se is the problem here
func (r *Requester) manageActiveProc(c chan *WikiResponse) {
    if r.activeRequests == r.maxRequests {
        resp := <- c
        r.allResponses = append(r.allResponses, resp)
    }
    r.activeRequests--
}

func (r *Requester) Analyze() {
    r.analyzeMeans()
    r.analyzePercentiles()
}

func (r *Requester) analyzePercentiles() {
    for _, resp := range r.allResponses {
        dates := FindRecentDates(resp, r.viewPercentile)
        if len(dates) == 0 {
            return
        }
        fmt.Println(
            stringit.Format(
                "Analyzed {} and found following dates within {} percentile", 
                resp.Symbol, 
                r.viewPercentile,
            ),
        )
        for date, views := range dates {
            fmt.Println(stringit.Format("\t{}:{}", date, views))    
        }
    }
}

func (r *Requester) analyzeMeans() {
    means := make(map[string]int)
    for _, resp := range r.allResponses {
        means[resp.Symbol] = FindMeanViews(resp)
    }
    fmt.Printf("Companies with mean views within the %f percentile were:\n", r.meanPercentile)
    for symbol, views := range FindHighestMeans(means, r.meanPercentile) {
        fmt.Println(stringit.Format("\t{}:{}", symbol, views))      
    }
}
