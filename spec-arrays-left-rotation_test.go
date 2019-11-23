package main

import (
	"fmt"
	"testing"
)

func TestRotLeft1(t *testing.T) {
	a := []int32{1, 2, 3, 4, 5}
	d := int32(4)

	returnedArray := rotLeft(a, d)

	fmt.Println("returnedArray", returnedArray)
}

func TestRotLeft2(t *testing.T) {
	a := []int32{1, 2, 3, 4, 5}
	d := int32(13)

	returnedArray := rotLeft(a, d)

	fmt.Println("returnedArray", returnedArray)
}
