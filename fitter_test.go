package fitter

import (
	"reflect"
	"testing"
)

func TestTransform(t *testing.T) {
	parallelogram := []float64{
		0, 2,
		1, 1,
		1, 0,
		0, 1,
	}
	inner := Square(0)
	result, _ := Transform(inner, parallelogram)

	// The result should be a parallelogram
	if !reflect.DeepEqual(result, parallelogram) {
		t.Errorf("Expected %v, got %v", parallelogram, result)
	}
}

func TestConcave(t *testing.T) {
	star := Star()
	inner := Square(0)
	_, err := Transform(inner, star)

	// The star is concave, so we should get an error
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
