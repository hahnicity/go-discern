package discern

import (
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

func (tr *TwitterRequest) Symbol() string {
    return tr.symbol
}

func (tr *TwitterRequest) Execute() {
    tr.client.Search.Tweets(tr.symbol, nil)
}
