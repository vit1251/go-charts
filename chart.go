package chart

import "github.com/fogleman/gg"

type Chart struct {
	Width int32
	Height int32
}

func New(width int32, height int32) (*Chart) {
	c := &Chart{
		Width: width,
		Height: height,
	}
	return c
}

func (c *Chart) Render(name string) {
	dc := gg.NewContext(1000, 1000)
	dc.DrawCircle(500, 500, 400)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.SavePNG(name)
}
