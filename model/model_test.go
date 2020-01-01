package model

import (
	"fmt"
	"testing"
)

func TestNote(t *testing.T) {
	_ = Connection("gospree:wdsrshi19971025@(lab.nimingshe.com:3306)/gospree?parseTime=True&loc=Local")
	notes := new(Notes)
	fmt.Println(notes.GetNoteList(0, 2))
	fmt.Println(notes.GetNoteList(1, 2))
	fmt.Println(notes.GetNoteList(2, 2))
	fmt.Println(notes.GetNoteList(3, 2))
	fmt.Println(notes.GetNoteList(4, 2))
	fmt.Println(notes.GetNoteList(5, 2))
	fmt.Println(notes.GetNoteList(6, 2))
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

func TestNote_CreateNote(t *testing.T) {

}
