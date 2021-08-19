package main

import (
	"example.com/exportdata"
	"fmt"
)

func main() {
	importData()
}

// Export data to excel
func exportData() {
	users, err := exportdata.QueryUsers(60000)
	if err != nil {
		fmt.Println("Error handle", err)
	}
	fmt.Printf("Total rows:%d\n", len(users))
	exportdata.SaveToExcel(users)
	fmt.Println("Export to excel ok!")
}

// Import data to database
func importData() {
	users, err := exportdata.ReadExcel()
	if err != nil {
		fmt.Println("Read Excel Error:", err)
	}

	ids := exportdata.WriteUsers(users)
	fmt.Printf("Import total of users:%s\n", len(ids))
	fmt.Println("Finished!")

}
