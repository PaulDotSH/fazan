package main

import (
	"embed"
	_ "embed"
	"fazan/words"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var a fyne.App
var w fyne.Window

////go:embed bundle/filenames
//var filenames string

//go:embed bundle/cuvinte/*
var content embed.FS

func init() {

	a = app.NewWithID("words.paul.abrudan")
	w = a.NewWindow("Fazan")

	w.Resize(fyne.Size{800, 600})

	joacaBtn := widget.NewButton("Joaca", func() {
		w.SetContent(container.NewGridWithColumns(1,
			widget.NewLabel("Cuvantul generat de "), nil, nil, nil, nil))
	})

	//inapoi button o sa dea w.SetContent(b) or smth

	despreBtn := widget.NewButton("Despre", func() {
		dialog.ShowInformation("Despre Fazan", "Aceasta aplicatie incearca sa implementeze jocul de \"fazan\",\nmomentan jocul poate avea unele probleme,\nputeti raporta problemele pe care le aveti pe github.com/bvckdoor/fazan/issues\nAplicatia a fost facuta in GO cu ajutorul librariei fyne, daca va place interfata ii puteti dona lui \"andy.xyz\"\nNu uitati sa dati 5 stele aplicatiei :D", w)
	})

	box1 := widget.NewLabel("")

	titleLabel := widget.NewLabel("FAZAN")
	titleLabel.Alignment = fyne.TextAlignCenter

	//dumb way, don't know of a better way tho
	b := container.NewGridWithColumns(1,
		titleLabel, joacaBtn, box1, despreBtn, box1)

	w.SetContent(b)
	words.Initialise(&content)
}

func DisplayNotification(title, content string) {
	fyne.CurrentApp().SendNotification(&fyne.Notification{
		Title:   title,
		Content: content,
	})
}

func main() {
	w.ShowAndRun()
}
