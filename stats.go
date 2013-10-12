package discern

import (
    "github.com/grd/statistics"
    "github.com/hahnicity/go-discern/config"
    "sort"
    "time"
)

const FIVE_DAYS_AGO int64 = 60 * 60 * 24 * 5

// Find all companies with abnormally high activity within the last 
// 5 days. 
func FindRecentDates(wr *WikiResponse, f float64) (dates *Dates) {
    dates = FindDates(wr, f)
    now := time.Now().Unix()
    for i, date := range *dates {
        parsed, err := time.Parse(config.TimeConf, date)
        if err != nil {
            panic(err)    
        }
        ut := parsed.Unix()
        if now - ut > FIVE_DAYS_AGO {
            *dates = append((*dates)[:i], (*dates)[i+1:]...)    
        }
    }
    return
}

// Find the dates where a given wiki page experienced abnormally high
// activity
func FindDates(wr *WikiResponse, f float64) (dates *Dates) {
    vals := GetValuesFromMap(wr.Yearly)
    sort.Sort(vals)
    result := statistics.QuantileFromSortedData(vals, f)
    dates = GetDatesGE(wr.Yearly, result)
    return
}

// Find the mean number of views for a wikipedia page
func FindMeanViews(wr *WikiResponse) (mean int) {
    vals := GetValuesFromMap(wr.Yearly)
    mean = int(statistics.Mean(vals))
    return
}

// Find the companies with the highest number of mean views according to a
// mean view percentile
func FindHighestMeans(means map[string]int, meanPercentile float64) map[string]int {
    
}
