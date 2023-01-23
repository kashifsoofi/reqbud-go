package ui

import (
	"io"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var requests = []string{}

func NewMainWindow(a fyne.App) fyne.Window {
	w := a.NewWindow("Request Buddy")

	leftNavigation := createNavigation()

	responseContent := widget.NewLabel("Response content")
	content := container.NewAppTabs(
		container.NewTabItem("Request", widget.NewLabel("Requst content")),
		container.NewTabItem("Response", responseContent),
	)

	method := widget.NewSelect([]string{"HEAD", "GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"}, func(s string) {})
	method.Selected = "GET"
	url := widget.NewEntry()
	send := widget.NewButton("Send", func() {
		responseText := makeRequest(method.Selected, url.Text, "")
		responseContent.SetText(responseText)
		responseContent.Refresh()
	})
	addressBar := container.NewBorder(nil, nil, method, send, url)

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

func createRequest(method, url, data string) (*http.Request, error) {
	var body io.Reader
	if len(data) > 0 {
		body = strings.NewReader(data)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func makeRequest(method, url, data string) string {
	req, err := createRequest(method, url, data)
	if err != nil {
		return "Error creating request"
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Error making request"
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response body"
	}

	return string(body)
}
