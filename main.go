package main

import (
	"fmt"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func homeScreen() fyne.CanvasObject {
	return container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Welcome to the Fyne toolkit demo app", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container.NewHBox(
			widget.NewLabel("-"),
			widget.NewLabel("-"),
		),
		widget.NewLabel(""), // balance the header on the tutorial screen we leave blank on this content
	))
}
func scanTap(win fyne.Window, progress *widget.ProgressBar) {
	fmt.Println("Scan started")
	dialog.ShowInformation("Loading", "Scan started", win)
	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Millisecond * 250)
			progress.SetValue(i)

		}
	}()

}

func scannerScreen(win fyne.Window) fyne.CanvasObject {
	progress := widget.NewProgressBar()
	scanScreen := container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("info.txt", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewButtonWithIcon("Start Scan", theme.SearchIcon(), func() { scanTap(win, progress) }),
		progress,
	))
	return scanScreen
}
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("BBruter")
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(600, 600))
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), homeScreen()),
		container.NewTabItemWithIcon("Scanner", theme.SearchIcon(), scannerScreen(myWindow)),
		container.NewTabItemWithIcon("Attacker", theme.ComputerIcon(), widget.NewLabel("Attacker tab")),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("Settings tab")),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
