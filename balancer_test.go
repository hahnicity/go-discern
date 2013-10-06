package discern

import "testing"

func TestWorker(t *testing.T) {
    c := make(chan WikiRequest)
    w := &Worker{c, 0}
    wr := makeWikiRequest("2012", "Apple_inc", "AAPL")
    go w.work()
    w.requests <- wr
    <- wr.Resp    
}
