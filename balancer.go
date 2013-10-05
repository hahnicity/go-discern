package discern

import "container/heap"
import "fmt"


type Worker struct {
    requests chan WikiRequest
    index    int
}

func (w *Worker) work() {
    go func() {
        req := <- w.requests
        req.Resp <- req.GetYearlyStats()
    }()
}

type Pool []*Worker

func (p Pool) Len() int { 
    return len(p) 
}

func (p Pool) Less(i, j int) bool {
    return p[i].index < p[j].index
}

func (p Pool) Swap(i, j int) { 
    p[i], p[j] = p[j], p[i] 
}

func (p *Pool) Push(x interface{}) {
    x.(*Worker).index = p.Len()
    *p = append(*p, x.(*Worker))    
}

func (p *Pool) Pop() interface{} {
    old := *p
    n := len(old)
    x := old[n-1]
    *p = old[0 : n-1]
    return x    
}

type Balancer struct {
    pool *Pool    
}

func (b *Balancer) Balance(work chan WikiRequest) {
    for {
        req := <-work // received a Request...
        b.dispatch(req) // ...so send it to a Worker
    }    
}

func (b *Balancer) dispatch(req WikiRequest) {
    w := heap.Pop(b.pool).(*Worker)
    fmt.Println("USING WORKER", w.index)
    // ...send it the task.
    w.work()
    w.requests <- req
}

func MakeBalancer(n int) *Balancer {
    b := &Balancer{makePool(n)}
    heap.Init(b.pool) //initialize the pool
    return b
}

func makePool(n int) *Pool {
    p := new(Pool)
    for i := 0; i < n; i++ {
        requests := make(chan WikiRequest)
        p.Push(&Worker{requests, i})
    }
    return p
}
