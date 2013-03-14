package goabc

import (
    "bytes"
    "strings"
    "errors"
)

func EncodeQuotedPrintable(bin []byte) (string, error) {
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
            hex, err := ToHex(b/16)
            if err != nil {
                return "", err
            }
            buf.WriteByte(hex)
            hex, err = ToHex(b%16)
            if err != nil {
                return "", err
            }
            buf.WriteByte(hex)
            l += 3
        } else {
            buf.WriteByte(b)
            l++
        }
    }
    return string(buf.Bytes()), nil
}

func DecodeQuotedPrintable(s string) ([]byte, error) {
    var (
        l = 0
        buf bytes.Buffer
        a = strings.ToLower(s) //convert to lower first
    )

    for i := 0; i < len(a); {
        if l > 74 {
            return nil, errors.New("invalid encoding source")
        }

        if a[i] == '=' {
            if len(a) - i < 2 {
                return nil, errors.New("too few encoding source")
            }

            if a[i+1] == 13 && a[i+2] == 10 {
                l = 0
                i += 3
            } else {
                hh, err := FromHex(a[i+1])
                if err != nil {
                    return nil, err
                }

                lh, err := FromHex(a[i+2])
                if err != nil {
                    return nil, err
                }
                buf.WriteByte(hh*16 + lh)
                i += 3
                l += 3
            }
        } else {
            if err := buf.WriteByte(a[i]); err != nil {
                return nil, err
            }
            l ++
            i ++
        }
    }

    return buf.Bytes(), nil
}

func ToHex(b byte) (byte, error) {
    if b < 10 {
        return b + '0', nil
    } else if b < 16 {
        return b - 10 + 'a', nil
    }

    return 0, errors.New("invalid number")
}

func FromHex(h byte) (byte, error) {
    if h >= 'a' && h <= 'f' {
        return h - 'a' + 10, nil
    } else if h >= '0' {
        return h - '0', nil
    }
    return 0, errors.New("invalid hex string")
}
