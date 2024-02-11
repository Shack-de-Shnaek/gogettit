package main

import (
	"fmt"
	"gogettit/modules/desktop"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	app := app.New()

	mainWidth := 768
	mainHeight := 456
	defaultSize := fyne.NewSize(float32(mainWidth), float32(mainHeight))

	window := app.NewWindow("GoGettIt")
	window.Resize(defaultSize)
	window.SetFixedSize(true)
	window.SetPadded(false)

	searchIcon := widget.NewIcon(theme.SearchIcon())
	searchIcon.Resize(fyne.NewSquareSize(32))
	searchIconContainer := container.NewVBox(container.NewPadded(searchIcon))
	searchIconContainer.Resize(fyne.NewSquareSize(32))

	searchBox := widget.NewEntry()
	searchBox.SetPlaceHolder("Search something...")
	searchBoxContainer := container.NewVBox(searchBox)

	searchContainer := container.NewBorder(nil, nil, searchIconContainer, nil, searchBoxContainer)
	searchContainer.Resize(fyne.NewSize(768, 32))

	mainContainer := container.NewGridWithRows(1, searchContainer)
	mainContainer.Resize(fyne.NewSize(768,  32))

	searchBox.OnChanged = func(bingus string) {
		// fmt.Println("Key pressed")
		// fmt.Println(bingus)
		files, err := desktop.GetDesktopFiles()
		if err != nil {
			fmt.Println(err)
		} else {
			for _, file := range(files) {
				fmt.Println(file.Name())
			}
		}
	}

	window.SetContent(mainContainer)
	window.ShowAndRun()
}
