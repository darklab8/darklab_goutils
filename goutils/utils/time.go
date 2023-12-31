package utils

import (
	"fmt"
	"time"

	utils_logus1 "github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_logus"
)

type timeMeasurer struct {
	msg          string
	ops          []utils_logus1.SlogParam
	time_started time.Time
}

func NewTimeMeasure(msg string, ops ...utils_logus1.SlogParam) *timeMeasurer {
	return &timeMeasurer{
		msg:          msg,
		ops:          ops,
		time_started: time.Now(),
	}
}

func (t *timeMeasurer) Close() {
	utils_logus.Log.Debug(fmt.Sprintf("time_measure %v | %s", time.Since(t.time_started), t.msg), t.ops...)
}

func TimeMeasure(callback func(), msg string, ops ...utils_logus1.SlogParam) {
	time_started := NewTimeMeasure(msg, ops...)
	defer time_started.Close()
	callback()
}
