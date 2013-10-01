package main

import (
    "fmt"
    "github.com/hahnicity/go-discern/wiki"
)


func main() {
    fmt.Println(wiki.GetYearlyStats("2012", "Apple_inc"))
}
