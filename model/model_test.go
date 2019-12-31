package model

import (
	"fmt"
	"testing"
)

func TestNote(t *testing.T) {
	_ = Connection("gospree:wdsrshi19971025@(lab.nimingshe.com:3306)/gospree?parseTime=True&loc=Local")
	notes := new(Notes)
	fmt.Println(notes.GetNoteIDList(0, 2))
	fmt.Println(notes.GetNoteIDList(1, 2))
	fmt.Println(notes.GetNoteIDList(2, 2))
	fmt.Println(notes.GetNoteIDList(3, 2))
	fmt.Println(notes.GetNoteIDList(4, 2))
	fmt.Println(notes.GetNoteIDList(5, 2))
	fmt.Println(notes.GetNoteIDList(6, 2))
}

type N struct {
	name string
	age  int
	St   Interface
}
type Interface interface {
}

func TestSt(t *testing.T) {
	n := N{}
	fmt.Println(n)
	fmt.Println(n.St == nil)
}
