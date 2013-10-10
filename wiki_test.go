package discern

import "testing"

func TestGetMonthlyStats(t *testing.T) {
    wr, _ := makeTestRequest()
    mo := make(chan map[string]int)
    go wr.getMonthlyStats("201201", mo)
    <- mo
}

func TestGetYearlyStats(t *testing.T) {
    wr, _ := makeTestRequest()
    wr.GetYearlyStats()
}
