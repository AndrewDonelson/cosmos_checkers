//go:build js && wasm

package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func jsSKDVersion(this js.Value, args []js.Value) interface{} {
	return SDKVersion("Cosmos-Checkers-SDK")
}

func jsConsumeAPINodeInfo(this js.Value, args []js.Value) interface{} {
	return consumeAPI_NodeInfo()
}

func jsSquare(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Invalid no of args passed"
	}

	input, err := strconv.ParseFloat(args[0].String(), 64)
	fmt.Printf("Input is [%v]", input)
	if err != nil {
		return err.Error()
	}
	square := square(input)
	fmt.Printf("Square of [%v] is [%v]", input, square)
	return square
}

func jsGetTestData(this js.Value, args []js.Value) interface{} {
	if len(args) != 0 {
		return "No args expected"
	}
	testData, err := getTestData()
	if err != nil {
		return err.Error()
	}
	return testData
}
