package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	md "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func main() {
	// Forzar el tema oscuro antes de crear la app
	os.Setenv("FYNE_THEME", "dark")
	myApp := app.New()

	myWindow := myApp.NewWindow("Convertidor de HTML a Markdown")
	myWindow.Resize(fyne.NewSize(900, 600))

	// Cuadros de texto multilínea
	inputEntry := widget.NewMultiLineEntry()
	inputEntry.SetPlaceHolder("Pega aquí el código fuente HTML de la página...")

	outputEntry := widget.NewMultiLineEntry()
	outputEntry.SetPlaceHolder("El Markdown aparecerá aquí...")

	// Botón de conversión
	convertBtn := widget.NewButton("Convertir a Markdown", func() {
		htmlContent := inputEntry.Text

		// Convertir usando la librería
		markdown, err := md.ConvertString(htmlContent)
		if err != nil {
			outputEntry.SetText("Error al convertir: " + err.Error())
			return
		}

		outputEntry.SetText(markdown)
	})

	// Organizar la interfaz en dos columnas (Entrada / Salida)
	// Arriba o en medio ponemos el botón para disparar la acción
	topContainer := container.NewVBox(convertBtn)

	split := container.NewHSplit(
		container.NewBorder(widget.NewLabel("HTML Fuente:"), nil, nil, nil, inputEntry),
		container.NewBorder(widget.NewLabel("Markdown Generado:"), nil, nil, nil, outputEntry),
	)
	split.SetOffset(0.5) // Dividir la pantalla 50% y 50%

	mainLayout := container.NewBorder(topContainer, nil, nil, nil, split)
	myWindow.SetContent(mainLayout)

	myWindow.ShowAndRun()
}
