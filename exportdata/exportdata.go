package exportdata

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
	"log"
	"os"
	"strconv"
)

var db *sql.DB

type User struct {
	AccountId   int64
	Username    string
	Nickname    string
	Phone       string
	AccountType int32
}

// Init mysql config
func connectDatabase(dbName string) {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "192.168.56.38:3306",
		DBName:               dbName,
		AllowOldPasswords:    true,
		AllowNativePasswords: true,
	}

	// Get a database handle
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database connected!")

}

// Query user records
func QueryUsers(limit int64) ([]User, error) {

	dbName := os.Getenv("DB1")
	connectDatabase(dbName)

	// An users slice hold data from return rows
	var users []User

	rows, err := db.Query("SELECT `account_id`, `username`, `nickname`, `phone`, `account_type` FROM `accounts`  LIMIT 0, ?", limit)
	if err != nil {
		return nil, fmt.Errorf("QueryUsers %q: %v", limit, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to user fields
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.AccountId, &user.Username, &user.Nickname, &user.Phone, &user.AccountType); err != nil {
			return nil, fmt.Errorf("QueryUsers %q: %v", limit, err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("QueryUsers %q: %v", limit, err)
	}
	return users, nil
}

// Save users to execel file
func SaveToExcel(users []User) {
	xlsx := excelize.NewFile()
	// Create a nwe sheet
	index := xlsx.NewSheet("Sheet1")
	var user User
	i := 1
	for _, user = range users {
		xlsx.SetCellValue("Sheet1", "A"+strconv.Itoa(i), user.AccountId)
		xlsx.SetCellValue("Sheet1", "B"+strconv.Itoa(i), user.Username)
		xlsx.SetCellValue("Sheet1", "C"+strconv.Itoa(i), user.Nickname)
		xlsx.SetCellValue("Sheet1", "D"+strconv.Itoa(i), user.Phone)
		xlsx.SetCellValue("Sheet1", "E"+strconv.Itoa(i), user.AccountType)
		i++
	}

	// Set value a cell
	xlsx.SetActiveSheet(index)

	err := xlsx.SaveAs("./Accounts.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}

func ReadExcel() ([]User, error) {
	fileName := "./Accounts.xlsx"
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("ReadExcel %q, %v", fileName, err)
	}

	// Get all the rows in the sheet1
	rows, err := xlsx.GetRows("Sheet1")
	if err != nil {
		return nil, fmt.Errorf("ReadExcel %q, %v", fileName, err)
	}

	var users []User
	for _, row := range rows {
		accountId, _ := strconv.ParseInt(row[0], 10, 64)
		accountType, _ := strconv.ParseInt(row[4], 10, 32)

		user := User{accountId, row[1], row[2], row[3], int32(accountType)}
		users = append(users, user)
	}

	return users, nil
}

func WriteUsers(users []User) []int64 {
	dbName := os.Getenv("DB2")
	connectDatabase(dbName)
	var ids []int64
	for _, user := range users {
		res, err := db.Exec("INSERT INTO `users` (`account_id`, `username`, `nickname`, `phone`, `account_type`) VALUES(?, ?, ?, ?, ?)",
			user.AccountId, user.Username, user.Nickname, user.Phone, user.AccountType)
		if err != nil {
			log.Fatal(err)
		}

		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}

		rowCount, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		ids = append(ids, lastId)

		log.Printf("ID = %d, affected = %d, user = %s\n", lastId, rowCount, user.Nickname)
	}
	return ids
}
