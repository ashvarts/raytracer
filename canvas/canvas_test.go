package canvas_test

import (
	"testing"

	"github.com/ashvarts/raytracer/artcolor"
	"github.com/ashvarts/raytracer/canvas"
)

func TestNewCanvas(t *testing.T) {
	c := canvas.NewCanvas(10, 20)
	if c.Width != 10 || c.Height != 20 {
		t.Errorf("expected width=10, got=%d; height=20,got=%d", c.Width, c.Height)
	}

	if c.Pixels == nil {
		t.Fatal("pixels should not be nil")
	}

	for _, col := range c.Pixels {
		if col != artcolor.NewColor(0, 0, 0) {
			t.Errorf("pixel should be color(0,0,0)")
		}
	}
}

func TestWritingPixelToCanvas(t *testing.T) {
	c := canvas.NewCanvas(10, 20)
	red := artcolor.NewColor(1, 0, 0)

	c.WritePixel(2, 3, red)
	gott := c.PixelAt(2, 3)
	_ = gott
	if got := c.PixelAt(2, 3); got != red {
		t.Errorf("Expected pixel to be %v, got:%v", red, got)
	}
}
