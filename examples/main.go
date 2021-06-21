package main

import (
	"fmt"
	"log"

	"github.com/eze-kiel/zlog"
)

func main() {
	fmt.Println("Hello world")
	l := zlog.NewLogger()
	l.Debug("message at debug level")
	l.Info("blah at info level")

	if err := l.ParseLevel("warn"); err != nil {
		log.Fatal(err)
	}

	l.Info("you cannot see me!")
	l.Warn("but you can see me!!")
	l.Error("oops, an error !")

	l.UseColors(false)
	l.Fatal("this one is going to be fatal... :(")
}
