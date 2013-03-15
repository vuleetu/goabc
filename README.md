#goabc

Ascii/Binary Conversion library for golang

**encoding supported currently**
* quoted printable

#Install

go get github.com/vuleetu/goabc

#Usage

import goabc first
```Go
import "github.com/vuleetu/goabc"
```

```Go
goabc.EncodeQuotedPrintable([]byte("你好"))
```
