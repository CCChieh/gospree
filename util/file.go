package util

import "io"

//Get a line in []byte.
func GetLineInByte(text []byte, s int) (line []byte, i int, err error) {
	l := len(text)
	if s >= l {
		err = io.EOF
		return
	}
	for i := s; i < l; i++ {
		if text[i] == '\n' {
			return text[s : i+1], i + 1, err
		}
	}
	return text[s:], l, err
}
