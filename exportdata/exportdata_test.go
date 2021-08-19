package exportdata

import (
	"fmt"
	"testing"
)

func TestExportData(t *testing.T) {

	users, err := QueryUsers(10)
	fmt.Println("Users", users)
	if err != nil {
		t.Errorf(`QueryUsers() = %q, %v  want "", error`, users, err)
	}
}
