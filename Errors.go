package main

//In Go it’s idiomatic to communicate errors via an explicit, separate return value

import (
	"errors"
	"fmt"
)

//By convention, errors are the last return value and have type error, a built-in interface.
//interface in next project

func f1(arg int) (int, error) {

	if arg == 42 {

		return -1, errors.New("can't work with 42")

	}

	//errors.New constructs a basic error value with the given error message

	return arg + 3, nil

}

//A nil value in the error position indicates that there was no error.

type argError struct {
	arg int

	prob string
}

//It’s possible to use custom types as errors by implementing the Error() method on them.
// Here’s a variant on the example above that uses a custom type to explicitly represent an argument error.

func (e *argError) Error() string {

	return fmt.Sprintf("%d - %s", e.arg, e.prob)

}

func f2(arg int) (int, error) {

	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}

	}

	return arg + 3, nil

}

//In this case we use &argError syntax to build a new struct, supplying values for the two fields arg and prob.

func main() {

	//The two loops below test out each of our error-returning functions.
	// Note that the use of an inline error check on the if line is a common idiom in Go code.

	for _, i := range []int{7, 42} {

		if r, e := f1(i); e != nil {

			fmt.Println("f1 failed:", e)

		} else {

			fmt.Println("f1 worked:", r)

		}

	}

	for _, i := range []int{7, 42} {

		if r, e := f2(i); e != nil {

			fmt.Println("f2 failed:", e)

		} else {

			fmt.Println("f2 worked:", r)

		}

	}

	_, e := f2(42)

	//_,  : This sign means do not consider

	if ae, ok := e.(*argError); ok {

		fmt.Println(ae.arg)

		fmt.Println(ae.prob)

	}

}

//If you want to programmatically use the data in a custom error,
// you’ll need to get the error as an instance of the custom error type via type assertion.

//output :
//f1 worked: 10
//f1 failed: can't work with 42
//f2 worked: 10
//f2 failed: 42 - can't work with it
//42
//can't work with it
