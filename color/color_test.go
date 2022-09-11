package color

import "testing"

func TestColor(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)
	if c.red != -0.5 || c.green != 0.4 || c.blue != 1.7 {
		t.Error("Color not made correctly")
	}
}

func TestColorAdd(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	want := NewColor(1.6, 0.7, 1.0)
	if got := c1.Add(c2); got != want {
		t.Errorf("wanted:%v, got:%v", want, got)
	}
}
