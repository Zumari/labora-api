package models

import (
	"testing"
	"time"
)

var itemTest = []Item{
	{1, "Samantha", time.Now(), "Memoria USB 1TB", 2, 150, " ", 0.0, 0},
	{2, "Pedro", time.Now(), "Vino La Rosa", 3, 200, " ", 0.0, 0},
	{3, "Selim", time.Now(), "Caja plástica", 1, 55, " ", 0.0, 0},
	{4, "Gloria", time.Now(), "Cepillo dental", 3, 70, " ", 0.0, 0},
	{5, "Alejandra", time.Now(), "Mouse", 3, 45, " ", 0.0, 0},
	{6, "Gloria", time.Now(), "Galleta", 2, 50, " ", 0.0, 0},
}

var expectedTotalValue = []float32{300.0, 600.0, 55.0, 210.0, 135.0, 100.0}

// TestTotalPriceWorks check that the total prices of items are generated correctly.
func TestTotalPriceWorks(t *testing.T) {

	for i, test := range itemTest {

		generatedTotalValue := test.GeneratorTotalPrice()

		if expectedTotalValue[i] != generatedTotalValue {
			t.Errorf("Output %v not equal to expected %v", generatedTotalValue, expectedTotalValue)
		}
	}
}
