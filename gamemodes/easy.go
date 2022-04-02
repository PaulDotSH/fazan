package gamemodes

import (
	"embed"
	"fazan/gamemodes/words"
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"strings"
)

var cuvantUtilizator binding.String
var cuvantBot binding.String
var isWordValid *bool
var pctUser binding.Int
var pctBot binding.Int

func Easy_init(files embed.FS, cuvUser, cuvBot binding.String, IsWordValid *bool, PctUser, PctBot binding.Int) {
	words.Initialise(&files)
	cuvantUtilizator = cuvUser
	cuvantBot = cuvBot
	isWordValid = IsWordValid
	pctUser = PctUser
	pctBot = PctBot
	//puncteUtilizator = PuncteUtilizator
	//puncteBot = PuncteBot

	s, _ := words.GetWordStartingWith("")
	cuvBot.Set(s)
}

func getLast2Characters(word string) string {
	var sb strings.Builder
	length := len(word)
	sb.WriteByte(word[length-2])
	sb.WriteByte(word[length-1])
	return sb.String()
}

func Easy_verifyCuvant() {
	cuvUserStr, _ := cuvantUtilizator.Get()
	if cuvUserStr == "" {
		return
	}
	fmt.Println(cuvUserStr)

	cuvBotStr, _ := cuvantBot.Get()

	//check string prefix
	if !strings.HasPrefix(cuvUserStr, getLast2Characters(cuvBotStr)) {
		return
	}

	*isWordValid = words.IsValidWord(cuvUserStr)
	if *isWordValid == false {
		return
	}

	//delete user's word
	cuvantUtilizator.Set("")

	//get next word
	GetNewWord(getLast2Characters(cuvUserStr))
}

func AddPctUser() {
	i, e := pctUser.Get()
	if e != nil {
		i = 0
	}
	pctUser.Set(i + 1)
	GetNewWord("")
}

func Easy_handleIdk() {
	i, e := pctBot.Get()
	if e != nil {
		i = 0
	}
	pctBot.Set(i + 1)
	GetNewWord("")
}

func GetNewWord(start string) {
	s, e := words.GetWordStartingWith(start)
	if e != nil {
		AddPctUser()
		return
	}
	cuvantBot.Set(s)
}
