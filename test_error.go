package main

import (
	"fmt"
	"errors"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagian Dengan NOL")
	}else {
		return nilai / pembagi, nil
	}
}

// custom error
type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "validation error"}
	}
	
	if id != "eko" {
		return &notFoundError{Message: "data not found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error:", validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error:", notFoundErr.Message)
		} else {
			fmt.Println("unknown error:", err.Error())
		}
	}
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("error", err.Error())
	}
}