package discern

import (
    "fmt"
    "github.com/hahnicity/go-stringit"
    "github.com/hahnicity/tweetlib"
    "net/http"
)

func GetToken(key, secret string) string {
    config := &tweetlib.Config{key, secret, ""}
    a := &tweetlib.ApplicationOnly{&http.Client{}, config}
    token, err := a.GetToken()
    if err != nil {
        panic(err)    
    }
    return token
}

func NewTwitterRequest(symbol, token string) *TwitterRequest {
    client, _ := tweetlib.NewApplicationClient(&http.Client{}, token)
    return &TwitterRequest{client, symbol}
}

type TwitterRequest struct {
    client *tweetlib.Client
    symbol string
}

func (tr *TwitterRequest) FindTweets() {
    opts := tweetlib.NewOptionals()
    opts.Add("lang", "en")
    opts.Add("result_type", "popular")
    r, err := tr.client.Search.Tweets(tr.symbol, opts)
    if err != nil {
        panic(err)
    }
    fmt.Println(stringit.Format("Analyzing tweets for {}", tr.symbol))
    for _, tweet := range r.Results {
        fmt.Println(
            stringit.Format(
                "\tTweet:\n\t\t{}\n\t\tRetweets:{}", tweet.Text, tweet.RetweetCount,
            ),
        )
    }
}
