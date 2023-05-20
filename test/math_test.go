package math_test

import (
	"testing"

	"go-intermediate/pkg/math"
)

type AddData struct {
	x, y, result int
}

func TestAdd(t *testing.T) {

	/*result := math.Add(1, 3)

	if result != 4 {
		t.Errorf("Add(1, 3) FAILED. Expected %d, got %d", 4, result)
	} else {
		t.Logf("Add(1, 3) SUCCESS. Expected %d, got %d", 4, result)
	}*/

	testData := []AddData{
		{1, 2, 3},
		{3, 5, 8},
		{7, -4, 3},
	}

	for _, datum := range testData {
		result := math.Add(datum.x, datum.y)

		if result != datum.result {
			t.Errorf("Add(1, 3) FAILED. Expected %d, got %d", datum.result, result)
		} else {
			t.Logf("Add(1, 3) SUCCESS. Expected %d, got %d", datum.result, result)
		}
	}
}

func TestSubstract(t *testing.T) {

	result := math.Substract(5, 2)

	if result != 3 {
		t.Errorf("Substract(5, 2) FAILED. Expected %d, got %d", 3, result)
	} else {
		t.Logf("Substract(5, 2) FAISUCCESSLED. Expected %d, got %d", 3, result)
	}

}
