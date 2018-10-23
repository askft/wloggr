package store

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("store_test: init()")
	err := SetupDB("root:password@(localhost:3306)/temp")
	if err != nil {
		panic(err)
	}
}

func TestStoreCreateWorkout(t *testing.T) {
	fmt.Println("store_test: TestStoreCreateWorkout")
}
