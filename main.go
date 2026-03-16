package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Temperature Converter")

	// 1. Setup the Celsius display (Above slider)
	cLabel := widget.NewRichTextFromMarkdown("# 0°C")

	// 2. Setup the Fahrenheit display (Result)
	fLabel := widget.NewRichTextFromMarkdown("# 32.0°F")

	slider := widget.NewSlider(-50, 50)

	slider.OnChanged = func(celsius float64) {
		// Update the Celsius label above the slider
		cLabel.ParseMarkdown(fmt.Sprintf("# %.0f°C", celsius))

		// Calculate and update the large Fahrenheit display
		fahrenheit := (celsius * 9 / 5) + 32
		fLabel.ParseMarkdown(fmt.Sprintf("# %.1f°F", fahrenheit))
	}

	resetButton := widget.NewButton("Reset to 0°C", func() {
		slider.SetValue(0)
	})

	// 3. Arrange layout: Celsius label -> Slider -> Fahrenheit Result
	content := container.NewVBox(
		container.NewCenter(cLabel),
		slider,
		container.NewCenter(fLabel),
		resetButton,
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 190))
	myWindow.ShowAndRun()
}
