package main

import (
	"syscall/js"

	"github.com/ewaldhorn/dommie/dom"
)

var background js.Value

// ----------------------------------------------------------------------------
// load the image assets
func setupImages() {
	background = dom.CreateImg("background.jpg")
}
