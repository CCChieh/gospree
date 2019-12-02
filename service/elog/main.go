package elog

import (
	"github.com/ccchieh/gospree/util"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"time"
)

var (
	Logger *logger
)

func init() {
	Logger = new(logger)
}

//Maintain log.
//Delete outdated log.
func (log *logger) Maintain(logSaveDays int) {
	if !log.saveFlag {
		log.Error("maintain error because of logger don't to save as a file")
		return
	}
	log.wg.Add(1)
	defer log.wg.Done()

	log.Info("Start maintain log.")
	log.mu.Lock()
	defer log.mu.Unlock()

	//set file offset.
	_, err := log.file.Seek(0, 0)
	Panic(errors.Wrap(err, "occur error while seek log file"))

	all, err := ioutil.ReadAll(log.file)
	Panic(errors.Wrap(err, "occur error while read log file"))

	//set file offset.
	_, err = log.file.Seek(0, 0)
	Panic(errors.Wrap(err, "occur error while seek log file"))
	initTime := time.Time{}
	var ls int

	for i := 0; ; i++ {
		line, lt, err := util.GetLineInByte(all, ls)
		if err != nil {
			//EOF.
			if i != 0 {
				//all maintain
				log.closeFile()
				log.openFile(os.O_TRUNC | os.O_CREATE | os.O_APPEND)
			}
			return
		}
		t := util.GetJsonTime(line)

		if i == 0 && t == initTime {
			//blank file.
			return
		}

		if time.Now().Sub(t).Hours()/24 <= float64(logSaveDays) {
			if i == 0 {
				//no maintain required.
				return
			}
			log.closeFile()
			log.openFile(os.O_TRUNC | os.O_CREATE | os.O_APPEND)
			//update log file.
			_, err = log.file.Write(all[ls:])
			Panic(errors.Wrap(err, "occur error while write log file"))
			break
		}
		ls = lt
	}

}

func (log *logger) Debug(args ...interface{}) {
	log.log(DebugLevel, args...)
}

func (log *logger) Info(args ...interface{}) {
	log.log(InfoLevel, args...)
}

func (log *logger) Warn(args ...interface{}) {
	log.log(WarnLevel, args...)
}

func (log *logger) Error(args ...interface{}) {
	log.log(ErrorLevel, args...)
}

func (log *logger) Panic(args ...interface{}) {
	log.log(PanicLevel, args...)
}

func (log *logger) Err(err error) {
	if err == nil {
		return
	}
	switch t := err.(type) {
	case Err:
		log.logErr(t)
	default:
		err := newErr()
		err.Msg = t.Error()
		err.Level = ErrorLevel
		log.logErr(err)
	}
}

//Initialization logger.
func (log *logger) SetSave(saveFlag bool, path string) {
	if saveFlag {
		log.mu.Lock()
		defer log.mu.Unlock()
		log.path = path
		log.openFile(os.O_CREATE | os.O_APPEND)
	} else if log.saveFlag {
		log.destroy()
	}
	log.saveFlag = saveFlag
	//err := core.Cron.AddFunc("@daily", log.Maintain)
	//Panic(errors.Wrap(err, "fail to add function in cron"))
}
