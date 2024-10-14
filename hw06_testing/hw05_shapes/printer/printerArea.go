package printerArea

import (
	"fmt"

	"github.com/Krovaldo/OtusHW/hw06_testing/hw05_shapes/area"
)

func PrintArea(s any) {
	areaVal, err := area.CalculateArea(s)
	if err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		switch shape := s.(type) {
		case *area.Circle:
			fmt.Printf("Круг: радиус %v\n", shape.Radius)
		case *area.Rectangle:
			fmt.Printf("Прямоугольник: ширина %v, высота %v\n", shape.Length, shape.Height)
		case *area.Triangle:
			fmt.Printf("Треугольник: основание %v, высота %v\n", shape.Base, shape.Height)
		}
		fmt.Println("Площадь:", areaVal)
	}
}
