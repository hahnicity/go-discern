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
func FindRecentDates(wr *WikiResponse, f float64) (dates map[string]int) {
    dates = FindDates(wr, f)
    now := time.Now().Unix()
    for date, _ := range dates {
        parsed, err := time.Parse(config.TimeConf, date)
        if err != nil {
            panic(err)    
        }
        ut := parsed.Unix()
        if now - ut > FIVE_DAYS_AGO {
            delete(dates, date)    
        }
    }
    return
}

// Find the dates where a given wiki page experienced abnormally high
// activity
func FindDates(wr *WikiResponse, viewPercentile float64) (dates map[string]int ) {
    vals := GetValuesFromMap(wr.Yearly)
    sort.Sort(vals)
    result := statistics.QuantileFromSortedData(vals, viewPercentile)
    dates = GetKeysGE(wr.Yearly, result)
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
func FindHighestMeans(means map[string]int, meanPercentile float64) (ret map[string]int) {
    vals := GetValuesFromMap(means)
    result := statistics.QuantileFromSortedData(vals, meanPercentile)
    means = GetKeysGE(means, result)
    return
}

// Convert a slice of floats to a Values struct
func FloatsToValues(floats []float64) (values *Values) {
     values = new(Values)
     for _, f := range floats {
        *values = append(*values, int(f))
    }
    return
}

// Get all keys mapped to their values in a map whose values are greater than or equal
// to some float
func GetKeysGE(Map map[string]int, val float64) map[string]int {
    for k, v := range Map {
        if v < int(val) {
            delete(Map, k)
        }
    }
    return Map
}

func GetValuesFromMap(Map map[string]int) *Values {
    v := &Values{}
    for _, j := range Map {
        *v  = append(*v, j)
    }
    return v
}

func JoinViewsMaps(map1 map[string]int, map2 map[string]int) map[string]int {
    for k, v := range map2 {
        map1[k] = v   
    }
    return map1
}
