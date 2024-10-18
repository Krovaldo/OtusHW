package area

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircleArea(t *testing.T) {
	circle := &Circle{Radius: 5}
	result, err := CalculateArea(circle)
	expected := 5 * 5 * math.Pi
	assert.NoError(t, err)
	assert.Equal(t, result, expected)
}

func TestTriangleArea(t *testing.T) {
	triangle := &Triangle{Base: 4, Height: 5}
	result, err := CalculateArea(triangle)
	expected := float64((4 * 5) / 2)
	assert.NoError(t, err)
	assert.Equal(t, result, expected)
}

func TestRectangleArea(t *testing.T) {
	rectangle := &Rectangle{Length: 5, Height: 8}
	result, err := CalculateArea(rectangle)
	expected := float64(5 * 8)
	assert.NoError(t, err)
	assert.Equal(t, result, expected)
}
