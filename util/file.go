package util

import (
	"io"
	"os"
)

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

// 判断所给路径文件/文件夹是否存在
func FileIsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//新建目录
func MkDir(path string) error {
	err := os.Mkdir(path, 0777)
	return err
}
