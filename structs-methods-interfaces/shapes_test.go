package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		checkArea(t, rectangle, 100)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{10, 10}
		checkArea(t, triangle, 50.0)
	})
}

func TestDrivenAreas(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10, Height: 10}, hasArea: 100},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 10}, hasArea: 50},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()
			if got != areaTest.hasArea {
				t.Errorf("got %g, want %g", got, areaTest.hasArea)
			}
		})
	}
}
