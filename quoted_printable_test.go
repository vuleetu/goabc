package gob2a

import (
    "testing"
)

func TestQuote(t *testing.T){
    s := EncodeQuotedPrintable([]byte("你好"))
    if s == "=e4=bd=a0=e5=a5=bd" {
        return
    }

    t.Fail()
}
