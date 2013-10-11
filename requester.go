package discern

import (
    "fmt"
    "github.com/hahnicity/go-stringit"
)


func Requester(year string, 
               companies map[string]string, 
               n int, 
               f float64, 
               cl bool, 
               work chan <-WikiRequest) {
    items := 0
    c := make(chan *WikiResponse)
    for symbol, page := range companies {
        items++
        req := makeWikiRequest(year, page, symbol, cl, c)
        work <- req
        if items > n { // statsgrok.se is the problem here
            resp := <- c
            analyzeResponse(resp, f)
            items--
        }
    }
}

func analyzeResponse(resp *WikiResponse, f float64) {
    dates := FindDates(resp, f)
    fmt.Println(
        stringit.Format(
            "Analyzed {} and found following dates within {} percentile", 
            resp.Symbol, 
            f,
        ),
    )
    for _, date := range *dates {
        fmt.Println(stringit.Format("\t {}:{}", date, resp.Yearly[date]))    
    }
}
