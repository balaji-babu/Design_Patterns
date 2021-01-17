package main

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type BookInfo struct{
	ISBN string `json:"ISBN"`
	Name string `json:"Name"`
}

type BookEntries struct{
	entries []BookInfo
}

type Persistence struct {

}

func (be *BookEntries) AddEntry(bi *BookInfo) {
	be.entries = append(be.entries, *bi)
	time.Sleep(10)
}

func (be *BookEntries) ListEntry() {

}

func (be *BookEntries) RemoveEntry(isbn string) {

}

func (p *Persistence) SaveToFile(be *BookEntries) {

	file, _ := json.Marshal(be.entries)
	_ = ioutil.WriteFile("/tmp/test.txt", file, 0644)
}

func main() {
	info1 := BookInfo{"978-0470138137","Common Sense on Mutual Funds"}
	info2 := BookInfo{"978-0743200400","One Up On Wall Street"}
	entry := BookEntries{}
	entry.AddEntry(&info1)
	entry.AddEntry(&info2)

	per := Persistence{}
	per.SaveToFile(&entry)
}