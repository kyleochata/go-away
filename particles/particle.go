package particles

import (
	"math"
	"slices"
	"strings"
	"time"
)

type Particle struct {
	X float64
	Y float64

	LifeTime int64
	Speed    float64
}

type ParticleParams struct {
	MaxLife  int64
	MaxSpeed float64

	ParticleCount int

	X int
	Y int

	NextPosition NextPositionFunc
	Ascii        Ascii
	Reset        Reset
}

type NextPositionFunc func(p *Particle, deltaMs int64)
type Ascii func(row, col int, count [][]int) string
type Reset func(particle *Particle, params *ParticleParams)

type ParticleSystem struct {
	ParticleParams
	lastTime  int64
	particles []*Particle
}

func NewParticleSystem(params ParticleParams) *ParticleSystem {
	particles := make([]*Particle, 0)
	for i := 0; i < params.ParticleCount; i++ {
		particles = append(particles, &Particle{})
	}
	return &ParticleSystem{
		ParticleParams: params,
		lastTime:       time.Now().UnixMilli(),
		particles:      particles,
	}
}
func (ps *ParticleSystem) Start() {
	for _, p := range ps.particles {
		ps.Reset(p, &ps.ParticleParams)
	}
}

func (ps *ParticleSystem) Update() {
	now := time.Now().UnixMilli()
	delta := now - ps.lastTime
	ps.lastTime = now

	for _, p := range ps.particles {
		ps.NextPosition(p, delta)
		if p.Y >= float64(ps.Y) || p.X >= float64(ps.X) || p.LifeTime <= 0 {
			ps.Reset(p, &ps.ParticleParams)

		}
	}

}

func (ps *ParticleSystem) Display() string {
	counts := make([][]int, 0)
	for row := 0; row < ps.Y; row++ {
		count := make([]int, 0)
		for col := 0; col < ps.X; col++ {
			count = append(count, 0)
		}
		counts = append(counts, count)
	}
	for _, p := range ps.particles {
		row := int(math.Floor(p.Y)) //Round gives out of range
		col := int(math.Round(p.X))
		if row >= 0 && row < ps.Y && col >= 0 && col < ps.X {
			counts[row][col]++
		}
	}
	// fmt.Printf("particles: %v\n", counts)
	out := make([][]string, 0)
	for r, row := range counts {
		outRow := make([]string, 0)
		for c := range row {
			outRow = append(outRow, ps.Ascii(r, c, counts))
		}
		out = append(out, outRow)
	}
	slices.Reverse(out)
	outStr := make([]string, 0)
	for _, row := range out {
		outStr = append(outStr, strings.Join(row, ""))
	}
	return strings.Join(outStr, "\n")
}
