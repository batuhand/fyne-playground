package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/storage"
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

func attackerScreen(win fyne.Window) fyne.CanvasObject {
	progress := widget.NewProgressBar()
	attackButton1 := widget.NewButton("Select File and Start Attack", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(r)
				testlist := strings.Split(string(data), "\n")
				for i := 0; i < len(testlist); i++ {
					print(testlist[i])
				}
			}, win)
		file_Dialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))
		file_Dialog.Show()
	})
	hashType := widget.NewSelect([]string{"Hash Type 1", "Hash Type 2"}, func(value string) {
		log.Println("Select set to", value)
	})
	sysInfoVal := false
	summaryVal := false
	sysInfoCheck := widget.NewCheck("System Info", func(value bool) {
		sysInfoVal = value
		fmt.Println(sysInfoVal)
	})
	summaryCheck := widget.NewCheck("Summary CSV File", func(value bool) {
		summaryVal = value
		fmt.Println(summaryVal)
	})
	attackScreen := container.NewCenter(container.NewVBox(
		sysInfoCheck,
		summaryCheck,
		hashType,
		attackButton1,
		progress,
	))

	return attackScreen
}

func scannerScreen(win fyne.Window) fyne.CanvasObject {
	progress := widget.NewProgressBar()
	scanScreen := container.NewCenter(container.NewVBox(
		widget.NewLabelWithStyle("Scan Types", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewButtonWithIcon("Scan with info.txt", theme.SearchIcon(), func() { scanTap(win, progress) }),
		widget.NewButtonWithIcon("Scan with servers.txt and credentials.txt", theme.SearchIcon(), func() { scanTap(win, progress) }),
		progress,
	))
	return scanScreen
}

func settingsScreen(win fyne.Window) fyne.CanvasObject {
	dialTimeoutField := widget.NewEntry()
	connectionTimeoutField := widget.NewEntry()
	threadCountField := widget.NewEntry()
	settingScreen := container.NewCenter(
		container.NewVBox(
			container.NewVBox(
				widget.NewLabel("Thread Count"),
				threadCountField,
			),
			container.NewHBox(
				container.NewVBox(
					widget.NewLabel("Dial Timeout"),
					dialTimeoutField,
				),
				container.NewVBox(
					widget.NewLabel("Connection Timeout"),
					connectionTimeoutField,
				),
			),
		),
	)
	return settingScreen
}
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("BBruter")
	myWindow.CenterOnScreen()
	myWindow.Resize(fyne.NewSize(600, 600))
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), homeScreen()),
		container.NewTabItemWithIcon("Scanner", theme.SearchIcon(), scannerScreen(myWindow)),
		container.NewTabItemWithIcon("Attacker", theme.ComputerIcon(), attackerScreen(myWindow)),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), settingsScreen(myWindow)),
	)
	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
