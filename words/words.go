package words

import (
	"bufio"
	"embed"
	"errors"
	"io/fs"
	"math/rand"
	"strings"
	"time"
)

type cuvinte map[string]byte

var dictionar = make(map[string]cuvinte)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

//call this before doing anything with words
func Initialise(files *embed.FS) {
	//get the text from the files that contain the words

	//get every file from the embedded folder that contains the words we want
	fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		//if folder is . or ..
		length := len(path)
		if length < 3 {
			return nil
		}

		//make a new dict for every file i.e aa
		dict := make(cuvinte)
		file, _ := files.Open(path)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//for line in file dict[line]=1
			dict[scanner.Text()] = 1
		}
		//by this time, we make sure dict["aa"] contains every word that starts with aa

		var builder strings.Builder
		builder.WriteByte(path[length-2])
		builder.WriteByte(path[length-1])
		dictionar[builder.String()] = dict
		return err
	})
}

//do prefix checking in the ui or smth
func IsValidWord(word string) bool {
	length := len(word)
	if length < 2 {
		return false
	}

	var sb strings.Builder
	sb.WriteByte(word[0])
	sb.WriteByte(word[1])

	//dictionar["aa"]["aalenian"]
	return dictionar[sb.String()][word] == 1
}

func GetWordStartingWith(word string) (string, error) {
	length := len(word)
	if length < 2 {
		return "", errors.New("Invalid Length")
	}

	var sb strings.Builder

	//if word == "re" then sb is re
	//if word == "restante" then sb is te
	if length == 2 {
		sb.WriteByte(word[0])
		sb.WriteByte(word[1])
	} else {
		sb.WriteByte(word[length-2])
		sb.WriteByte(word[length-1])
	}

	Cuvinte := dictionar[sb.String()]
	//If there is no word starting with that
	if len(Cuvinte) == 0 {
		return "", errors.New("no word starting with that prefix was found")
	}

	randomIndex := random.Intn(len(Cuvinte))
	i := 0
	for k, _ := range Cuvinte {
		if i == randomIndex {
			return k, nil
		}
		i += 1
	}

	return "", errors.New("no word starting with that prefix was found")
}
