package main

import (
	"github.com/kataras/golog"
)

func main() {
	golog.SetLevel("debug")
	golog.Debugf("Hello world!")
}
