package discern

func FloatsToValues(floats []float64) (values *Values) {
     values = new(Values)
     for _, f := range floats {
        *values = append(*values, int(f))
    }
    return
}

type Values []int

func (v Values) Value(i int) float64 {
    return float64(v[i])
}
 
func (v *Values) SetValue(i int, val float64) {
    (*v)[i] = int(val)    
}

func (v Values) Len() int {
    return len(v)    
}

func (v Values) Less(i, j int) bool {
    return v[i] < v[j]  
}

func (v *Values) Swap(i, j int) {
    (*v)[i], (*v)[j] = (*v)[j], (*v)[i]     
}

type Dates []string

func GetValuesFromMap(Map map[string]int) *Values {
    v := &Values{}
    for _, j := range Map {
        *v  = append(*v, j)
    }
    return v
}

func GetDatesGE(Map map[string]int, val float64) *Dates {
    dates := &Dates{}
    for k, v := range Map {
        if v >= int(val) {
            (*dates) = append((*dates), k)
        }
    }
    return dates
}

func JoinViewsMaps(map1 map[string]int, map2 map[string]int) map[string]int {
    for k, v := range map2 {
        map1[k] = v   
    }
    return map1
}
