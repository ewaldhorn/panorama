package main

import (
	"math"
	"syscall/js"
)

const (
	speed = 100 // lower is faster
	scale = 1.05
	y     = -4.5 // vertical offset
	dx    = -4.5
)

var (
	imgW   = 0.0
	imgH   = 0.0
	x      = 0.0
	clearX = 0.0
	clearY = 0.0
)

// ----------------------------------------------------------------------------
func setupAnimation() {
	// hardcoded, I know I know...
	imgW = 1299.0 * scale
	imgH = 700 * scale

	if imgW > canvasWidth {
		// Image larger than canvas
		x = float64(canvasWidth) - imgW
	}

	// Check if image dimension is larger than canvas
	clearX = math.Max(imgW, float64(canvasWidth))
	clearY = math.Max(imgH, float64(canvasHeight))

	js.Global().Call("setInterval", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		render()
		return nil
	}), speed)
}

// ----------------------------------------------------------------------------
func render() {
}
