package fibonacci

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci0(t *testing.T) {
	result := Fibonacci(0)
	expected := 0
	assert.Equal(t, expected, result, "Fibonacci(0) should return 0")
}

func TestFibonacci1(t *testing.T) {
	result := Fibonacci(1)
	expected := 1
	assert.Equal(t, expected, result, "Fibonacci(1) should return 1")
}

func TestFibonacci2(t *testing.T) {
	result := Fibonacci(2)
	expected := 1
	assert.Equal(t, expected, result, "Fibonacci(2) should return 1")
}

func TestFibonacci5(t *testing.T) {
	result := Fibonacci(5)
	expected := 5
	assert.Equal(t, expected, result, "Fibonacci(5) should return 5")
}

func TestFibonacci10(t *testing.T) {
	result := Fibonacci(10)
	expected := 55
	assert.Equal(t, expected, result, "Fibonacci(10) should return 55")
}
func TestFibonacci1000(t *testing.T) {
	result := Fibonacci(50)
	expected := 12586269025
	assert.Equal(t, expected, result, "Fibonacci(50) should return 12586269025")
}

func TestFibonacciNegative(t *testing.T) {
	result := Fibonacci(-5)
	expected := 0
	assert.Equal(t, expected, result, "Fibonacci(-5) should return 0 (based on our assumption that negative inputs return 0)")
}
