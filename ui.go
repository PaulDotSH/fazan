package main

import (
	"embed"
	"fazan/gamemodes"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

////go:embed bundle/filenames
//var filenames string

//go:embed bundle/cuvinte/*
var content embed.FS

var a fyne.App
var w fyne.Window
var gameModeSelect *widget.Select
var cuvantLabel *widget.Label
var cuvantEntry *widget.Entry

var nustiuButton *widget.Button
var verifyButton *widget.Button
var emptyWidgetLabel *widget.Label

var mainScreenContainer *fyne.Container

var CuvantBot = binding.NewString()
var CuvantUtilizator = binding.NewString()

var IsWordValid *bool

var PuncteUtilizator = binding.NewInt()
var PuncteBot = binding.NewInt()

var PuncteBotLabel *widget.Label
var PuncteUtilizatorLabel *widget.Label

//fa alt folder cu gamemodes si in el o sa ai words
//in gamemodes o sa pasezi cumva bindings la cuvantgenerat si cuvantul de la user

func init() {
	b := false
	IsWordValid = &b
	a = app.NewWithID("words.paul.abrudan")

	w = a.NewWindow("Fazan")
	w.SetTitle("Fazan")
	emptyWidgetLabel = widget.NewLabel("")

	w.Resize(fyne.Size{800, 600})

	cuvantLabel = widget.NewLabelWithData(CuvantBot)
	cuvantLabel.Alignment = fyne.TextAlignCenter

	cuvantEntry = widget.NewEntryWithData(CuvantUtilizator)
	cuvantEntry.PlaceHolder = "Introdu cuvantul tau aici"

	PuncteBotLabel = widget.NewLabelWithData(binding.IntToStringWithFormat(PuncteBot, "Punctele botului: %v"))
	PuncteUtilizatorLabel = widget.NewLabelWithData(binding.IntToStringWithFormat(PuncteUtilizator, "Punctele tale: %v"))

	joacaBtn := widget.NewButton("Joaca", joacaBtnClick)

	//inapoi button o sa dea w.SetContent(b) or smth

	despreBtn := widget.NewButton("Despre", func() {
		dialog.ShowInformation("Despre Fazan", "Aceasta aplicatie incearca sa implementeze jocul de \"fazan\",\nmomentan jocul poate avea unele probleme,\nputeti raporta problemele pe care le aveti\npe github.com/bvckdoor/fazan/issues\nAplicatia a fost facuta in GO cu ajutorul librariei fyne.\nDaca va place interfata ii puteti dona lui \"andy.xyz\"\nIn modul de joc \"usor\" nu se verifica cuvintele duplicate\nNu uitati sa dati 5 stele aplicatiei :D", w)
	})

	titleLabel := widget.NewLabel("FAZAN")
	titleLabel.Alignment = fyne.TextAlignCenter

	//Daca e usor, fara cuvinte care nu se pot repeta, normal e normal
	gameModeSelect = widget.NewSelect([]string{"Usor"}, nil)
	gameModeSelect.SetSelectedIndex(0)

	gameModeSelect.PlaceHolder = "Selecteaza un mod de joc"

	//dumb way, don't know of a better way tho
	mainScreenContainer = container.NewGridWithColumns(1,
		titleLabel, emptyWidgetLabel, gameModeSelect, joacaBtn, emptyWidgetLabel, emptyWidgetLabel, emptyWidgetLabel, despreBtn, emptyWidgetLabel)

	w.SetContent(mainScreenContainer)

}

func joacaBtnClick() {
	var idkFunc func()
	var verifyFunc func()

	switch gameModeSelect.SelectedIndex() {
	case 0:
		gamemodes.Easy_init(content, CuvantUtilizator, CuvantBot, IsWordValid, PuncteUtilizator, PuncteBot)
		//gamemodes.Easy_init(content, CuvantUtilizator, CuvantBot, IsWordValid, PuncteUtilizator, PuncteBot)
		idkFunc, verifyFunc = gamemodes.Easy_handleIdk, gamemodes.Easy_verifyCuvant
	default:
		return
	}

	nustiuButton = widget.NewButton("Nu stiu", idkFunc)

	verifyButton = widget.NewButton("Verifica cuvantul", func() {
		verifyFunc()
		cuvantEntry.FocusGained()
	})

	w.SetContent(container.NewGridWithColumns(1,
		emptyWidgetLabel, cuvantLabel, PuncteUtilizatorLabel, PuncteBotLabel, cuvantEntry, verifyButton, nustiuButton, emptyWidgetLabel, emptyWidgetLabel))
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
