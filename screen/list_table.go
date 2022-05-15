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

func generateStudentsList(th *material.Theme, list widget.List, students []storage.StudentEntry, delete []widget.Clickable, edit []widget.Clickable) func(gtx layout.Context) layout.Dimensions {
	lightContrast := th.ContrastBg
	lightContrast.A = 0x11
	darkContrast := th.ContrastBg
	darkContrast.A = 0x33
	return func(gtx layout.Context) layout.Dimensions {
		return material.List(th, &list).Layout(gtx, len(students), func(gtx layout.Context, index int) layout.Dimensions {
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
						layout.Rigid(rowInset(material.Body1(th, fmt.Sprintf("%s %s  ", student.Surname, student.Name)).Layout)),
						layout.Rigid(material.Button(th, &delete[index], "Delete").Layout),
					)
				})),
				layout.Stacked(rowInset(func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{}.Layout(gtx,
						layout.Rigid(rowInset(material.Body1(th, fmt.Sprintf("%s %s                   ", student.Surname, student.Name)).Layout)),
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

	students, err := state.Students()
	if err != nil {
		// TODO: Show user an error toast.
		log.Printf("failed to fetch students: %v", err)
		return nil
	}
	// var maxlength storage.StudentEntry                // DELETE AFTER !!!
	// for _, i := range students {
	//	   fmt.Println("i:", i)
	// }
	delete := make([]widget.Clickable, len(students))
	edit := make([]widget.Clickable, len(students))

	studentsLayout := generateStudentsList(th, list, students, delete, edit)

	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
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
				state.DeleteRecordByID(students[i].ID)
				return AddStudent(th, state), d // it works, but it doesn't replace the student in the same place, it adds him below
			}
		}
		students, err = state.Students()
		if err != nil {
			// TODO: Show user an error toast.
			log.Printf("failed to fetch students: %v", err)
			os.Exit(1)
		}
		studentsLayout = generateStudentsList(th, list, students, delete, edit)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		return nil, d
	}
}
