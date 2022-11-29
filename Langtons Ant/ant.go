package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

type Ant struct {
	color color.Color
	speed int
	pos   pixel.Vec
	dir   pixel.Vec
}

var (
	UP    = pixel.V(0, -1)
	DOWN  = pixel.V(0, 1)
	LEFT  = pixel.V(-1, 0)
	RIGHT = pixel.V(1, 0)
)

func (a *Ant) update() {
	for i := 0; i < a.speed; i++ {
		i := int(a.pos.X)
		j := int(a.pos.Y)

		a.turn(!grid[i][j])

		// Flip the cell
		grid[i][j] = !grid[i][j]

		// Update ant's position
		a.pos.X += a.dir.X
		a.pos.Y += a.dir.Y

		a.handle_edges()
	}
}

func (a *Ant) draw() {
	img := imdraw.New(nil)
	img.Color = a.color
	img.Push(a.pos)
	img.Push(pixel.V(a.pos.X+1, a.pos.Y+1))
	img.Rectangle(0)
	img.Draw(gui.canvas)
}

func (a *Ant) turn(clockwise bool) {
	ccw := map[pixel.Vec]pixel.Vec{
		UP:    LEFT,
		DOWN:  RIGHT,
		LEFT:  DOWN,
		RIGHT: UP,
	}

	cw := map[pixel.Vec]pixel.Vec{
		UP:    RIGHT,
		DOWN:  LEFT,
		LEFT:  UP,
		RIGHT: DOWN,
	}

	if clockwise {
		a.dir = cw[a.dir]
	} else {
		a.dir = ccw[a.dir]
	}
}

func (a *Ant) handle_edges() {
	if a.pos.X == window_width {
		a.pos.X = 0
	} else if a.pos.X < 0 {
		a.pos.X = window_width - 1
	}

	if a.pos.Y == window_height {
		a.pos.Y = 0
	} else if a.pos.Y < 0 {
		a.pos.Y = window_height - 1
	}
}
