package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type GUI struct {
	canvas *pixelgl.Window
}

const (
	window_width  = 750
	window_height = 750
)

var (
	window_cfg = pixelgl.WindowConfig{
		Title:     "Langton's Ant",
		Resizable: false,
		Bounds:    pixel.R(0, 0, float64(window_width), float64(window_height)),
	}

	bg_color   = color.RGBA{255, 255, 255, 255}
	cell_color = color.RGBA{37, 37, 37, 255}
	grid       = [window_width][window_height]bool{}
	gui        = GUI{}

	ant = Ant{
		color: color.RGBA{100, 0, 0, 255},
		speed: 800,
		pos:   pixel.V(window_width/2, window_height/2),
		dir:   UP,
	}
)

func main() {
	pixelgl.Run(setup)
}

func setup() {
	gui = GUI{
		canvas: create_window(window_cfg),
	}

	for !gui.canvas.Closed() {
		draw()
		update()

		gui.canvas.Update()
	}
}

func draw() {
	gui.canvas.Clear(bg_color)

	cells := imdraw.New(nil)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if !grid[i][j] {
				continue
			}

			x1 := float64(i)
			y1 := float64(j)
			x2 := float64(i + 1)
			y2 := float64(j + 1)

			cells.Color = cell_color
			cells.Push(pixel.V(x1, y1))
			cells.Push(pixel.V(x2, y2))
			cells.Rectangle(0)
		}
	}

	cells.Draw(gui.canvas)
	ant.draw()
}

func update() {
	ant.update()
}
