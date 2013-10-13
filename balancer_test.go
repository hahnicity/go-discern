package discern

import (
    "github.com/hahnicity/go-stringit"
    "testing"
    "time"
)

func makeTestRequest() (WikiRequest, chan WikiRequest) {
    resp := make(chan *WikiResponse)
    c := make(chan WikiRequest)
    return makeWikiRequest("2012", "Apple_inc", "AAPL", resp), c
}

func TestWorker(t *testing.T) {
    wr, c := makeTestRequest()
    w := &Worker{c, 0}
    done := make(chan *Worker)
    go w.work(done)
    w.requests <- wr
    <- wr.Resp
}

func TestDispatch(t *testing.T) {
    work := make(chan WikiRequest)
    b := MakeBalancer(10)
    go b.Balance(work)
    wr, _ := makeTestRequest()
    work <- wr
    <- wr.Resp
}

func TestCompleted(t *testing.T) {
    work := make(chan WikiRequest)
    b := MakeBalancer(10)
    go b.Balance(work)
    wr, _ := makeTestRequest()
    work <- wr
    <- wr.Resp
    time.Sleep(time.Second)
    if b.Pool.Len() != 10 {
        panic(stringit.Format("The pool is not of size 10. Size: {}", b.Pool.Len()))
    }
}
