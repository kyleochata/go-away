package particles

import (
	"fmt"
	"math"
	"math/rand"
)

type Coffee struct {
	*ParticleSystem
}

func ascii(row, col int, counts [][]int) string {
	count := counts[row][col]
	if count < 1 {
		return "   "
	}
	if count < 6 {
		return ".  " // coffee bean
	}
	if count < 9 {
		return ":  "
	}
	return "}  "
}

// particle either out of bounds or life time is over
func reset(p *Particle, params *ParticleParams) {
	p.LifeTime = int64(math.Floor(float64(params.MaxLife) * rand.Float64()))
	p.Speed = params.MaxSpeed * rand.Float64()

	//normal distribution for the position of x.
	//rand.NormFloat64() returns a random number from a normal distribution with mean 0 and standard deviation 1
	// math.Max(-maxX, math.Min(maxX, rand.NormFloat64())) ensures that the x position is within the bounds of the screen;
	maxX := math.Floor(float64(params.X) / 2)
	x := math.Max(-maxX, math.Min(maxX, rand.NormFloat64()))
	p.X = x + maxX
	p.Y = 0
}

// raise the particle straight up
func nextPos(p *Particle, deltaMs int64) {
	p.LifeTime -= deltaMs
	if p.LifeTime <= 0 {
		return
	}
	p.Y += p.Speed * (float64(deltaMs) / 1000.0)
	// fmt.Printf("particle (%d %f) %f\n", p.LifeTime, p.Speed, p.Y)
}
func NewCoffee(width, height int) Coffee {
	if width%2 == 0 {
		tmp := width
		width++
		fmt.Sprintf("Width must be odd, got %d, changed to: %d", tmp, width)
	}
	return Coffee{
		NewParticleSystem(
			ParticleParams{
				MaxLife:       5000,
				MaxSpeed:      1,
				ParticleCount: 1000,
				X:             width,
				Y:             height,
				Reset:         reset,
				Ascii:         ascii,
				NextPosition:  nextPos,
			},
		),
	}
}
