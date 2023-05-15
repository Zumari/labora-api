package math

import (
	"fmt"
	"testing"
)

// arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
type addTest struct {
	arg1, arg2, expected int
}

// Prueba de factorial
type factorialTest struct {
	arg, expected int
}

var factTests = []factorialTest{
	{0, 1},
	{1, 1},
	{2, 2},
	{5, 120},
}

func TestFact(t *testing.T) {

	for _, test := range factTests {
		if output := Factorial(test.arg); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func ExampleFact() {
	fmt.Println(Factorial(5))
	// Output: 120
}

// Fin de pruebas en factorial

var addTests = []addTest{
	{2, 3, 5},
	{4, 8, 12},
	{6, 9, 15},
	{3, 10, 13},
}

func TestAdd(t *testing.T) {

	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(4, 6)
	}
}

func ExampleAdd() {
	fmt.Println(Add(4, 6))
	// Output: 10
}
