package discern


func Requester(year string, companies map[string]string, n int, work chan <-WikiRequest) {
    c := make(chan *WikiResponse)
    items := 0
    for symbol, page := range companies {
        items++
        req := makeWikiRequest(year, page, symbol)
        work <- req
        if items > n { // Seems like the number of outbound network connections is limiting me here
            <-c
            items--
        }
    }
}
