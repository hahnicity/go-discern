package discern

import (
    "github.com/grd/statistics"
    "sort"
)


func FindDates(wr *WikiResponse, f float64) *Dates {
    vals := GetValuesFromMap(wr.Yearly)
    sort.Sort(vals)
    result := statistics.QuantileFromSortedData(vals, f)
    return GetDatesGE(wr.Yearly, result)
}
