package main

import "github.com/hahnicity/go-discern"

type Worker struct {
    requests chan *discern.WikiRequest
    pending  int  // count of pending tasks
}

func (w *Worker) handleRequest(done chan *Worker) {
    for  {
        req := <- w.requests
        req.GetYearlyStats()
        done <- w           // we've finished this request
    }    
}

func requester(work chan<- discern.WikiRequest) {
    c := make(chan int)
    for {
        work <- discern.WikiRequest{workFn, c}
        result := <-c                        
    }    
}

type Pool []*Worker

func (p Pool) Less(i, j int) bool {
    return p[i].pending < p[j].pending
}


type Balancer struct {
    pool Pool
    done chan *Worker
}

func (b *Balancer) balance(work chan discern.WikiRequest) {
    for {
        select {
        case req := <-work: // received a Request...
            b.dispatch(req) // ...so send it to a Worker
        case w := <-b.done: // a worker has finished ...
            b.completed(w)  // ...so update its info
        }
    }
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
    // Grab the least loaded worker...
    w := heap.Pop(&b.pool).(*Worker)
    // ...send it the task.
    w.requests <- req
    // One more in its work queue.
    w.pending++
    // Put it into its place on the heap.
    heap.Push(&b.pool, w)
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
    w.pending--
    // Remove it from heap.                  
    heap.Remove(&b.pool, w.index)
    // Put it into its place on the heap.
    heap.Push(&b.pool, w)
}
