package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func main() {
	// Define map
	m := make(map[string]*student)
	// Some students
	stus := []student{
		{Name: "zhang", Age: 18},
		{Name: "Li", Age: 28},
		{Name: "Ba", Age: 32},
	}
	// Add student to map
	// for _, stu := range stus {
	//     // foreach 中, stu 是结构体的拷贝副本，所以m[stu.Name] = &stu 实际一致指几同一个个指针
	//     // 最终该指针的值遍历的最后个 struct 的值拷贝
	// 	m[stu.Name] = &stu
	// }

	// Correct
	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}

	// Traverse map
	for k, v := range m {
		fmt.Println("k", "=>", k, v.Name)
	}

	fmt.Println("Main func done!")
}
