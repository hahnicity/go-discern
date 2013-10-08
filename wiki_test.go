package discern

import "testing"

func TestGetMonthlyStats(t *testing.T) {
    wr := makeWikiRequest("2012", "Apple_inc", "APPL")       
    mo := make(chan map[string]int)
    go wr.getMonthlyStats("201201", mo)
    <- mo
}

func TestGetYearlyStats(t *testing.T) {
    wr := makeWikiRequest("2012", "Apple_inc", "APPL")       
    wr.GetYearlyStats()
}
