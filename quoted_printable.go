package gob2a

import (
    "bytes"
)

func EncodeQuotedPrintable(bin []byte) string {
    var (
        l = 0
        buf bytes.Buffer
    )

    for _, b := range bin {
        if l >= 72 {
            buf.WriteByte('=')
            buf.WriteByte(13)
            buf.WriteByte(10)
            l = 0
        }

        if b < 32 || b >= 127 || b == '=' || b == '.' {
            buf.WriteByte('=')
            buf.WriteByte(ToHex(b/16))
            buf.WriteByte(ToHex(b%16))
            l += 3
        } else {
            buf.WriteByte(b)
            l++
        }
    }
    return string(buf.Bytes())
}

func ToHex(b byte) byte {
    if b < 10 {
        return b + '0'
    }
    return b - 10 + 'a'
}
