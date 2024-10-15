package printer

import (
	"fmt"
	"strings"

	"github.com/Krovaldo/OtusHW/hw06_testing/hw02_fix_app/types"
)

func PrintStaff(staff []types.Employee) string {
	var res strings.Builder
	for i := 0; i < len(staff); i++ {
		str := fmt.Sprintf("User ID: %d; Age: %d; Name: %s; Department ID: %d; ",
			staff[i].UserID, staff[i].Age, staff[i].Name, staff[i].DepartmentID)
		res.WriteString(str)
		res.WriteString("\n")
	}
	return res.String()
}
