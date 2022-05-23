package screen

import (
	"eklase/state"
	"eklase/storage"
	"fmt"
	"image"
	"log"
	"os"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func generateStudentsList(th *material.Theme, list *widget.List, students []storage.StudentEntry, delete []widget.Clickable, edit []widget.Clickable, name widget.Editor, surname widget.Editor) func(gtx layout.Context) layout.Dimensions {
	lightContrast := th.ContrastBg
	lightContrast.A = 0x11
	darkContrast := th.ContrastBg
	darkContrast.A = 0x33

	return func(gtx layout.Context) layout.Dimensions {
		return material.List(th, list).Layout(gtx, len(students), func(gtx layout.Context, index int) layout.Dimensions {
			student := students[index]

			return layout.Stack{}.Layout(gtx,
				layout.Expanded(func(gtx layout.Context) layout.Dimensions {
					color := lightContrast
					if index%2 == 0 {
						color = darkContrast
					}

					max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
					paint.FillShape(gtx.Ops, color, clip.Rect{Max: max}.Op())
					return layout.Dimensions{Size: gtx.Constraints.Min}
				}),
				layout.Stacked(rowInset(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(rowInset(material.Body1(th, fmt.Sprintf("%s %s", student.Surname, student.Name)).Layout)),
					)
				})),
				layout.Stacked(rowInset(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(rowInset(material.Body1(th, fmt.Sprintln("                                                                                                                                              ")).Layout)),
						layout.Rigid(material.Button(th, &delete[index], "Delete").Layout),
					)
				})),
				layout.Stacked(rowInset(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(rowInset(material.Body1(th, fmt.Sprintf("%s %s", student.Surname, student.Name)).Layout)),
					)
				})),
				layout.Stacked(rowInset(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(rowInset(material.Body1(th, fmt.Sprintln("                                                                                                                                                              ")).Layout)),
						layout.Rigid(material.Button(th, &edit[index], "Edit").Layout),
					)
				})),
			)
		})
	}
}

// ListStudent defines a screen layout for listing existing students.
func ListTable(th *material.Theme, state *state.State) Screen {
	var close widget.Clickable
	list := widget.List{List: layout.List{Axis: layout.Vertical}}
	var name, surname widget.Editor
	students, err := state.Students(name.Text(), surname.Text())
	if err != nil {
		// TODO: Show user an error toast.
		log.Printf("failed to fetch students: %v", err)
		return nil
	}

	delete := make([]widget.Clickable, len(students))
	edit := make([]widget.Clickable, len(students))

	studentsLayout := generateStudentsList(th, &list, students, delete, edit, name, surname)
	editsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(1, material.Editor(th, &name, "First name").Layout),
			layout.Rigid(spacer.Layout),
			layout.Flexed(1, material.Editor(th, &surname, "Last name").Layout),
		)
	}
	return func(gtx layout.Context) (Screen, layout.Dimensions) {

		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(editsRowLayout)),
			layout.Flexed(1, rowInset(studentsLayout)),
			layout.Rigid(rowInset(material.Button(th, &close, "Close").Layout)),
		)

		for i := range delete {
			if delete[i].Clicked() {
				state.DeleteRecordByID(students[i].ID)
			}
		}
		for i := range edit {
			if edit[i].Clicked() {
				return EditStudent(th, state, students[i].ID, students[i].Name, students[i].Surname), d
			}
		}

		students, err = state.Students(name.Text(), surname.Text())
		if err != nil {
			// TODO: Show user an error toast.
			log.Printf("failed to fetch students: %v", err)
			os.Exit(1)
		}
		studentsLayout = generateStudentsList(th, &list, students, delete, edit, name, surname)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		return nil, d
	}
}
