package model

import (
	"fmt"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	fmt.Println("Test calculateDistance...")
	distance_expected := 7.280109889280518
	distance_calculated := calculateDistance(20, 10, 27, 12)
	fmt.Println("Distance Expected: ", distance_expected)
	fmt.Println("Distante Calculated: ", distance_calculated)
	if distance_calculated != distance_expected {
		t.Error("Error calculateDistance. ")
	}
}
