package main

import "math"

type Point struct {
	X, Y float64
}

type Point3 struct {
	X, Y, Z float64
}

type Polar struct {
	R, T float64
}

func Abs(p *Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func Scale(p *Point, s float64) (q Point) {
	q.X = p.X * s
	q.Y = q.Y * s
	return
}
