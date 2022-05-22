package screen

import (
	"eklase/state"
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// MainMenu defines a main menu screen layout.
func MainMenu(th *material.Theme, state *state.State) Screen {
	var (
		add  widget.Clickable
		list widget.Clickable
		quit widget.Clickable
	)
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		matAddBut := material.Button(th, &add, "Add student")
		matAddBut.Font = text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
		matAddBut.Background = color.NRGBA{A: 0xff, R: 0x1e, G: 0x4d, B: 0x24}
		matListBut := material.Button(th, &list, "List students")
		matListBut.Font = text.Font{Variant: "Mono", Weight: text.Bold, Style: text.Italic}
		matQuitBut := material.Button(th, &quit, "Quit")
		matQuitBut.Font = text.Font{Variant: "Smallcaps", Style: text.Italic}

		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(matAddBut.Layout)),
			layout.Rigid(rowInset(matListBut.Layout)),
			layout.Rigid(rowInset(matQuitBut.Layout)),
		)
		if add.Clicked() {
			return AddStudent(th, state), d
		}
		if list.Clicked() {
			return ListTable(th, state), d
		}
		if quit.Clicked() {
			state.Quit()
		}
		return nil, d
	}
}
