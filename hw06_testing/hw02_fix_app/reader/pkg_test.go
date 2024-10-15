package reader

import (
	"os"
	"testing"

	"github.com/Krovaldo/OtusHW/hw06_testing/hw02_fix_app/types"

	"github.com/stretchr/testify/assert"
)

func TestReadJSON(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_data_*.json")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	testData := `[
		{"userId": 1, "age": 30, "name": "John Doe", "departmentId": 101},
		{"userId": 2, "age": 25, "name": "Jane Smith", "departmentId": 102}
	]`

	_, err = tempFile.Write([]byte(testData))
	assert.NoError(t, err)

	err = tempFile.Close()
	assert.NoError(t, err)

	employees, err := ReadJSON(tempFile.Name())
	assert.NoError(t, err)
	assert.Len(t, employees, 2)

	expectedEmployees := []types.Employee{
		{UserID: 1, Age: 30, Name: "John Doe", DepartmentID: 101},
		{UserID: 2, Age: 25, Name: "Jane Smith", DepartmentID: 102},
	}
	assert.Equal(t, employees, expectedEmployees)
}
