package printer

import (
	"testing"

	"github.com/Krovaldo/OtusHW/hw06_testing/hw02_fix_app/types"
	"github.com/stretchr/testify/assert"
)

func TestPrinter(t *testing.T) {
	employees := []types.Employee{
		{UserID: 1, Age: 30, Name: "John Doe", DepartmentID: 101},
		{UserID: 2, Age: 25, Name: "Jane Smith", DepartmentID: 102},
	}

	expected := "User ID: 1; Age: 30; Name: John Doe; Department ID: 101; \n" +
		"User ID: 2; Age: 25; Name: Jane Smith; Department ID: 102; \n"

	result := PrintStaff(employees)
	assert.Equal(t, result, expected)
}
