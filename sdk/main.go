//go:build js && wasm

package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"syscall/js"
)

func main() {
	// This is a channel that we will use to block the main thread
	c := make(chan bool)

	// Register the functions to be called from JS
	js.Global().Set("SDKVersion", js.FuncOf(jsSKDVersion))
	js.Global().Set("NodeInfo", js.FuncOf(jsConsumeAPINodeInfo))
	// Test Functions
	js.Global().Set("square", js.FuncOf(jsSquare))
	js.Global().Set("getTestData", js.FuncOf(jsGetTestData))

	fmt.Println(SDKVersion("Cosmos-Checkers-SDK"), "ready")

	// Block the main thread forever
	<-c

	// This is a hack to prevent the program from exiting
	fmt.Println(SDKVersion("Cosmos-Checkers-SDK"), "exited")
}

/* getTestData() is a function that makes a GET request to a REST API and returns the response body as a string.
 * This function is called from the JS code.
 */
func getTestData() (string, error) {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return "", err
	}

	// We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert the body to type string
	sb := string(body)
	return sb, nil
}

/* square() is a function that takes a float64 as input and returns the square of the input.
 * This function is called from the JS code.
 */
func square(input float64) float64 {
	return math.Pow(input, 2)
}
