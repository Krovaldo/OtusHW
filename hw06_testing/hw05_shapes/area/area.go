package area

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

type Rectangle struct {
	Length float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Length * r.Height
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t *Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func CalculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, fmt.Errorf("не реализует интерфейс \"Shape\"")
	}
	return shape.Area(), nil
}
