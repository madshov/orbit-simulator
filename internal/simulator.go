package internal

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/madshov/data-structures/algebraic"
	"gonum.org/v1/gonum/unit/constant"
)

const (
	scalingFactor = 20
)

type Simulator struct {
	bodies []*Body
}

func NewSimulator(bodies []*Body) *Simulator {
	return &Simulator{
		bodies: bodies,
	}
}

func (s *Simulator) CalcForceVectors() []algebraic.Vector {
	numOfBodies := len(s.bodies)

	forces := make([][]algebraic.Vector, numOfBodies)
	dists := make([][]algebraic.Vector, numOfBodies)

	for i := 0; i < numOfBodies; i++ {
		for j := 0; j < numOfBodies; j++ {
			deltaDist := s.bodies[j].Position.Sub(s.bodies[i].Position)
			dists[i] = append(dists[i], deltaDist)
		}
	}

	for i := 0; i < numOfBodies; i++ {
		for j := 0; j < numOfBodies; j++ {
			if dists[i][j].Magnitude() != 0 {
				// F = G*((m1*m2)/r^2)
				force := float64(constant.Gravitational) *
					s.bodies[i].Mass *
					s.bodies[j].Mass *
					(1 / math.Pow(dists[i][j].Magnitude(), 2)) *
					scalingFactor

				nv := dists[i][j]
				nv.Normalize()
				nnv := nv.Scale(force)
				forces[i] = append(forces[i], nnv)
			} else {
				nv := algebraic.NewZeroVector(3)
				forces[i] = append(forces[i], nv)
			}
		}
	}

	var netForce []algebraic.Vector
	for _, f := range forces {
		f[0].Add(f[1])
		netForce = append(netForce, f[0])
	}

	return netForce
}

func (s *Simulator) DrawBodies(screen *ebiten.Image) {
	for _, body := range s.bodies {
		vector.DrawFilledCircle(
			screen,
			float32(body.Position.X()),
			float32(body.Position.Y()),
			float32(body.Radius),
			body.Color, true,
		)
	}
}

func (s *Simulator) UpdateBodies(netForce []algebraic.Vector) {
	for k, body := range s.bodies {
		vel := netForce[k].Scale(1 / body.Mass * timeDelay)
		body.AddToVelocity(vel)

		pos := body.Velocity.Scale(timeDelay)
		body.AddToPosition(pos)
	}
}
