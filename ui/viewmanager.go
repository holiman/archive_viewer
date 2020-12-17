// This file is part of the goevmlab library.
//
// The library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the goevmlab library. If not, see <http://www.gnu.org/licenses/>.

package ui

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/holiman/archive_viewer/model"
	"github.com/rivo/tview"
)

const (
	headingCol = tcell.ColorYellow
)

type viewManager struct {
	room *model.Room

	msgView    *tview.Table
	detailView *tview.TextView
	root       *tview.Grid
}

func NewViewManager(room *model.Room) *viewManager {

	msgView := tview.NewTable()
	msgView.SetTitle("Messages").SetBorder(true)

	detailView := tview.NewTextView()

	root := tview.NewGrid().
		SetRows(0, 15).
		SetColumns(0).
		SetBorders(true)

	mgr := viewManager{
		room:       room,
		msgView:    msgView,
		root:       root,
		detailView: detailView,
	}

	var search = func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'f' {
			row, _ := mgr.msgView.GetSelection()
			row++
			for ; row < len(room.Messages); row++ {
				if txt := room.Messages[row].Text; strings.Index(txt, "consensus") > 0 {
					mgr.detailView.SetText(txt)
					mgr.msgView.Select(row+1, 0)
					break
				}
			}
			return nil
		}
		return event
	}

	mgr.init(room)
	msgView.SetInputCapture(search)

	root.
		AddItem(msgView, 0, 0, 1, 1, 0, 50, true).
		AddItem(detailView, 1, 0, 1, 1, 0, 50, true)

	return &mgr
}

// Starts the UI compoments
func (mgr *viewManager) Run() {
	if err := tview.NewApplication().SetRoot(mgr.root, true).Run(); err != nil {
		panic(err)
	}
}
func setHeadings(headings []string, table *tview.Table) {

	table.SetFixed(1, 0).SetSelectable(false, false)
	for col, title := range headings {
		table.SetCell(0, col,
			tview.NewTableCell(strings.ToUpper(title)).
				SetTextColor(headingCol).
				SetAlign(tview.AlignRight))
	}
}
func (mgr *viewManager) init(room *model.Room) {
	{ // The operations table
		table := mgr.msgView
		headings := []string{"time", "from", "message"}

		table.SetSelectable(true, false).
			SetSelectionChangedFunc(func(row, col int) {
				r := row - 1
				txt := ""
				if r > 0 && r < len(mgr.room.Messages) {
					txt = mgr.room.Messages[r].Text
				}
				mgr.detailView.SetText(txt)
			}).
			Select(1, 1).SetFixed(1, 1)

		// Headings
		for col, title := range headings {
			table.SetCell(0, col,
				tview.NewTableCell(strings.ToUpper(title)).
					SetTextColor(headingCol).
					SetAlign(tview.AlignCenter))
		}
		// Ops table body
		for i, elem := range room.Messages {
			if elem == nil {
				break
			}
			row := i + 1
			table.SetCell(row, 0, tview.NewTableCell(elem.Sent.Format("2006-01-02")))
			table.SetCell(row, 1, tview.NewTableCell(elem.From.UserName))
			table.SetCell(row, 2, tview.NewTableCell(elem.Text))
		}
	}
}
