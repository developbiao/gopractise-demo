package main

import (
	"errors"
	"fmt"
)

type areaError struct {
	width  float64
	length float64
	err    string
}

func (e *areaError) Error() string {
	return e.err
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func reactArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "lenght is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError{err: err, width: width, length: length}
	}
	return length * width, nil
}

func main() {
	length, width := -9.5, -5.36
	area, err := reactArea(length, width)
	if err != nil {
		var areaError *areaError
		if errors.As(err, &areaError) {
			if areaError.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", areaError.length)
			}
			if areaError.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", areaError.width)
			}
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("Area of rect", area)
}
