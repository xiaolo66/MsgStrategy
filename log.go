package MsgStrategy

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"io/ioutil"
	"path"
	"time"
)

var Log = logrus.New()

func InitLog(filePath, filename string, maxAge, rotationTime time.Duration) error {
	var writeMap = make(lfshook.WriterMap)
	for _, level := range logrus.AllLevels {
		fullPath := path.Join(filePath, filename, level.String())
		writer, err := rotatelogs.New(
			fullPath+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(fullPath),         //
			rotatelogs.WithMaxAge(maxAge),             //
			rotatelogs.WithRotationTime(rotationTime), //
		)
		if err == nil {
			writeMap[level] = writer
		}
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{DisableColors: true})
	Log.AddHook(lfHook)
	Log.Out = ioutil.Discard

	return nil
}
