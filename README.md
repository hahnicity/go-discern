go-discern
==========

Golang translation of discerner. Objective here is to track public interest in a 
meaningful way, not tie it too heavily with financials like in the python repo.

## Installation
Install the following 4 packages. I need to add in a dependency manager

    go get github.com/hahnicity/go-stringit
    go get github.com/hahnicity/tweetlib
    go get github.com/grd/statistics
    go get github.com/go-discern 

## Function
Looks for a S&P 500 company's page views on wikipedia. If they are above some
threshold (a quantile) then `go-discern` will flag that date to be of interest
to the viewer. Afterwards we have the options of looking for tweets in relation to
companies of interest

## Configuration
All configuration of `go-discern` should be taken care of automatically, but if
you ever want to get information from twitter you will need to 

 * Register a new Twitter Application
 * Place the application's key and secret in some kind of file accessible by the `config` package
 * Key should be defined as `TwitterKey`
 * Secret should be defined as `TwitterSecret`

Example:

Create a new file name `twitterconfig.go` in the `config` package

    package config
        
    var (
        TwitterKey    string = "hsfkhklj"
        TwitterSecret string = "aklhjfkhafkjh"
    )
        
Unfortunately the ability to specify the key and secret is currently not supported 
over the command line

## Usage
All usage of this package can be done through the command line.

    go run main/main.go

Of course we have the option of specifying different command line parameters as well

    -max => The number of concurrent requests we can make to wikipedia
    -workers => The number of workers we want to create in our heap
    -year => The year we want to analyze stats for
    -viewFunc => (FindDates or FindRecentDates). FindDates will return all dates with
                 higher than normal activity. FindRecentDates will find dates within
                 past 5 days
    -viewQuant => All dates with views inside this quantile will be shown
    -closeRequests => Close all requests made to wikipedia immediately after getting a
                      response
    -analyzeMeans => Show companies with the mean views within a quantile
    -meanQuant => The quantile to display mean views for
    -analyzeTweets => Set to false if you do not want to analyze tweets
