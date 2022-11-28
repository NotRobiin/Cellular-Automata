package main

import (
	"github.com/faiface/pixel/pixelgl"
)

func create_window(window_settings pixelgl.WindowConfig) *pixelgl.Window {
	win, err := pixelgl.NewWindow(window_settings)

	if err != nil {
		panic(err)
	}

	return win
}
