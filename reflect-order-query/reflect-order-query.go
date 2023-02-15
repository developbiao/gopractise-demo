package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	firstName string
	lastName  string
	id        int
	salary    int
	country   string
	address   string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct {
		t := reflect.TypeOf(q).Name()
		query := fmt.Sprintf("INSERT INTO %s VALUES(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int:
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else {
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())

				} else {
					query = fmt.Sprintf("%s \"%s\"", query, v.Field(i).String())
				}
			default:
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println("SQL:", query)
	}
}

func main() {
	order := order{
		ordId:      123,
		customerId: 8795,
	}
	createQuery(order)

	xiaoming := employee{
		id:        132,
		address:   "Tian an meng",
		country:   "china",
		firstName: "xiao",
		lastName:  "ming",
		salary:    8000,
	}
	createQuery(xiaoming)

}
