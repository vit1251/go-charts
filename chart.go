package chart

import "github.com/fogleman/gg"

type AxisX struct {
	width int32
	height int32
}

type AxisY struct {
	width int32
	height int32
}

type Chart struct {
	width int32
	height int32
}

func New(width int32, height int32) (*Chart) {
	c := &Chart{
		width: width,
		height: height,
	}
	return c
}

func (c *Chart) RegisterInterval(y int32, start_x int32, stop_y int32) {
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
