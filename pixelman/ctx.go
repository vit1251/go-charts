package pixelman

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"errors"
)

type Context struct {
	i *image.RGBA

	width int
	height int

	r uint8
	g uint8
	b uint8

}

func NewContext(width int, height int) (*Context) {

	ctx := &Context{}

	ctx.i = image.NewRGBA(image.Rect(0, 0, width, height))
	ctx.width = width
	ctx.height = height

	return ctx
}

func (ctx *Context) SetRGB(r, g, b uint8) {
	ctx.r = r
	ctx.g = g
	ctx.b = b
}

func (ctx *Context) Clear() {

	c := color.RGBA{}
	c.R = ctx.r
	c.G = ctx.g
	c.B = ctx.b
	c.A = 255

	for x := 0; x < ctx.width; x++ {
		for y := 0; y < ctx.height; y++ {
			ctx.i.Set(x, y, c)
		}
	}
}

func (ctx *Context) SetLineWidth(width int) {
}

func (ctx *Context) vLine(x int, y1 int, y2 int) {

	c := color.RGBA{}
	c.R = ctx.r
	c.G = ctx.g
	c.B = ctx.b
	c.A = 255

	for y := y1; y <= y2; y++ {
		ctx.i.Set(x, y, c)
	}

}

func (ctx *Context) hLine(y int, x1 int, x2 int) {

	c := color.RGBA{}
	c.R = ctx.r
	c.G = ctx.g
	c.B = ctx.b
	c.A = 255

	for x := x1; x <= x2; x++ {
		ctx.i.Set(x, y, c)
	}

}

func (ctx *Context) DrawLine(x1, y1, x2, y2 int) (error) {

	if (x1 == x2) {
		ctx.vLine(x1, min(y1, y2), max(y1, y2))
		return nil
	} else if (y1 == y2) {
		ctx.hLine(y1, min(x1, x2), max(x1, x2))
		return nil
	} else {
		return errors.New("No brezinhame line implementation here")
	}

}

func (ctx *Context) Stroke() {
}

func (ctx *Context) SavePNG(filename string) (error) {

	stream, err1 := os.Create(filename)
	if err1 != nil {
		return err1
	}
	png.Encode(stream, ctx.i)  
	//os.Close(stream)
	return nil
}
