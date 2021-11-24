# go-charts

`go-charts` is simple pixel precision Go real world process chart implementation

![Example](/chart.png?raw=true)

## Source code

```golang
c := New(320, 240)
c.RegisterInterval(1, 1, 2)
c.RegisterInterval(2, 2, 3)
c.RegisterInterval(3, 3, 5)
c.RegisterInterval(4, 5, 5)
c.RegisterInterval(5, 6, 12)
c.Render("chart.png")
```

## Quick overview

Debug real world process for example work with hardware device or debug time bases process require understanding start and stop time quant.

You may manipulate with scale and provide second, microsecond or nanosecond percission.

Main limitation is using integer values elsewhere.
