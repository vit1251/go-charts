package chart

import "log"

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

}

func NewAxisX(c *Chart) (*AxisX) {

	/* Create new axis instance */
	a_x := &AxisX{}

	/* Setup members */
	a_x.step = 8
	a_x.size = 8
	a_x.marker = 4
	a_x.markerSize = 16

	return a_x
}

type AxisY struct {
}

func NewAxisY(c *Chart) (*AxisY) {

	/* Create new axis instance */
	a_y := &AxisY{}

	/* Setup members */

	return a_y
}

type Interval struct {
	Y int
	StartX int
	StopX int
}

type Chart struct {
	width int
	height int
	padding Padding
	intervals []Interval
}

func New(width int, height int) (*Chart) {

	/* Create chart instance */
	c := &Chart{
		width: width,
		height: height,
	}

	/* Setup padding */
	c.padding.Left = 32
	c.padding.Right = 8
	c.padding.Top = 8
	c.padding.Bottom = 32

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

	/* Draw values */
	for _, i := range c.intervals {

		/* Calculate scale */
		scaleX := 10.0
		scaleY := 10.0

		/* Prepare interval coords */
		x1 := scaleX * float64(i.StartX)
		y1 := scaleY * float64(i.Y)

		x2 := scaleX * float64(i.StopX)
		y2 := scaleY * float64(i.Y)

		/* Make clipping */
		if x1 < 0.0 {
			x1 = 0.0
		}
		if x2 < 0.0 {
			x2 = 0.0
		}
		if x1 > float64(c.width) { /* TODO - Calculate rect with padding ... */
			x1 = float64(c.width)
		}
		if x2 > float64(c.width) { /* TODO - Calculate rect with padding ... */
			x2 = float64(c.width)
		}

		/* Apply padding */
		x1 += float64(c.padding.Left)
		y1 += float64(c.padding.Top)
		x2 += float64(c.padding.Left)
		y2 += float64(c.padding.Top)

		/* Debug message */
		log.Printf("DrawLine( %f, %f, %f, %f )", x1, y1, x2, y2 )

		/* Draw visible interval */
		dc.SetRGB(0.4, 0.4, 0.4)
		dc.SetLineWidth( 3 )
		dc.DrawLine( x1, y1, x2, y2 )
		dc.Stroke()
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

	/* Create AxisX structure */
	a_x := NewAxisX(c)
	log.Printf("a_x = %v", a_x)

        /* Draw baseline */
//	dc.SetRGB(0.4, 0.4, 0.4)
//	dc.SetLineWidth( 0.0 )
//	dc.DrawLine( float64(a_x.StartX), float64(a_x.StartY), float64(a_x.StopX), float64(a_x.StopY) )
//	dc.Stroke()

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

	/* Create AxisY structure */
	a_y := NewAxisY(c)
	log.Printf("a_y = %v", a_y)
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

	/* Clear image */
	dc.SetRGB(255, 255, 255)
	dc.Clear()

	/* Create and draw grids*/

	/* Create and draw axis */
	c.RenderAxes(dc)

	/* Draw values */
	c.RenderValues(dc)

	/* Store chart */
	dc.SavePNG(name)
}
