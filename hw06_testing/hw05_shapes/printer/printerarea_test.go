package printerarea

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/Krovaldo/OtusHW/hw06_testing/hw05_shapes/area"
	"github.com/stretchr/testify/assert"
)

func TestPrinter(t *testing.T) {
	testCases := []struct {
		name     string
		shape    any
		expected string
	}{
		{
			name:     "Круг",
			shape:    &area.Circle{Radius: 5},
			expected: "Круг: радиус 5\nПлощадь: 78.53981633974483\n",
		},
		{
			name:     "Прямоугольник",
			shape:    &area.Rectangle{Length: 5, Height: 8},
			expected: "Прямоугольник: ширина 5, высота 8\nПлощадь: 40\n",
		},
		{
			name:     "Треугольник",
			shape:    &area.Triangle{Base: 4, Height: 5},
			expected: "Треугольник: основание 4, высота 5\nПлощадь: 10\n",
		},
		{
			name:     "Не является фигурой",
			shape:    "not a shape",
			expected: "Ошибка: переданный объект не является фигурой.\n",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.name, func(t *testing.T) {
			var buf bytes.Buffer
			originalStdout := os.Stdout

			r, w, _ := os.Pipe() // создание нового файла для перенаправления ввода

			os.Stdout = w
			PrintArea(tC.shape)
			w.Close()
			io.Copy(&buf, r)
			os.Stdout = originalStdout

			assert.Equal(t, buf.String(), tC.expected)
		})
	}
}
