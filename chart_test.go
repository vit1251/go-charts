package chart

import (
	"testing"
	"image/color"
)

func TestChart(t *testing.T) {
	c := New(320, 240)
	c.SetScale(4, 4)

	/* Classic red color interval */
	c.RegisterInterval(25, 1, 2)
	c.RegisterInterval(23, 4, 12)
	c.RegisterInterval(26, 23, 34)
	c.RegisterInterval(21, 43, 52)

	/* Modern intervals */
	c.RegisterIntervalEx(5, 6, 12, color.RGBA{234,28,15,255})
	c.RegisterIntervalEx(5, 12, 21, color.RGBA{0,128,0,255})
	c.RegisterIntervalEx(5, 21, 23, color.RGBA{128,128,0,255})

	/* Random range */
	//var pos = 0
	for i := 0; i < 12; i++ {
		c.RegisterIntervalEx(10 + i, 6 + i, 12 + i, color.RGBA{234,28,15,255})
		c.RegisterIntervalEx(10 + i, 12 + i, 21 + i, color.RGBA{0,128,0,255})
		c.RegisterIntervalEx(10 + i, 21 + i, 23 + i, color.RGBA{128,128,0,255})
	}

	c.Render("chart.png")

}
