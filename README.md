# must

Wraps a call to a function returning (T, error)
and panics if the error is non-nil.

## Installation

```sh
go get github.com/g-must/must
```

## Usage

```go
package main

import (
	"net/http"

	. "github.com/g-must/must"
)

func main() {
	_ = Must(http.Get(""))
}
```
