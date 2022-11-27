package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type Ant struct {
	color color.Color
	speed int
	pos   pixel.Vec
	dir   pixel.Vec
}

const (
	window_width  = 750
	window_height = 750
)

var (
	UP    = pixel.V(0, -1)
	DOWN  = pixel.V(0, 1)
	LEFT  = pixel.V(-1, 0)
	RIGHT = pixel.V(1, 0)

	window_cfg = pixelgl.WindowConfig{
		Title:     "Langton's Ant",
		Resizable: false,
		Bounds:    pixel.R(0, 0, float64(window_width), float64(window_height)),
	}

	bg_color   = color.RGBA{255, 255, 255, 255}
	cell_color = color.RGBA{37, 37, 37, 255}
	grid       = [window_width][window_height]bool{}

	ant = Ant{
		color: color.RGBA{100, 0, 0, 255},
		speed: 500,
		pos:   pixel.V(window_width/2, window_height/2),
		dir:   UP,
	}
)

func main() {
	pixelgl.Run(setup)
}

func setup() {
	window, err := pixelgl.NewWindow(window_cfg)

	if err != nil {
		panic(err)
	}

	for !window.Closed() {
		draw(window)
		update(window)

		window.Update()
	}
}

func draw(window *pixelgl.Window) {
	window.Clear(bg_color)

	draw_cells(window)
	draw_ant(window)
}

func update(window *pixelgl.Window) {
	for i := 0; i < ant.speed; i++ {
		i := int(ant.pos.X)
		j := int(ant.pos.Y)

		// Spot taken (black)
		if grid[i][j] {
			turn_counter_clockwise()
		} else {
			turn_clockwise()
		}

		// Flip the cell
		grid[i][j] = !grid[i][j]

		// Update ant's position
		ant.pos.X += ant.dir.X
		ant.pos.Y += ant.dir.Y

		handle_x_edge()
		handle_y_edge()
	}
}

func draw_cells(window *pixelgl.Window) {
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

	cells.Draw(window)
}

func draw_ant(window *pixelgl.Window) {
	img := imdraw.New(nil)
	img.Color = ant.color
	img.Push(ant.pos)
	img.Push(pixel.V(ant.pos.X+1, ant.pos.Y+1))
	img.Rectangle(0)
	img.Draw(window)
}

func turn_clockwise() {
	switch ant.dir {
	case UP:
		ant.dir = RIGHT
	case DOWN:
		ant.dir = LEFT
	case LEFT:
		ant.dir = UP
	case RIGHT:
		ant.dir = DOWN
	}
}

func turn_counter_clockwise() {
	switch ant.dir {
	case UP:
		ant.dir = LEFT
	case DOWN:
		ant.dir = RIGHT
	case LEFT:
		ant.dir = DOWN
	case RIGHT:
		ant.dir = UP
	}
}

func handle_x_edge() {
	// X-edge
	if ant.pos.X == window_width {
		ant.pos.X = 0
	}

	if ant.pos.X < 0 {
		ant.pos.X = window_width - 1
	}
}

func handle_y_edge() {
	// Y-edge
	if ant.pos.Y == window_height {
		ant.pos.Y = 0
	}

	if ant.pos.Y < 0 {
		ant.pos.Y = window_height - 1
	}
}
