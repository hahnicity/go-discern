package discern

type Request interface {
    SetResponse(Response)
    GetResponse() Response
    Symbol()      string
    Execute()     interface{}
}

type Response interface {
    Symbol() string
    Data()   interface{}
}

func Requester(year string, companies map[string]string, n int, work chan <-WikiRequest) {
    items := 0
    for symbol, page := range companies {
        items++
        req := makeWikiRequest(year, page, symbol)
        work <- req
        if items > n { // Seems like the number of outbound network connections is limiting
            <-req.Resp
            items--
        }
    }
}
