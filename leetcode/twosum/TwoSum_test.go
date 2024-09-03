package twosum

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
import _ "github.com/stretchr/testify/assert"

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	result := twoSum(nums, target)
	assert.Equal(t, result, expected,
		"Expected %v, but got %v", expected, result,
	)
}
func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}
	result := twoSum(nums, target)
	assert.Equal(t, expected, result, "Expected %v, but got %v", expected, result)
}
func TestTwoSum3(t *testing.T) {
	nums := []int{3, 3}
	target := 6
	expected := []int{0, 1}
	result := twoSum(nums, target)
	assert.Equal(t, expected, result, "Expected %v, but got %v", expected, result)
}

func TestTwoSum(t *testing.T) {

	arr := [...]int{2, 3, 4, 5, 6, 7}
	mikes := details{
		name: "mike",
		age:  25,
	}
	println("Print using print")
	println(mikes)
	print("Print using fmt")
	fmt.Println(arr)
}

type details struct {
	name string
	age  int
}
