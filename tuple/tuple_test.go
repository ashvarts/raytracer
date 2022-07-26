package tuple

import (
	"math"
	"reflect"
	"testing"

	"github.com/ashvarts/raytracer/test"
)

func Test_Tuple_isPoint(t *testing.T) {
	type fields struct {
		x float64
		y float64
		z float64
		w float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"isPoint", fields{4.3, -4.2, 3.1, 1.0}, true},
		{"isPoint", fields{4.3, -4.2, 3.1, 0.0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tuple{
				X: tt.fields.x,
				Y: tt.fields.y,
				Z: tt.fields.z,
				W: tt.fields.w,
			}
			if got := tr.IsPoint(); got != tt.want {
				t.Errorf("Tuple.isPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Tuple_isVector(t *testing.T) {
	type fields struct {
		x float64
		y float64
		z float64
		w float64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"isVector", fields{4.3, -4.2, 3.1, 1.0}, false},
		{"isVector", fields{4.3, -4.2, 3.1, 0.0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tuple{
				X: tt.fields.x,
				Y: tt.fields.y,
				Z: tt.fields.z,
				W: tt.fields.w,
			}
			if got := tr.IsVector(); got != tt.want {
				t.Errorf("Tuple.isVector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_point(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want Tuple
	}{
		{"new point", args{4, -4, 3}, Tuple{4, -4, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Point(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("point() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_vector(t *testing.T) {
	type args struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name string
		args args
		want Tuple
	}{
		{"new vector", args{4, -4, 3}, Tuple{4, -4, 3, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Vector(tt.args.x, tt.args.y, tt.args.z); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("vector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTupleAdd(t *testing.T) {
	t1 := Point(3, -2, 5)
	t2 := Vector(-2, 3, 1)

	sum := t1.Add(t2)
	want := Tuple{1, 1, 6, 1}
	if sum != want {
		t.Errorf("expected: %v, got: %v", sum, want)
	}
}

func TestTupleSub(t *testing.T) {
	p1 := Point(3, 2, 1)
	p2 := Point(5, 6, 7)

	sum := p1.Sub(p2)
	want := Vector(-2, -4, -6)
	if sum != want {
		t.Errorf("expected: %v, got: %v", sum, want)
	}
}

func TestDot(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)

	dotProd := v1.Dot(v2)
	want := 20.0
	if dotProd != want {
		t.Errorf("want: %v, got: %v", want, dotProd)
	}
}
func TestCross(t *testing.T) {
	v1 := Vector(1, 2, 3)
	v2 := Vector(2, 3, 4)

	dotProd := v1.Cross(v2)
	want := Vector(-1, 2, -1)
	if dotProd != want {
		t.Errorf("want: %v, got: %v", want, dotProd)
	}
	dotProd = v2.Cross(v1)
	want = Vector(1, -2, 1)
	if dotProd != want {
		t.Errorf("want: %v, got: %v", want, dotProd)
	}

}
func TestTuple_Sub(t *testing.T) {
	tests := []struct {
		name string
		t1   Tuple
		t2   Tuple
		want Tuple
	}{
		{"Subtract two points", Point(3, 2, 1), Point(5, 6, 7), Vector(-2, -4, -6)},
		{"Subtract a vector from point", Point(3, 2, 1), Vector(5, 6, 7), Point(-2, -4, -6)},
		{"Subtract two vectors", Vector(3, 2, 1), Vector(5, 6, 7), Vector(-2, -4, -6)},
		{"Subtract two vectors", Vector(0, 0, 0), Vector(1, -2, 3), Vector(-1, 2, -3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t1.Sub(tt.t2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTupleMultiply(t *testing.T) {
	tests := []struct {
		name string
		t1   Tuple
		s    float64
		want Tuple
	}{
		{"Multiply a tuple by a scalar", Tuple{1, -2, 3, -4}, 3.5, Tuple{3.5, -7, 10.5, -14}},
		{"Multiply a tuple by a fraction", Tuple{1, -2, 3, -4}, 0.5, Tuple{0.5, -1, 1.5, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t1.Multiply(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tuple.Multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTupleMagnitude(t *testing.T) {
	tests := []struct {
		name string
		t1   Tuple
		want float64
	}{
		{"Compute magnitude", Vector(1, 0, 0), 1},
		{"Compute magnitude", Vector(0, 1, 0), 1},
		{"Compute magnitude", Vector(0, 0, 1), 1},
		{"Compute magnitude", Vector(1, 2, 3), math.Sqrt(14)},
		{"Compute magnitude", Vector(-1, -2, -3), math.Sqrt(14)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t1.Magnitude(); got != tt.want {
				t.Errorf("Tuple.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTupleNormalize(t *testing.T) {
	tests := []struct {
		name string
		t1   Tuple
		want Tuple
	}{
		{"Compute vector normal", Vector(4, 0, 0), Vector(1, 0, 0)},
		{"Compute vector normal", Vector(1, 2, 3), Vector(0.26726, 0.53452, 0.80178)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t1.Normalize(); !AssertTupleEqual(got, tt.want) {
				t.Errorf("Tuple.Normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestTupleNegate(t *testing.T) {
	tup := Tuple{1, -2, 3, -4}
	got := tup.Negate()
	want := Tuple{-1, 2, -3, 4}

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

// AssertTupleEqual is a helper function that tests if two tuples are equial using an aproximate equilancy for each element
func AssertTupleEqual(t1 Tuple, t2 Tuple) bool {
	if test.AproxEquall(t1.X, t2.X) && test.AproxEquall(t1.Y, t2.Y) && test.AproxEquall(t1.Z, t2.Z) && test.AproxEquall(t1.W, t2.W) {
		return true
	}
	return false
}
