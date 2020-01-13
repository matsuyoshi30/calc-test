package main

import (
	"math/rand"
	"strconv"
	"syscall/js"
)

// input and clear number

var input = ""

func inputNum(this js.Value, i []js.Value) interface{} {
	if input != "0" {
		input += i[0].String()
		setContent("inputnum", input)
	}
	return nil
}

func clearNum(this js.Value, i []js.Value) interface{} {
	reset()
	return nil
}

// check answer

var started = false

func submit(this js.Value, i []js.Value) interface{} {
	if !started {
		generate()
		reset()

		started = true
	} else {
		inputnum := getNum("inputnum")
		n, _ := strconv.Atoi(inputnum)
		lval := getNum("lval")
		l, _ := strconv.Atoi(lval)
		rval := getNum("rval")
		r, _ := strconv.Atoi(rval)

		if l+r != n {
			setContent("result", "WRONG")
		} else {
			setContent("result", "CORRECT")
		}

		generate()
		reset()
	}

	return nil
}

func setContent(id, str string) {
	js.Global().Get("document").Call("getElementById", id).Set("textContent", str)
}

func getNum(id string) string {
	return js.Global().Get("document").Call("getElementById", id).Get("textContent").String()
}

func generate() {
	lval := rand.Intn(100)
	rval := rand.Intn(100)

	setContent("lval", strconv.Itoa(lval))
	setContent("rval", strconv.Itoa(rval))
}

func reset() {
	input = ""
	setContent("inputnum", "0")
}

func registerCallbacks() {
	js.Global().Set("inputNum", js.FuncOf(inputNum))
	js.Global().Set("clearNum", js.FuncOf(clearNum))
	js.Global().Set("submit", js.FuncOf(submit))
}

var c chan struct{}

func init() {
	registerCallbacks()
	c = make(chan struct{}, 0)
}

func main() {
	<-c
}
