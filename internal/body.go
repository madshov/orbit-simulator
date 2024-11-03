package internal

import (
	"image/color"

	"github.com/madshov/data-structures/algebraic"
)

const (
	timeDelay = 0.01
)

type Body struct {
	Position *algebraic.Vector
	Velocity *algebraic.Vector
	Mass     float64
	Radius   float64

	Color color.Color
}

func NewBody(pos *algebraic.Vector, mass, radius float64, color color.Color) *Body {
	vel, _ := algebraic.NewZeroVector(3)

	return &Body{
		Position: pos,
		Velocity: vel,
		Mass:     mass,
		Radius:   radius,
		Color:    color,
	}
}

func (b *Body) AddToVelocity(vel *algebraic.Vector) {
	b.Velocity = b.Velocity.Add(vel)
}

func (b *Body) AddToPosition(pos *algebraic.Vector) {
	b.Position = b.Position.Add(pos)
}
