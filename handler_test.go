package main

import (
	"fmt"
	"math"
	"testing"
)

func TestAreaOfRect(t *testing.T) {
	obj := Rect{Width: 2, Length: 3}
	actual := AreaCounter{obj}.Count()
	expected := float64(6)
	if actual != expected {
		t.Errorf("TestAreaOfRect() expected:%f but got:%f", expected, actual)
	}
}

func TestAreaOfCircle(t *testing.T) {
	obj := Circle{Radius: 10}
	actual := AreaCounter{obj}.Count()
	expected := math.Pi * 10 * 10
	if actual != expected {
		t.Errorf("TestAreaOfCircle() expected:%f but got:%f", expected, actual)
	}
}

func TestAreOfRectButWrong(t *testing.T) {
	obj := Rect{Width: 2, Length: 3}
	actual := AreaCounter{obj}.Count()
	expected := float64(5)
	if actual != expected {
		t.Errorf("TestAreOfRectButWrong() expected:%f but got:%f", expected, actual)
	}
}

func TestPerimeters(t *testing.T) {
	params := []Geometry{
		Rect{Width: 2, Length: 3},
		Circle{Radius: 10},
		Rect{Width: 3, Length: 2},
	}
	expectedResults := []float64{
		float64(10),
		float64(math.Pi * 10 * 2),
		float64(11),
	}

	for index, param := range params {
		testname := fmt.Sprintf("TestPerimeters() : %d", index)
		t.Run(testname, func(t *testing.T) {
			actual := PerimeterCounter{param}.Count()
			expected := expectedResults[index]
			if actual != expected {
				t.Errorf("expected:%f but got:%f", expected, actual)
			}
		})
	}
}
