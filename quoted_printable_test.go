package goabc

import (
    "testing"
)

const TEST_STRING = "你好"

func TestQuote(t *testing.T){
    s, err := EncodeQuotedPrintable([]byte(TEST_STRING))
    if err != nil || s != "=e4=bd=a0=e5=a5=bd" {
        t.Fail()
        return
    }

    b, err := DecodeQuotedPrintable(s)
    if err != nil || string(b) != TEST_STRING {
        t.Fail()
        return
    }
}
