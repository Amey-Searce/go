package math

import "testing"

func TestAverage(t *testing.T) {
	v := []int{34, 34}
	avg := 0
	sum := 0
	for _, value := range v {
		sum += value
	}
	avg = sum / 2
	if avg != 34 {
		t.Error("Expected 34, got ", v)
	}
}
