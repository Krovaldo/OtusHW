package main

import (
	"github.com/Krovaldo/OtusHW/hw06_testing/hw05_shapes/area"
	printerarea "github.com/Krovaldo/OtusHW/hw06_testing/hw05_shapes/printer"
)

func main() {
	c := &area.Circle{Radius: 7}
	r := &area.Rectangle{Length: 5, Height: 6}
	t := &area.Triangle{Base: 8, Height: 5}
	n := "not a shape"

	printerarea.PrintArea(c)
	printerarea.PrintArea(r)
	printerarea.PrintArea(t)
	printerarea.PrintArea(n)
}
