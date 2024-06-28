package must_test

import (
	"fmt"
	"os"

	. "github.com/g-must/must"
)

func ExampleMust() {
	f := Must(os.Stat("."))

	fmt.Println(f.Name())
	// Output: .
}

func ExampleMust_error() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error")
		}
	}()

	Must(os.Stat(""))
	// Output: error
}
