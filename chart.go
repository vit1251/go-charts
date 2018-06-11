package chart

import "github.com/fogleman/gg"

type AxisX struct {
	width int
	height int
}

type AxisY struct {
	width int
	height int
}

type Chart struct {
	width int
	height int
}

func New(width int, height int) (*Chart) {
	c := &Chart{
		width: width,
		height: height,
	}
	return c
}

func (c *Chart) RegisterInterval(y int, start_x int, stop_y int) {
}

func (c *Chart) RenderAxes(int dc, a_x *AxisX, a_y *AxisY) {
	axis_x = &AxisX{width, height}
	axis_y = &AxisY{width, height}
}

func (c *Chart) RenderValues(int dc) {
//	dc.SetRGB(0, 0, 0)
//	dc.Fill()
}

func (c *Chart) Render(name string) {

	/* Create new drawing canvas */
	dc := gg.NewContext(c.width, c.height)

	/* Create and draw axis */
	c.RenderAxes(dc)

	/* Draw values */
	c.RenderValues(dc)

	/* Store chart */
	dc.SavePNG(name)
}
