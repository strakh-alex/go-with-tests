package structs

import "testing"
import "fmt"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 10.0}, want: 40.0},
		{shape: Circle{Radius: 10.0}, want: 62.83185307179586},
		{shape: Triangle{Base: 12, Height: 6}, want: 31.41640786499874},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("%#v gotPerimeter %g hasPerimeter %g", tt.shape, got, tt.want)
		}
	}

	//checkPerimeter := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Perimeter()
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}
	//
	//t.Run("rectangle perimeter", func(t *testing.T) {
	//	rectangle := Rectangle{10.0, 10.0}
	//	checkPerimeter(t, rectangle, 40.0)
	//})
	//
	//t.Run("circle perimeter", func(t *testing.T) {
	//	circle := Circle{10.0}
	//	checkPerimeter(t, circle, 62.83185307179586)
	//})
	//
	//t.Run("triangle perimeter", func(t *testing.T) {
	//	triangle := Triangle{12, 6}
	//	checkPerimeter(t, triangle, 31.41640786499874)
	//})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
		desc  string
	}{
		{Rectangle{12, 6}, 72.0, "retangle"},
		{Triangle{12, 6}, 36.0, "triangle"},
		{Circle{10}, 314.1592653589793, "circle"},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		fmt.Println(tt.desc)
		if got != tt.want {
			t.Errorf("%#v gotArea %g hasArea %g", tt.shape, got, tt.want)
		}
	}

	//	checkArea := func(t testing.TB, shape Shape, want float64) {
	//		t.Helper()
	//		got := shape.Area()
	//		if got != want {
	//			t.Errorf("got %g want %g", got, want)
	//		}
	//	}
	//
	//	t.Run("rectangle", func(t *testing.T) {
	//		rectangle := Rectangle{12.0, 6.0}
	//		checkArea(t, rectangle, 72.0)
	//	})
	//
	//	t.Run("circle", func(t *testing.T) {
	//		circle := Circle{10}
	//		checkArea(t, circle, 314.1592653589793)
	//	})
}
