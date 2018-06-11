package chart

import "github.com/fogleman/gg"

type Padding struct {
	Left int /* Default 16 */
	Right int /* Default 16 */
	Top int /* Default 16 */
	Bottom int /* Default 16 */
}

type AxisX struct {
    width int
    height int
    step int /* Default 8 */
    size int /* Default 8 */
    marker int /* Default 4 */
    markerSize int /* Default 16 */
    padding Padding
}

func NewAxisX() (*AxisX) {

	/* Create new axis instance */
	a_x := &AxisX{}

	/* Setup members */
	a_x.step = 8
	a_x.size = 8
	a_x.marker = 4
	a_x.markerSize = 16

	/* Setup padding */
	a_x.padding.Left = 32
	a_x.padding.Right = 8
        a_x.padding.Ttop = 8
        a_x.padding.Bottom = 32

	return a_x
}

type AxisY struct {
}

type Interval struct {
	Y int
	StartX int
	StopX int
}

type Chart struct {
	width int
	height int
	intervals []Interval
}

func New(width int, height int) (*Chart) {
	c := &Chart{
		width: width,
		height: height,
	}
	return c
}

func (c *Chart) RegisterInterval(y int, startX int, stopX int) {
	interval := Interval{
		Y: y,
		StartX: startX,
		StopX: stopX,
	}
	c.intervals = append(c.intervals, interval)
}

func (c *Chart) RenderValues(dc *gg.Context) {

	dc.SetRGB(0, 0, 0)
	dc.Fill()

	/* Draw values */
	for _, i := range c.intervals {
//        start = (axis_x.grid_start_x + axis_x.step * value_start_x, axis_x.grid_stop_y - axis_x.step * value_y)
//        stop = (axis_x.grid_start_x + axis_x.step * value_stop_x, axis_x.grid_stop_y - axis_x.step * value_y)
//        draw.line([start, stop], fill=color, width=3)
	}
}

func (c *Chart) RendexGridX(dc *gg.Context) {

    /* Draw grid on X */
//    for x in range(axis_x.grid_start_x, axis_x.grid_stop_x, axis_x.step * axis_x.marker):
//        start = (x, 50)
//        stop  = (x, axis_x.start_y)
//        draw.line([start, stop], fill=axis_x.grid_color, width=0)

}

func (c *Chart) RenderAxesX(dc *gg.Context) {

	a_x := NewAxisX()

        /* Draw baseline */
//        draw.line([(axis_x.start_x, axis_x.start_y), (axis_x.stop_x, axis_x.stop_y)], fill=axis_x.color, width=0)

        /* Draw smaller scale */
//        for x in range(axis_x.grid_start_x, axis_x.grid_stop_x, axis_x.step):
//            start = (x, axis_x.start_y)
//            stop  = (x, axis_x.start_y + axis_x.size)
//            draw.line([start, stop], fill=axis_x.color, width=0)

        /* Draw medium scale */
//        for x in range(axis_x.grid_start_x, axis_x.grid_stop_x, axis_x.step * axis_x.marker):
//            start = (x, axis_x.start_y)
//            stop  = (x, axis_x.start_y + axis_x.marker_size)
//            draw.line([start, stop], fill=axis_x.color, width=0)

}

func (c *Chart) RenderAxesY(dc *gg.Context) {
	axis_y := &AxisY{c.width, c.height}
}

func (c *Chart) RenderAxes(dc *gg.Context) {

	/* Draw Acis X */
	c.RenderAxesX(dc)

	/* Draw Acis Y */
	c.RenderAxesY(dc)

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
