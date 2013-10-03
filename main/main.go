package main

import "github.com/hahnicity/go-discern/data"


// Make argument parser


func main() {
    for symbol, page := range data.SP500 {
        HandleWikiCompany(page)
    }
}
