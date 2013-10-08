package discern

import (
    "encoding/json"
    "fmt"
    "github.com/hahnicity/go-discern/config"
    "github.com/hahnicity/go-stringit"
    "io/ioutil"
    "net/http"
)

type MonthlyStats struct {
    Daily_views map[string]int
    Project     string
    Month       string
    Rank        int
    Title       string
}

type WikiRequest struct {
    Client *http.Client
    Resp   chan *WikiResponse
    Symbol string
    Page   string
    Year   string
}

type WikiResponse struct {
    Symbol string
    Yearly map[string]int
}

func (w *WikiRequest) composeStats(monthly chan map[string]int) map[string]int {
    var monthsReceived int = 0
    aggregateViews := make(map[string]int)
    for received := range monthly {
        monthsReceived += 1
        JoinViewsMaps(aggregateViews, received)
        if monthsReceived >= config.NumberMonths {
            close(monthly)
        }
    }
    return aggregateViews
}

func (w *WikiRequest) getMonthlyStats(date string, monthly chan map[string]int) {
    resp, err := http.Get(stringit.Format("{}/{}/{}", config.WikiUrl, date, w.Page))
    if err != nil {
        panic(err)    
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)    
    }
    monthlyStats := new(MonthlyStats)
    err = json.Unmarshal(body, monthlyStats)
    if err != nil {
        panic(err)    
    }
    monthly <- monthlyStats.Daily_views
    return  // tell the goroutine to close
}

func (w *WikiRequest) GetYearlyStats() *WikiResponse {
    fmt.Println(stringit.Format("Search for wiki stats for {}", w.Symbol))
    monthly := make(chan map[string]int)
    for i := 1; i < 1 + config.NumberMonths; i++ { 
        date := stringit.Format("{}{}", w.Year, toMonthStr(i))
        go w.getMonthlyStats(date, monthly)
    }
    return &WikiResponse{w.Symbol, w.composeStats(monthly)}
}

func toMonthStr(number int) string {
    if number < 10 {
        return stringit.Format("0{}", number)    
    } else {
        return stringit.Format("{}", number)    
    }
}

func makeWikiRequest(year, page, symbol string) WikiRequest{
    wikiResp := make(chan *WikiResponse)
    tr := &http.Transport{DisableKeepAlives: true}
    c := &http.Client{Transport: tr}
    return WikiRequest{c, wikiResp, symbol, page, year}
}
