package chart

import (
	"testing"
)

func TestChart(t *testing.T) {
	c := New(320, 240)
	c.SetScale(4, 4)
	c.RegisterInterval(1, 1, 2)
	c.RegisterInterval(2, 2, 3)
	c.RegisterInterval(3, 3, 5)
	c.RegisterInterval(4, 5, 5)
	c.RegisterInterval(5, 6, 12)
	c.Render("chart.png")

}
