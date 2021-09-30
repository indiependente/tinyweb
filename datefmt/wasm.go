package main

import (
	"syscall/js"
	"time"
)

const (
	htmlDateLayout = "2006-01-02"
)

func main() {
}

//export format
func format(d string) string {
	return time.Now().UTC().Format(time.RFC3339)
}

//export update
func update() {
	document := js.Global().Get("document")
	dd := document.Call("getElementById", "dd").Get("value").String()
	result := format(dd)
	document.Call("getElementById", "result").Set("value", result)
}
