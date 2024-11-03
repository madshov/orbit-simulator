package main

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/madshov/data-structures/algebraic"
	"golang.org/x/image/colornames"

	"github.com/madshov/orbit-simulator/internal"
)

const (
	ScreenWidth  = 1000
	ScreenHeight = 800
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

func main() {

	vec1, _ := algebraic.NewVector(3, ScreenWidth/2, ScreenHeight/2, 0)
	body1 := internal.NewBody(vec1, 6*math.Pow(10, 15), 10, colornames.White)

	vec2, _ := algebraic.NewVector(3, 500, 595, 0)
	body2 := internal.NewBody(vec2, 6*math.Pow(10, 9), 10, colornames.Blue)

	vel1, _ := algebraic.NewVector(3, 200, 0, 0)
	body2.AddToVelocity(vel1)

	simulator := internal.NewSimulator([]*internal.Body{body1, body2})

	sim := NewSimulation(simulator)

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Orbit Simulation")
	if err := ebiten.RunGame(sim); err != nil {
		log.Fatal(err)
	}
}
