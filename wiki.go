package discern

import (
    "encoding/json"
    "github.com/hahnicity/go-discern/config"
    "github.com/hahnicity/go-stringit"
    "io/ioutil"
    "net/http"
)

type WikiResponse struct {
    Daily_views map[string]int
    Project     string
    Month       string
    Rank        int
    Title       string
}

type WikiRequest struct {
    monthly  chan map[string]int
    yearly   chan map[string]int
    symbol   string
    page     string
    year     string
}

func (w *WikiRequest) composeStats() map[string]int {
    var monthsReceived int = 0
    aggregateViews := make(map[string]int)
    for received := range w.monthly {
        monthsReceived += 1
        JoinViewsMaps(aggregateViews, received)
        if monthsReceived == config.NumberMonths {
            close(w.monthly)
        }
    }
    return aggregateViews
}

func (w *WikiRequest) getMonthlyStats(date string) {
    wikiResp := new(WikiResponse)
    resp, err := http.Get(stringit.Format("{}/{}/{}", config.WikiUrl, date, w.page))
    if err != nil {
        panic(err)    
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(body, wikiResp)
    if err != nil {
        panic(err)    
    }
    w.monthly <- wikiResp.Daily_views
}

func (w *WikiRequest) GetYearlyStats() {
    for i := 1; i < 1 + config.NumberMonths; i++ { 
        date := stringit.Format("{}{}", w.year, toMonthStr(i))
        go w.getMonthlyStats(date)
    }
    w.yearly <- w.composeStats()
}

func toMonthStr(number int) string {
    if number < 10 {
        return stringit.Format("0{}", number)    
    } else {
        return stringit.Format("{}", number)    
    }
}
