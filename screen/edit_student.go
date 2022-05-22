package screen

import (
	"eklase/state"
	"fmt"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// AddStudent defines a screen layout for adding a new student.
func EditStudent(th *material.Theme, state *state.State, id int, nameT string, surnameT string) Screen {
	var (
		name    widget.Editor
		surname widget.Editor

		close widget.Clickable
		save  widget.Clickable
	)
	enabledIfNameOK := func(w layout.Widget) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			name := strings.TrimSpace(name.Text())
			surname := strings.TrimSpace(surname.Text())
			if name == "" && surname == "" || name == nameT && surname == surnameT {
				gtx = gtx.Disabled()
			}
			return w(gtx)
		}
	}
	editsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(1, material.Editor(th, &name, "First name").Layout),
			layout.Rigid(spacer.Layout),
			layout.Flexed(1, material.Editor(th, &surname, "Last name").Layout),
		)
	}
	buttonsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceStart}.Layout(gtx,
			layout.Rigid(material.Button(th, &close, "Close").Layout),
			layout.Rigid(spacer.Layout),
			layout.Rigid(enabledIfNameOK(material.Button(th, &save, "Save").Layout)),
		)
	}
	name.SetText(nameT)
	surname.SetText(surnameT)
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(editsRowLayout)),
			layout.Rigid(rowInset(buttonsRowLayout)),
		)
		if close.Clicked() {
			return ListTable(th, state), d
		}
		if save.Clicked() {
			fmt.Println(state.EditRecordByID(id, name.Text(), surname.Text()))
			return ListTable(th, state), d
		}
		return nil, d
	}
}
