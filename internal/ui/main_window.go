package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var requests = []string{}

func NewMainWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Request Buddy")

	requestList := widget.NewList(
		func() int { return len(requests) },
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(requests[i])
		})
	requestNavigation := container.NewBorder(
		widget.NewButton("New Request", func() {
			requests = append(requests, "New Request")
			requestList.Refresh()
		}),
		nil,
		nil,
		nil,
		requestList,
	)

	requestContent := widget.NewMultiLineEntry()
	requestContent.SetText("Request content go here")
	responseContent := widget.NewLabel("Response content go here")

	mainContent := container.NewBorder(
		nil, nil, nil, nil, container.NewMax(container.NewVSplit(requestContent, responseContent)))

	split := container.NewHSplit(requestNavigation, mainContent)

	w.SetContent(split)
	w.Resize(fyne.NewSize(640, 460))
	return w
}
