package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmployeeString(t *testing.T) {
	employee := Employee{
		UserID:       1,
		Age:          30,
		Name:         "John Doe",
		DepartmentID: 101,
	}

	expected := "User ID: 1; Age: 30; Name: John Doe; Department ID: 101; "
	assert.Equal(t, expected, employee.String())
}
