package discern

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
