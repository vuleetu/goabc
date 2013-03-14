goabc
=====

Ascii/Binary Conversion library for golang

**encoding supported currently**
* quoted printable

Install
===
go get github.com/vuleetu/goabc

Usage
===
import goabc first
```golang
import "github.com/vuleetu/goabc"
```

```golang
goabc.EncodeQuotedPrintable([]byte("你好"))
```
