package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount, text)

	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// separation of concerns

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	//
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	//
}

var LineSeparator = "\n"

func SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, LineSeparator)), 0644)
}

type Persistence struct {
	LineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename,
		[]byte(strings.Join(j.entries, p.LineSeparator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I cried today")
	j.AddEntry("I ate an apple")

	fmt.Println(j.String())

	//
	SaveToFile(&j, "journal.txt")

	//
	p := Persistence{"\r\n"}
	p.SaveToFile(&j, "journal.txt")
}
