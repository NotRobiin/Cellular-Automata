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
	for i := 0; i < ant.speed; i++ {
		i := int(ant.pos.X)
		j := int(ant.pos.Y)

		ant.turn(!grid[i][j])

		// Flip the cell
		grid[i][j] = !grid[i][j]

		// Update ant's position
		ant.pos.X += ant.dir.X
		ant.pos.Y += ant.dir.Y

		ant.handle_edges()
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
		ant.dir = cw[ant.dir]
	} else {
		ant.dir = ccw[ant.dir]
	}
}

func (a *Ant) handle_edges() {
	if ant.pos.X == window_width {
		ant.pos.X = 0
	} else if ant.pos.X < 0 {
		ant.pos.X = window_width - 1
	}

	if ant.pos.Y == window_height {
		ant.pos.Y = 0
	} else if ant.pos.Y < 0 {
		ant.pos.Y = window_height - 1
	}
}
