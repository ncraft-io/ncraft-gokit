package filter

import (
    "github.com/alecthomas/assert"
    "net/url"
    "testing"
)

func TestHttpQueryFilter_Compile(t *testing.T) {
    query := url.Values{
        "foo":    []string{`"bar"`},
        "far":    []string{"in(a, b, c)"},
        "name":   []string{`"*bb"`},
        "count":  []string{"1..<6"},
        "filter": []string{"a > 0 && a < 100"},
    }

    filter := HttpQueryFilter{}.Compile(query)
    assert.NotEmpty(t, filter)
}
