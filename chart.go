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
	a_x.StartY = c.size.Height - c.padding.Bottom

	a_x.StopX = c.size.Width - c.padding.Right
	a_x.StopY = c.size.Height - c.padding.Bottom

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
	a_y.StopY = c.size.Height - c.padding.Bottom

	a_y.Step = 8

	return a_y
}

type Interval struct {
	Y int
	StartX int
	StopX int
}

type Grid struct {
	ScaleX int
	ScaleY int
}

func NewGrid() *Grid {
	g := &Grid{}
	g.ScaleX = 16
	g.ScaleY = 16
	return g
}

type Size struct {
	Width int
	Height int
}

type Chart struct {
	size Size
	padding Padding
	intervals []Interval
	grid *Grid
}

func New(width int, height int) (*Chart) {

	/* Create chart instance */
	c := &Chart{}

	/* Setup size */
	c.size.Width = width
	c.size.Height = height

	/* Setup padding */
	c.padding.Left = 32
	c.padding.Right = 32
	c.padding.Top = 32
	c.padding.Bottom = 32

	/* Setup grid */
	c.grid = NewGrid()

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
	rect.Right = c.size.Width - c.padding.Right
	rect.Bottom = c.size.Height - c.padding.Bottom

	/* Draw values */
	for _, i := range c.intervals {

		var scaleX float64
		var scaleY float64

		/* Calculate scale */
		if c.grid != nil {
			scaleX = float64(c.grid.ScaleX)
			scaleY = float64(c.grid.ScaleY)
		} else {
			scaleX = float64(8.0)
			scaleY = float64(8.0)
		}

		/* Prepare interval coords */
		x1 := float64(rect.Left) + scaleX * float64(i.StartX)
		y1 := float64(rect.Bottom) - scaleY * float64(i.Y)

		x2 := float64(rect.Left) + scaleX * float64(i.StopX)
		y2 := float64(rect.Bottom) - scaleY * float64(i.Y)

		/* Make clipping X coords */
		if x1 < float64(rect.Left) {
			x1 = float64(rect.Left)
		}
		if x2 < float64(rect.Left) {
			x2 = float64(rect.Left)
		}
		if x1 > float64(rect.Right) {
			x1 = float64(rect.Right)
		}
		if x2 > float64(rect.Right) {
			x2 = float64(rect.Right)
		}

		/* Make clipping Y coords */
		if y1 < float64(rect.Top) {
			y1 = float64(rect.Top)
		}
		if y2 < float64(rect.Top) {
			y2 = float64(rect.Top)
		}
		if y1 > float64(rect.Bottom) {
			y1 = float64(rect.Bottom)
		}
		if y2 > float64(rect.Bottom) {
			y2 = float64(rect.Bottom)
		}

		/* Debug message */
//		log.Printf("DrawLine( %f, %f, %f, %f )", x1, y1, x2, y2 )

		/* Draw visible interval */
		dc.SetRGB(0.0, 0.0, 1.0)
		dc.SetLineWidth( 3 )
		dc.DrawLine( x1, y1, x2, y2 )
		dc.Stroke()
	}
}

func (c *Chart) RenderGrids(dc *gg.Context) {

	/* Drawing area */
	rect := NewRect()
	rect.Left = c.padding.Left
	rect.Top = c.padding.Top
	rect.Right = c.size.Width - c.padding.Right
	rect.Bottom = c.size.Height - c.padding.Bottom

	/* Grid size */
	if c.grid != nil {

		/* Draw grid on X */
		for c_x := rect.Left + c.grid.ScaleX; c_x < rect.Right; c_x += c.grid.ScaleX {

			/* Prepare parameters */
			x1 := float64(c_x)
			y1 := float64(rect.Top)

			x2 := float64(c_x)
			y2 := float64(rect.Bottom)

			/* Draw grid */
			dc.SetRGB(0.4, 0.4, 0.4)
			dc.SetLineWidth( 1 )
			dc.DrawLine( x1, y1, x2, y2 )
			dc.Stroke()
		}

		/* Draw grid on Y */
		for c_y := rect.Top + c.grid.ScaleY; c_y < rect.Bottom; c_y += c.grid.ScaleY {

			/* Prepare parameters */
			x1 := float64(rect.Left)
			y1 := float64(c_y)

			x2 := float64(rect.Right)
			y2 := float64(c_y)

			/* Draw grid */
			dc.SetRGB(0.4, 0.4, 0.4)
			dc.SetLineWidth( 1 )
			dc.DrawLine( x1, y1, x2, y2 )
			dc.Stroke()
		}
	}

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

func (c *Chart) RenderBorder(dc *gg.Context) {

	/* Drawing area */
	rect := NewRect()
	rect.Left = c.padding.Left
	rect.Top = c.padding.Top
	rect.Right = c.size.Width - c.padding.Right
	rect.Bottom = c.size.Height - c.padding.Bottom

	/* Top border  */
	dc.SetRGB(0.0, 0.0, 0.0)
	dc.SetLineWidth( 1.0 )
	dc.DrawLine( float64(rect.Left), float64(rect.Top), float64(rect.Right), float64(rect.Top) )
	dc.Stroke()

	/* Bottom border  */
	dc.SetRGB(0.0, 0.0, 0.0)
	dc.SetLineWidth( 1.0 )
	dc.DrawLine( float64(rect.Left), float64(rect.Bottom), float64(rect.Right), float64(rect.Bottom) )
	dc.Stroke()

	/* Left border  */
	dc.SetRGB(0.0, 0.0, 0.0)
	dc.SetLineWidth( 1.0 )
	dc.DrawLine( float64(rect.Left), float64(rect.Top), float64(rect.Left), float64(rect.Bottom) )
	dc.Stroke()

	/* Right border  */
	dc.SetRGB(0.0, 0.0, 0.0)
	dc.SetLineWidth( 1.0 )
	dc.DrawLine( float64(rect.Right), float64(rect.Top), float64(rect.Right), float64(rect.Bottom) )
	dc.Stroke()

}


func (c *Chart) Render(name string) {

	/* Create new drawing canvas */
	dc := gg.NewContext(c.size.Width, c.size.Height)

	/* Clear image */
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Clear()

	/* Create and draw grids*/
	c.RenderGrids(dc)

	/* Create and draw axis */
	c.RenderAxes(dc)

	/* Draw values */
	c.RenderValues(dc)

	/* Draw chart border */
	c.RenderBorder(dc)

	/* Store chart */
	dc.SavePNG(name)
}
