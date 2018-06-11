package chart

import "log"

import "github.com/fogleman/gg"

type Rect struct {
	Left int /* Default 16 */
	Right int /* Default 16 */
	Top int /* Default 16 */
	Bottom int /* Default 16 */
}

func NewRect() (*Rect) {
	rect := &Rect{}
	return rect
}

type Padding struct {
	Left int /* Default 16 */
	Right int /* Default 16 */
	Top int /* Default 16 */
	Bottom int /* Default 16 */
}

type AxisX struct {

	StartX int
	StartY int
	StopX int
	StopY int

	Step int

}

func NewAxisX(c *Chart) (*AxisX) {

	/* Create new axis instance */
	a_x := &AxisX{}

	/* Setup members */
	a_x.StartX = c.padding.Left
	a_x.StartY = c.height - c.padding.Bottom

	a_x.StopX = c.width - c.padding.Right
	a_x.StopY = c.height - c.padding.Bottom

	a_x.Step = 8

	return a_x
}

type AxisY struct {

	StartX int
	StartY int
	StopX int
	StopY int

	Step int

}

func NewAxisY(c *Chart) (*AxisY) {

	/* Create new axis instance */
	a_y := &AxisY{}

	/* Setup members */
	a_y.StartX = c.padding.Left
	a_y.StartY = c.padding.Top

	a_y.StopX = c.padding.Left
	a_y.StopY = c.height - c.padding.Bottom

	a_y.Step = 8

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
	c.padding.Right = 32
	c.padding.Top = 32
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

	/* Drawing area */
	rect := NewRect()
	rect.Left = c.padding.Left
	rect.Top = c.padding.Top
	rect.Right = c.width - c.padding.Right
	rect.Bottom = c.height - c.padding.Bottom

	/* Draw values */
	for _, i := range c.intervals {

		/* Calculate scale */
		scaleX := 10.0
		scaleY := 10.0

		/* Prepare interval coords */
		x1 := scaleX * float64(i.StartX)
		y1 := rect.Bottom - scaleY * float64(i.Y)

		x2 := scaleX * float64(i.StopX)
		y2 := rect.Bottom - scaleY * float64(i.Y)

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
	dc.SetRGB(0.0, 0.0, 0.0)
	dc.SetLineWidth( 1.0 )
	dc.DrawLine( float64(a_x.StartX), float64(a_x.StartY), float64(a_x.StopX), float64(a_x.StopY) )
	dc.Stroke()

        /* Draw scale */
	for c_x := a_x.StartX + a_x.Step; c_x < a_x.StopX; c_x += a_x.Step {

		/* Prepare step position */
		x1 := float64(c_x)
		y1 := float64(a_x.StartY)

		x2 := float64(c_x)
		y2 := float64(a_x.StartY) - 4.0

		/* Draw risk */
		dc.SetRGB(0.0, 0.0, 0.0)
		dc.SetLineWidth( 1.0 )
		dc.DrawLine( x1, y1, x2, y2 )
		dc.Stroke()

	}

}

func (c *Chart) RenderAxesY(dc *gg.Context) {

	/* Create AxisY structure */
	a_y := NewAxisY(c)
	log.Printf("a_y = %v", a_y)

	/* Draw baseline */
	dc.SetRGB(0.0, 0.0, 0.0)
	dc.SetLineWidth( 1.0 )
	dc.DrawLine( float64(a_y.StartX), float64(a_y.StartY), float64(a_y.StopX), float64(a_y.StopY) )
	dc.Stroke()

        /* Draw scale */
	for c_y := a_y.StartY + a_y.Step; c_y < a_y.StopY; c_y += a_y.Step {

		/* Prepare step position */
		x1 := float64(a_y.StartX)
		y1 := float64(c_y)

		x2 := float64(a_y.StartX) + 4.0
		y2 := float64(c_y)

		/* Draw risk */
		dc.SetRGB(0.0, 0.0, 0.0)
		dc.SetLineWidth( 1.0 )
		dc.DrawLine( x1, y1, x2, y2 )
		dc.Stroke()

	}

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
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Clear()

	/* Create and draw grids*/

	/* Create and draw axis */
	c.RenderAxes(dc)

	/* Draw values */
	c.RenderValues(dc)

	/* Store chart */
	dc.SavePNG(name)
}
