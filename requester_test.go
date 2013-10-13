package discern

import "testing"


func TestAnalyzeResponse(t *testing.T) {
    resp := &WikiResponse{"YCK", map[string]int{"10-10-2012": 100}}    
    analyzeResponse(resp, 99.9)
}
