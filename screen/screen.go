package screen

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// Screen defines the current layout.
type Screen func(gtx layout.Context) (Screen, layout.Dimensions)

var (
	s      = unit.Dp(5)
	in     = layout.UniformInset(s) // Default inset.
	spacer = layout.Spacer{Width: s, Height: s}
)

func rowInset(w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions { return in.Layout(gtx, w) }
}

//func widgetColour(th *material.Theme) layout.Dimensions {
//	var gtx layout.Context
//	darkContrast := th.ContrastBg
//	darkContrast.A = 0x33
//	return layout.Stack{}.Layout(gtx,
//		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
//			color := darkContrast
//
//			max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
//			paint.FillShape(gtx.Ops, color, clip.Rect{Max: max}.Op())
//			return layout.Dimensions{Size: gtx.Constraints.Max}
//		}),
//	)
//}

//var studentsLayout = func(gtx layout.Context) layout.Dimensions {
//	var th *material.Theme
//	return layout.Stack{}.Layout(gtx,
//		layout.Expanded(func(gtx layout.Context) layout.Dimensions {
//			darkContrast := th.ContrastBg
//			darkContrast.A = 0x33
//			color := darkContrast
//
//			max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
//			paint.FillShape(gtx.Ops, color, clip.Rect{Max: max}.Op())
//			return layout.Dimensions{Size: gtx.Constraints.Min}
//		}),
//		layout.Stacked(rowInset(material.Body1(th, fmt.Sprintln("")).Layout)),
//	)
//}

//func widgetColour(gtx layout.Context) (Screen, layout.Dimensions) {
//	d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
//		layout.Flexed(1000, rowInset(studentsLayout)),
//	)
//	return nil, d
//}
