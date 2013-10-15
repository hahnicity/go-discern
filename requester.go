package discern

import (
    "fmt"
    "github.com/hahnicity/go-stringit"
)

type Requester struct {
    // The number of active requests that are being run with wikipedia
    activeRequests int

    // All response received from wikipedia
    allResponses   []*WikiResponse

    // Analyze the mean views of all companies on wikipedia
    analyzeMeans   bool

    // Analyze tweets regarding companies that are currently relevant on wikipedia
    analyzeTweets  bool

    // Close all wikipedia requests after they are made
    closeRequests  bool

    // The maximum number of requests that can be run concurrently
    maxRequests    int

    // The percentile at which comapnies with mean views above that should be shown
    meanPercentile float64

    // Channel of active work
    Work           chan WikiRequest

    // The percentile at which a date with number of views above that should be shown
    viewPercentile float64

    // The year to analyze requests
    year           string
}

// Make a new Requester object. 
func NewRequester(analyzeMeans bool,
                  analyzeTweets bool,
                  closeRequests bool, 
                  maxRequests int, 
                  meanPercentile float64, 
                  viewPercentile float64, 
                  year string) (r *Requester){
    r = new(Requester)    
    r.activeRequests = 0
    r.analyzeMeans = analyzeMeans
    r.analyzeTweets = analyzeTweets
    r.closeRequests = closeRequests
    r.maxRequests = maxRequests
    r.meanPercentile = meanPercentile
    r.viewPercentile = viewPercentile
    r.Work = make(chan WikiRequest)
    r.year = year
    return
}

// Given a map of companies and their corresponding wikipedia pages, make
// requests to stats.grok.se so that we can get statistics as to how frequently
// people are viewing their pages
func (r *Requester) MakeRequests(companies map[string]string) {
    c := make(chan *WikiResponse)
    for symbol, page := range companies {
        r.activeRequests++
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

// Throttle number of active requests stats.grok.se is the problem here
func (r *Requester) manageActiveProc(c chan *WikiResponse) {
    if r.activeRequests == r.maxRequests {
        resp := <- c
        r.allResponses = append(r.allResponses, resp)
    }
    r.activeRequests--
}

// Analyze all responses received from wikipedia
func (r *Requester) Analyze() {
    if r.analyzeMeans { r.means() }
    r.percentiles()
}

func (r *Requester) percentiles() {
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

func (r *Requester) means() {
    means := make(map[string]int)
    for _, resp := range r.allResponses {
        means[resp.Symbol] = FindMeanViews(resp)
    }
    fmt.Printf("Companies with mean views within the %f percentile were:\n", r.meanPercentile)
    for symbol, views := range FindHighestMeans(means, r.meanPercentile) {
        fmt.Println(stringit.Format("\t{}:{}", symbol, views))      
    }
}
