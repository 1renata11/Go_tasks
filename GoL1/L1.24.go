package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(X float64, Y float64) *Point {
	return &Point{
		x: X,
		y: Y,
	}
}

func (p Point) getX() float64 {
	return p.x
}

func (p Point) getY() float64 {
	return p.y
}

func countLen(p1 Point, p2 Point) float64 {
	return math.Sqrt((p1.getX()-p2.getX())*(p1.getX()-p2.getX()) + (p1.getY()-p2.getY())*(p1.getY()-p2.getY()))
}

func L124() {
	p1 := NewPoint(1.2, 4.5)
	p2 := NewPoint(5.9, 2.1)
	fmt.Print(countLen(*p1, *p2))
}
