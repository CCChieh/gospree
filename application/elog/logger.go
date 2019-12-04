package elog

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"os"
	"runtime"
	"sync"
	"time"
)

type logger struct {
	mu       sync.Mutex
	wg       sync.WaitGroup
	file     *os.File
	oldest   time.Time
	level    int
	path     string
	saveFlag bool
}

//destroy log.
//Wait() is used to ensure that log output is complete.
func (log *logger) destroy() {
	log.wg.Wait()
	log.closeFile()
}

//Open log file.
func (log *logger) openFile(mode int) {
	var logfile *os.File
	var err error
	logfile, err = os.OpenFile(log.path, mode, 0666)
	Panic(errors.Wrap(err, "occur error while Opening log file"))
	log.file = logfile
}

func (log *logger) closeFile() {
	fmt.Println(log.file.Close())
}

//Return level string
func (log *logger) levelText(level int) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO "
	case WarnLevel:
		return "WARN "
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	}
	return ""
}

//Output log
func (log *logger) log(levelInt int, args ...interface{}) {
	log.wg.Add(1)
	defer log.wg.Done()
	//defer log.Done()
	if levelInt < log.level {
		return
	}

	level := log.levelText(levelInt)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	//content saved log data
	content := &struct {
		Level   string
		Time    time.Time
		File    string
		Line    int
		Message string
	}{
		level,
		time.Now(),
		file,
		line,
		fmt.Sprint(args...)}
	log.output(levelInt, content)
}

//logErr use to output Err Type
func (log *logger) logErr(err Err) {
	log.wg.Add(1)
	defer log.wg.Done()
	//defer log.Done()
	if err.GetLevel() < log.level {
		return
	}

	level := log.levelText(err.GetLevel())
	file, line := err.Local()
	//content saved log data
	content := &struct {
		Level   string
		Time    time.Time
		File    string
		Line    int
		Message string
	}{
		level,
		time.Now(),
		file,
		line,
		err.Error()}
	log.output(err.GetLevel(), content)
}

//output.
func (log *logger) output(levelInt int, content *struct {
	Level   string
	Time    time.Time
	File    string
	Line    int
	Message string
}) {
	log.wg.Add(1)
	defer log.wg.Done()
	colorPrint := func(format string, a ...interface{}) {}
	switch levelInt {
	case DebugLevel:
		colorPrint = color.Blue
	case InfoLevel:
		colorPrint = color.Cyan
	case WarnLevel:
		colorPrint = color.Yellow
	case ErrorLevel:
		colorPrint = color.Red
	case PanicLevel:
		colorPrint = color.Magenta
	}

	log.mu.Lock()
	defer log.mu.Unlock()

	colorPrint("%s | %s | %s:%d | %s\n",
		content.Level,
		content.Time.Format("2006-01-02 15:04:05"),
		content.File,
		content.Line,
		content.Message)
	js, err := json.Marshal(content)
	Panic(errors.Wrap(err, "occur error while convert to json"))

	if log.saveFlag {
		_, err = log.file.Write(append(js, '\n'))
		Panic(errors.Wrap(err, "occur error while json write to file %s"))
	}

}
