package wiki

import (
    "encoding/json"
    "github.com/hahnicity/go-discern"
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

func composeStats(views chan map[string]int) map[string]int {
    var monthsReceived int = 0
    aggregateViews := make(map[string]int)
    for received := range views {
        monthsReceived += 1
        discern.JoinViewsMaps(aggregateViews, received)
        if monthsReceived == config.NumberMonths {
            close(views)
        }
    }
    return aggregateViews
}

func getMonthlyStats(date string, page string, views chan map[string]int) {
    wikiResp := new(WikiResponse)
    resp, err := http.Get(stringit.Format("{}/{}/{}", config.WikiUrl, date, page))
    if err != nil {
        panic(err)    
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    err = json.Unmarshal(body, wikiResp)
    if err != nil {
        panic(err)    
    }
    views <- wikiResp.Daily_views
}

func GetYearlyStats(year string, page string) map[string]int {
    views := make(chan map[string]int)
    for i := 1; i < 1 + config.NumberMonths; i++ { 
        date := stringit.Format("{}{}", year, numToString(i))
        go getMonthlyStats(date, page, views)
    }
    return composeStats(views)
}

func numToString(number int) string {
    if number < 10 {
        return stringit.Format("0{}", number)    
    } else {
        return stringit.Format("{}", number)    
    }
}
