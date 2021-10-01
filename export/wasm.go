package main

import (
	"errors"
	"strconv"
	"syscall/js"
)

//export add

var (
	ErrInvalidInput = errors.New("negative input not valid")
)

func main() {
}

// Copyright (c) 2015 Project Nayuki
// All rights reserved. Contact Nayuki for licensing.
// https://www.nayuki.io/page/fast-fibonacci-algorithms
//export fibonacci
func fibonacci(n int64) (uint64, error) {
	if n < 0 {
		return 0, ErrInvalidInput
	}
	fibN, _ := fastFibonacci(uint64(n))
	return fibN, nil
}

func fastFibonacci(n uint64) (uint64, uint64) {
	if n == 0 {
		return 0, 1
	}
	a, b := fastFibonacci(n / 2)
	c := a * (b*2 - a)
	d := a*a + b*b
	if n%2 == 0 {
		return c, d
	}
	return d, c + d
}

//export update
func update() {
	document := js.Global().Get("document")
	nStr := document.Call("getElementById", "n").Get("value").String()
	n, _ := strconv.Atoi(nStr)
	result, err := fibonacci(int64(n))
	if err != nil {
		document.Call("getElementById", "result").Set("value", err.Error())
		return
	}
	document.Call("getElementById", "result").Set("value", result)
}
