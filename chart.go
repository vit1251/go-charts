package chart

import (
	gg "github.com/vit1251/go-charts/pixelman"
	"log"
)

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

	a_x.Step = 16

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

	a_y.Step = 16

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

func NewGrid(scaleX int, scaleY int) *Grid {
	g := &Grid{
		ScaleX: scaleX,
		ScaleY: scaleY,
	}
	return g
}

type Size struct {
	Width int
	Height int
}

type Chart struct {
	size		Size
	padding		Padding
	intervals	[]Interval
	grid		*Grid
	scaleX		int
	scaleY		int
}

func (c *Chart) SetScale(scaleX int, scaleY int) {

	/* Update scale */
	c.scaleX = scaleX
	c.scaleY = scaleY

	/* Setup grid */
	c.grid = NewGrid(c.scaleX, c.scaleY)

}

func New(width int, height int) (*Chart) {

	/* Create chart instance */
	c := new(Chart)

	/* Setup size */
	c.size.Width = width
	c.size.Height = height

	/* Setup padding */
	c.padding.Left = 32
	c.padding.Right = 32
	c.padding.Top = 32
	c.padding.Bottom = 32

	/* Set scale X and Y */
	c.scaleX = 8
	c.scaleY = 8

	/* Setup grid */
	c.grid = NewGrid(c.scaleX, c.scaleY)

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

		var scaleX int
		var scaleY int

		/* Calculate scale */
		if c.grid != nil {
			scaleX = c.grid.ScaleX
			scaleY = c.grid.ScaleY
		} else {
			scaleX = 8
			scaleY = 8
		}

		/* Prepare interval coords */
		x1 := rect.Left + scaleX * i.StartX
		y1 := rect.Bottom - scaleY * i.Y

		x2 := rect.Left + scaleX * i.StopX
		y2 := rect.Bottom - scaleY * i.Y

		/* Make clipping X coords */
		if x1 < rect.Left {
			x1 = rect.Left
		}
		if x2 < rect.Left {
			x2 = rect.Left
		}
		if x1 > rect.Right {
			x1 = rect.Right
		}
		if x2 > rect.Right {
			x2 = rect.Right
		}

		/* Make clipping Y coords */
		if y1 < rect.Top {
			y1 = rect.Top
		}
		if y2 < rect.Top {
			y2 = rect.Top
		}
		if y1 > rect.Bottom {
			y1 = rect.Bottom
		}
		if y2 > rect.Bottom {
			y2 = rect.Bottom
		}

		/* Draw visible interval */
		dc.SetRGB( 128, 0, 0 )
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
			x1 := c_x
			y1 := rect.Top

			x2 := c_x
			y2 := rect.Bottom

			/* Draw grid */
			dc.SetRGB(192, 192, 192)
			dc.SetLineWidth( 1 )
			dc.DrawLine( x1, y1, x2, y2 )
			dc.Stroke()
		}

		/* Draw grid on Y */
		for c_y := rect.Top + c.grid.ScaleY; c_y < rect.Bottom; c_y += c.grid.ScaleY {

			/* Prepare parameters */
			x1 := rect.Left
			y1 := c_y

			x2 := rect.Right
			y2 := c_y

			/* Draw grid */
			dc.SetRGB(192, 192, 192)
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
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth( 1 )
	dc.DrawLine( a_x.StartX, a_x.StartY, a_x.StopX, a_x.StopY )
	dc.Stroke()

        /* Draw scale */
	for c_x := a_x.StartX + a_x.Step; c_x < a_x.StopX; c_x += a_x.Step {

		/* Prepare step position */
		x1 := c_x
		y1 := a_x.StartY

		x2 := c_x
		y2 := a_x.StartY - 4

		/* Draw risk */
		dc.SetRGB(0, 0, 0)
		dc.SetLineWidth( 1 )
		dc.DrawLine( x1, y1, x2, y2 )
		dc.Stroke()

	}

}

func (c *Chart) RenderAxesY(dc *gg.Context) {

	/* Create AxisY structure */
	a_y := NewAxisY(c)
	log.Printf("a_y = %v", a_y)

	/* Draw baseline */
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth( 1 )
	dc.DrawLine( a_y.StartX, a_y.StartY, a_y.StopX, a_y.StopY )
	dc.Stroke()

        /* Draw scale */
	for c_y := a_y.StartY + a_y.Step; c_y < a_y.StopY; c_y += a_y.Step {

		/* Prepare step position */
		x1 := a_y.StartX
		y1 := c_y

		x2 := a_y.StartX + 4.0
		y2 := c_y

		/* Draw risk */
		dc.SetRGB( 0, 0, 0 )
		dc.SetLineWidth( 1 )
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
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth( 1 )
	dc.DrawLine( rect.Left, rect.Top, rect.Right, rect.Top )
	dc.Stroke()

	/* Bottom border  */
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth( 1 )
	dc.DrawLine( rect.Left, rect.Bottom, rect.Right, rect.Bottom )
	dc.Stroke()

	/* Left border  */
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth( 1 )
	dc.DrawLine( rect.Left, rect.Top, rect.Left, rect.Bottom )
	dc.Stroke()

	/* Right border  */
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth( 1 )
	dc.DrawLine( rect.Right, rect.Top, rect.Right, rect.Bottom )
	dc.Stroke()

}


func (c *Chart) Render(name string) {

	/* Create new drawing canvas */
	dc := gg.NewContext(c.size.Width, c.size.Height)

	/* Clear image */
	dc.SetRGB(255, 255, 255)
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
