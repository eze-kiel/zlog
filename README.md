# zlog

[![Go Reference](https://pkg.go.dev/badge/github.com/eze-kiel/zlog.svg)](https://pkg.go.dev/github.com/eze-kiel/zlog)

A small logging package, with cool features.

## Getting started

```
$ go get github.com/eze-kiel/zlog
```

## Usage

Here is a simple example about the main features of the package:

```go
package main

import (
	"fmt"
	"log"

	"github.com/eze-kiel/zlog"
)

func main() {
	l := zlog.NewLogger() // create logger
	l.Debug("message at debug level")
	l.Info("blah at info level")

    // change log level to warn
	if err := l.ParseLevel("warn"); err != nil {
		log.Fatal(err)
	}

	l.Info("you cannot see me!")
	l.Warn("but you can see me!!")
	l.Error("oops, an error !")

    // disable the colors
	l.UseColors(false)
	l.Fatal("this one is going to be fatal... :(")
}
```

## License

MIT