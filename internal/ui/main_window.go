package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var requests = []string{}

func NewMainWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Request Buddy")

	leftNavigation := createNavigation()

	method := widget.NewSelect([]string{"HEAD", "GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"}, func(s string) {})
	method.Selected = "GET"
	url := widget.NewEntry()
	send := widget.NewButton("Send", func() {})
	addressBar := container.NewBorder(nil, nil, method, send, url)

	content := container.NewAppTabs(
		container.NewTabItem("Request", widget.NewLabel("Requst content")),
		container.NewTabItem("Response", widget.NewLabel("Response content")),
	)

	mainContent := container.NewBorder(
		addressBar, nil, nil, nil, content)

	split := container.NewHSplit(leftNavigation, mainContent)
	split.Offset = 0.2

	w.SetContent(split)
	w.Resize(fyne.NewSize(800, 600))
	return w
}

func createNavigation() fyne.CanvasObject {
	requestList := widget.NewList(
		func() int { return len(requests) },
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(requests[i])
		})

	return container.NewBorder(
		widget.NewButton("New Request", func() {
			requests = append(requests, "New Request")
			requestList.Refresh()
		}),
		nil,
		nil,
		nil,
		requestList,
	)
}
