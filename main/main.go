package main

import (
    "container/heap"
    "fmt"
    "github.com/hahnicity/go-discern"
    "github.com/hahnicity/go-discern/data"
)

//XXX Make argument parser

type Worker struct {
    requests chan discern.WikiRequest
    index    int
}

type Pool []*Worker    

func (p Pool) Less() {

}

type Balancer struct {
    pool Pool    
}

func (b *Balancer) Balance(work chan discern.WikiRequest) {
    for {
        req := <-work // received a Request...
        b.dispatch(req) // ...so send it to a Worker
    }    
}

func (b *Balancer) dispatch(req discern.WikiRequest) {
    w := heap.Pop(b.pool).(*Worker)
    // ...send it the task.
    w.requests <- req
    // One more in its work queue.
    w.pending++
    // Put it into its place on the heap.
    heap.Push(&b.pool, w)
}


func Requester(year string, companies map[string]string, work chan <-discern.WikiRequest) {
    c := make(chan *discern.WikiResponse)
    for symbol, page := range companies {
        req := discern.WikiRequest{Resp: c, Symbol: symbol, Page: page, Year: year}
        work <- req
        resp <- req.Resp
        fmt.Println(resp)
    }
}

func main() {
    work := make(chan discern.WikiRequest)
    Requester("2012", data.SP500, work)
    for w := range work {
        fmt.Println(w)    
    }
}
