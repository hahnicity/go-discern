package discern


func JoinViewsMaps(map1 map[string]int, map2 map[string]int) map[string]int {
    for k, v := range map2 {
        map1[k] = v   
    }
    return map1
}
