// https://tinygo.org/docs/reference/lang-support/stdlib/

package main

import (
	"os/exec"
	"syscall/js"
)

func main() {
}

//export format
func run(c string) string {
	cmd := exec.Command("go", "version")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	return string(out)
}

//export update
func update() {
	document := js.Global().Get("document")
	dd := document.Call("getElementById", "code").Get("value").String()
	result := run(dd)
	document.Call("getElementById", "result").Set("value", result)
}
