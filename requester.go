package discern

import "net/http"


func Requester(year string, companies map[string]string, n int, work chan <-WikiRequest) {
    c := make(chan *WikiResponse)
    tr := &http.Transport{DisableKeepAlives: true}
    client := &http.Client{Transport: tr}
    items := 0
    for symbol, page := range companies {
        items++
        req := WikiRequest{Client: client, Resp: c, Page: page, Symbol: symbol, Year: year}
        work <- req
        if items > n { // Seems like the number of outbound network connections is limiting me here
            <-c
            tr.CloseIdleConnections()
            items--
        }
    }
}
