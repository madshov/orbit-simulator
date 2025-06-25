package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/madshov/data-structures/algebraic"
	"golang.org/x/image/colornames"

	"github.com/madshov/orbit-simulator/internal"
)

const (
	ScreenWidth  = 1000
	ScreenHeight = 800

	StarMass     = 6e15
	PlanetMass   = 6e9
	StarRadius   = 15
	PlanetRadius = 5
)

var (
	backgroundColor = colornames.Steelblue
)

type Simulation struct {
	simulator *internal.Simulator
}

func (s *Simulation) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (s *Simulation) Update() error {
	netForce := s.simulator.CalcForceVectors()
	s.simulator.UpdateBodies(netForce)

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	return nil
}

func (s *Simulation) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	s.simulator.DrawBodies(screen)
}

func NewSimulation(simulator *internal.Simulator) *Simulation {
	return &Simulation{
		simulator: simulator,
	}
}

func stableCircularOrbit() []*internal.Body {
	body1 := internal.NewBody(
		algebraic.NewVector(3, ScreenWidth/2, ScreenHeight/2, 0),
		StarMass,
		StarRadius,
		colornames.White,
	)

	body2 := internal.NewBody(
		algebraic.NewVector(3, 500, 600, 0),
		PlanetMass,
		PlanetRadius,
		colornames.Blue,
	)
	body2.AddToVelocity(algebraic.NewVector(3, 200, 0, 0))

	return []*internal.Body{body1, body2}
}

func stableEllipticalOrbit() []*internal.Body {
	body1 := internal.NewBody(
		algebraic.NewVector(3, ScreenWidth/2, ScreenHeight/2, 0),
		StarMass,
		StarRadius,
		colornames.White,
	)

	body2 := internal.NewBody(
		algebraic.NewVector(3, ScreenWidth/2, 550, 0),
		PlanetMass,
		PlanetRadius,
		colornames.Blue,
	)
	body2.AddToVelocity(algebraic.NewVector(3, 250, 100, 0))

	return []*internal.Body{body1, body2}
}

func erraticOrbit() []*internal.Body {
	body1 := internal.NewBody(
		algebraic.NewVector(3, 500, 300, 0),
		StarMass,
		10,
		colornames.White,
	)
	body1.AddToVelocity(algebraic.NewVector(3, 40, 0, 0))

	body2 := internal.NewBody(
		algebraic.NewVector(3, 600, 200, 0),
		StarMass,
		10,
		colornames.White,
	)
	body2.AddToVelocity(algebraic.NewVector(3, -40, 0, 0))

	body3 := internal.NewBody(
		algebraic.NewVector(3, 300, 500, 0),
		StarMass,
		10,
		colornames.White,
	)
	body3.AddToVelocity(algebraic.NewVector(3, 50, 0, 0))

	return []*internal.Body{body1, body2, body3}
}

func stableBinaryStarOrbit() []*internal.Body {
	body1 := internal.NewBody(
		algebraic.NewVector(3, 500, 350, 0),
		StarMass,
		10,
		colornames.White,
	)
	body1.AddToVelocity(algebraic.NewVector(3, -200, 0, 0))

	body2 := internal.NewBody(
		algebraic.NewVector(3, 500, 450, 0),
		StarMass,
		10,
		colornames.White,
	)
	body2.AddToVelocity(algebraic.NewVector(3, 200, 0, 0))

	return []*internal.Body{body1, body2}
}

func stableQuadrupleStarOrbit() []*internal.Body {
	body1 := internal.NewBody(
		algebraic.NewVector(3, 500, 300, 0),
		StarMass,
		10,
		colornames.White,
	)
	body1.AddToVelocity(algebraic.NewVector(3, -100, 0, 0))

	body2 := internal.NewBody(
		algebraic.NewVector(3, 400, 400, 0),
		StarMass,
		10,
		colornames.White,
	)
	body2.AddToVelocity(algebraic.NewVector(3, 0, 100, 0))

	body3 := internal.NewBody(
		algebraic.NewVector(3, 500, 500, 0),
		StarMass,
		10,
		colornames.White,
	)
	body3.AddToVelocity(algebraic.NewVector(3, 100, 0, 0))

	body4 := internal.NewBody(
		algebraic.NewVector(3, 600, 400, 0),
		StarMass,
		10,
		colornames.White,
	)
	body4.AddToVelocity(algebraic.NewVector(3, 0, -100, 0))

	return []*internal.Body{body1, body2}
}

func main() {
	fmt.Print(`Please enter desired orbit type:
1) Stable circular orbit
2) Stable elliptical orbit
3) Erratic orbit
4) Stable binary star orbit
5) Stable quadruple star orbit
`)

	var (
		inputScene int
		scene      []*internal.Body
	)

	fmt.Scanln(&inputScene)

	switch inputScene {
	case 5:
		scene = stableQuadrupleStarOrbit()
	case 4:
		scene = stableBinaryStarOrbit()
	case 3:
		scene = erraticOrbit()
	case 2:
		scene = stableEllipticalOrbit()
	default:
		scene = stableCircularOrbit()
	}

	simulator := internal.NewSimulator(scene)

	sim := NewSimulation(simulator)

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Orbit Simulation")
	if err := ebiten.RunGame(sim); err != nil {
		log.Fatal(err)
	}
}
