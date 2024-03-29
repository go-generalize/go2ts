# go2ts
[![PkgGoDev](https://pkg.go.dev/badge/go-generalize/go2ts)](https://pkg.go.dev/github.com/go-generalize/go2ts)

## What is go2ts?
- go2ts is a TypeScript interface generator from Go struct.
- Automatically recognize the Go module in the directory

## Installation
```console
$ go get github.com/go-generalize/go2ts
```

## Usage

```go
// ./example/main.go
package main

import (
    "time"
)

type Status string

const (
    StatusOK Status = "OK"
    StatusFailure Status = "Failure"
)

type Param struct {
    Status    Status
    Version   int
    Action    string
    CreatedAt time.Time
}
```

```console
$ go2ts ./example
export type Param = {
        Action: string;
        CreatedAt: string;
        Status: "Failure" | "OK";
        Version: number;
}
```

## Known Issues
- [#18](https://github.com/go-generalize/go2ts/issues/18): enums in indirectly parsed packages are not recognized.

## TODO
- Handle MarshalJSON/UnmarshalJSON
