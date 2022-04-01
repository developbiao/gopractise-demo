package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// response 1
type response1 struct {
	Page   int
	Fruits []string
}

// response 2
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(123)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.1415926)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("Hi gopher")
	fmt.Println(string(strB))

	sclD := []string{"apple", "peach", "pear"}
	sclB, _ := json.Marshal(sclD)
	fmt.Println(string(sclB))

	mapD := map[string]int{"apple": 5, "letuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   3,
		Fruits: []string{"apple", "peach", "pear", "blana"},
	}

	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "moto", "oppo", "TCL"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.314, "deveices":["moto", "iphone", "TCL"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	str := `{"page":1, "fruits": ["apple", "moto"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
