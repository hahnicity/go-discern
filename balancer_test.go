package discern

import (
    "github.com/hahnicity/go-stringit"
    "testing"
    "time"
)

func TestWorker(t *testing.T) {
    c := make(chan WikiRequest)
    w := &Worker{c, 0}
    wr := makeWikiRequest("2012", "Apple_inc", "AAPL")
    done := make(chan *Worker)
    go w.work(done)
    w.requests <- wr
    <- wr.Resp
}

func TestDispatch(t *testing.T) {
    work := make(chan WikiRequest)
    b := MakeBalancer(10)
    go b.Balance(work)
    wr := makeWikiRequest("2012", "Apple_inc", "AAPL")
    work <- wr
    <- wr.Resp
}

func TestCompleted(t *testing.T) {
    work := make(chan WikiRequest)
    b := MakeBalancer(10)
    go b.Balance(work)
    wr := makeWikiRequest("2012", "Apple_inc", "AAPL")
    work <- wr
    <- wr.Resp
    time.Sleep(time.Second)
    if b.Pool.Len() != 10 {
        panic(stringit.Format("The pool is not of size 10. Size: {}", b.Pool.Len()))
    }
}
