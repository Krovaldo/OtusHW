package main

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

func calculateArea(s any) (float64, error) {
	shape, ok := s.(Shape)
	if !ok {
		return 0, fmt.Errorf("не реализует интерфейс \"Shape\"")
	}
	return shape.Area(), nil
}

func printArea(s any) {
	area, err := calculateArea(s)
	if err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		switch shape := s.(type) {
		case *Circle:
			fmt.Printf("Круг: радиус %v\n", shape.Radius)
		case *Rectangle:
			fmt.Printf("Прямоугольник: ширина %v, высота %v\n", shape.Length, shape.Height)
		case *Triangle:
			fmt.Printf("Треугольник: основание %v, высота %v\n", shape.Base, shape.Height)
		}
		fmt.Println("Площадь:", area)
	}
}

func main() {
	c := &Circle{Radius: 7}
	r := &Rectangle{Length: 5, Height: 6}
	t := &Triangle{Base: 8, Height: 5}
	n := "не фигура"

	printArea(c)
	printArea(r)
	printArea(t)
	printArea(n)
}
