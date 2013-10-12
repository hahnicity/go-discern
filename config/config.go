package config

var NumberMonths int    = 12
var TimeConf     string = "2006-01-02"
var WikiUrl      string = "http://stats.grok.se/json/en"


type Config struct {
    CloseReq       bool
    MeanPercentile float64
    Processes      int
    ViewPercentile float64
    Workers        int
    Year           string
}
