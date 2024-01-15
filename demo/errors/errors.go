package main

import (
	"errors"
	"fmt"
)

type DivError string

func (de DivError) Error() string {
	return fmt.Sprintf("error dividing: %v", string(de))
}

func divide(lhs, rhs int) (int, error) {
	if rhs == 0 {
		divError := DivError("ZeroDivError")
		return 0, divError
	}
	return lhs / rhs, nil
}

func main() {
	lhs, rhs := 10, 0
	res, err := divide(lhs, rhs)
	if err != nil {
		zeroDivErr := DivError("ZeroDivError")
		if errors.Is(err, zeroDivErr) {
			fmt.Println("Division by Zero Error:", err)
		} else {
			fmt.Println("OtherError:", err)
		}
	} else {
		fmt.Println(res)
	}
}
