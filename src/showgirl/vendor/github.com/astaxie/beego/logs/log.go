// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logs provide a general log interface
// Usage:
//
// import "github.com/astaxie/beego/logs"
//
//	log := NewLogger(10000)
//	log.SetLogger("console", "")
//
//	> the first params stand for how many channel
//
// Use it like this:
//
//	log.Trace("trace")
//	log.Info("info")
//	log.Warn("warning")
//	log.Debug("debug")
//	log.Critical("critical")
//
//  more docs http://beego.me/docs/module/logs.md
package logs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// RFC5424 log message levels.
const (
	LevelEmergency = iota
	LevelAlert
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
)

const (
	y1  = `0123456789`
	y2  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	y3  = `0000000000111111111122222222223333333333444444444455555555556666666666777777777788888888889999999999`
	y4  = `0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789`
	mo1 = `000000000111`
	mo2 = `123456789012`
	d1  = `0000000001111111111222222222233`
	d2  = `1234567890123456789012345678901`
	h1  = `000000000011111111112222`
	h2  = `012345678901234567890123`
	mi1 = `000000000011111111112222222222333333333344444444445555555555`
	mi2 = `012345678901234567890123456789012345678901234567890123456789`
	s1  = `000000000011111111112222222222333333333344444444445555555555`
	s2  = `012345678901234567890123456789012345678901234567890123456789`
)

// Legacy loglevel constants to ensure backwards compatibility.
//
// Deprecated: will be removed in 1.5.0.
const (
	LevelInfo  = LevelInformational
	LevelTrace = LevelDebug
	LevelWarn  = LevelWarning
)

type loggerType func() Logger

// Logger defines the behavior of a log provider.
type Logger interface {
	Init(config string) error
	WriteMsg(when time.Time, msg string, level int) error
	Destroy()
	Flush()
}

var adapters = make(map[string]loggerType)

// Register makes a log provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, log loggerType) {
	if log == nil {
		panic("logs: Register provide is nil")
	}
	if _, dup := adapters[name]; dup {
		panic("logs: Register called twice for provider " + name)
	}
	adapters[name] = log
}

// BeeLogger is default logger in beego application.
// it can contain several providers and log message into all providers.
type BeeLogger struct {
	lock                sync.Mutex
	level               int
	enableFuncCallDepth bool
	loggerFuncCallDepth int
	asynchronous        bool
	msgChan             chan *logMsg
	outputs             []*nameLogger
}

type nameLogger struct {
	Logger
	name string
}

type logMsg struct {
	level int
	msg   string
	when  time.Time
}

var logMsgPool *sync.Pool

// NewLogger returns a new BeeLogger.
// channelLen means the number of messages in chan(used where asynchronous is true).
// if the buffering chan is full, logger adapters write to file or other way.
func NewLogger(channelLen int64) *BeeLogger {
	bl := new(BeeLogger)
	bl.level = LevelDebug
	bl.loggerFuncCallDepth = 2
	bl.msgChan = make(chan *logMsg, channelLen)
	return bl
}

// Async set the log to asynchronous and start the goroutine
func (bl *BeeLogger) Async() *BeeLogger {
	bl.asynchronous = true
	logMsgPool = &sync.Pool{
		New: func() interface{} {
			return &logMsg{}
		},
	}
	go bl.startLogger()
	return bl
}

// SetLogger provides a given logger adapter into BeeLogger with config string.
// config need to be correct JSON as string: {"interval":360}.
func (bl *BeeLogger) SetLogger(adapterName string, config string) error {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	if log, ok := adapters[adapterName]; ok {
		lg := log()
		err := lg.Init(config)
		if err != nil {
			fmt.Fprintln(os.Stderr, "logs.BeeLogger.SetLogger: "+err.Error())
			return err
		}
		bl.outputs = append(bl.outputs, &nameLogger{name: adapterName, Logger: lg})
	} else {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}
	return nil
}

// DelLogger remove a logger adapter in BeeLogger.
func (bl *BeeLogger) DelLogger(adapterName string) error {
	bl.lock.Lock()
	defer bl.lock.Unlock()
	outputs := []*nameLogger{}
	for _, lg := range bl.outputs {
		if lg.name == adapterName {
			lg.Destroy()
		} else {
			outputs = append(outputs, lg)
		}
	}
	if len(outputs) == len(bl.outputs) {
		return fmt.Errorf("logs: unknown adaptername %q (forgotten Register?)", adapterName)
	}
	bl.outputs = outputs
	return nil
}

func (bl *BeeLogger) writeToLoggers(when time.Time, msg string, level int) {
	for _, l := range bl.outputs {
		err := l.WriteMsg(when, msg, level)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to WriteMsg to adapter:%v,error:%v\n", l.name, err)
		}
	}
}

func (bl *BeeLogger) writeMsg(logLevel int, msg string) error {
	when := time.Now()
	if bl.enableFuncCallDepth {
		_, file, line, ok := runtime.Caller(bl.loggerFuncCallDepth)
		if !ok {
			file = "???"
			line = 0
		}
		_, filename := path.Split(file)
		msg = "[" + filename + ":" + strconv.FormatInt(int64(line), 10) + "]" + msg
	}
	if bl.asynchronous {
		lm := logMsgPool.Get().(*logMsg)
		lm.level = logLevel
		lm.msg = msg
		lm.when = when
		bl.msgChan <- lm
	} else {
		bl.writeToLoggers(when, msg, logLevel)
	}
	return nil
}

// SetLevel Set log message level.
// If message level (such as LevelDebug) is higher than logger level (such as LevelWarning),
// log providers will not even be sent the message.
func (bl *BeeLogger) SetLevel(l int) {
	bl.level = l
}

// SetLogFuncCallDepth set log funcCallDepth
func (bl *BeeLogger) SetLogFuncCallDepth(d int) {
	bl.loggerFuncCallDepth = d
}

// GetLogFuncCallDepth return log funcCallDepth for wrapper
func (bl *BeeLogger) GetLogFuncCallDepth() int {
	return bl.loggerFuncCallDepth
}

// EnableFuncCallDepth enable log funcCallDepth
func (bl *BeeLogger) EnableFuncCallDepth(b bool) {
	bl.enableFuncCallDepth = b
}

// start logger chan reading.
// when chan is not empty, write logs.
func (bl *BeeLogger) startLogger() {
	for {
		select {
		case bm := <-bl.msgChan:
			bl.writeToLoggers(bm.when, bm.msg, bm.level)
			logMsgPool.Put(bm)
		}
	}
}

// Emergency Log EMERGENCY level message.
func (bl *BeeLogger) Emergency(format string, v ...interface{}) {
	if LevelEmergency > bl.level {
		return
	}
	msg := fmt.Sprintf("[M] "+format, v...)
	bl.writeMsg(LevelEmergency, msg)
}

// Alert Log ALERT level message.
func (bl *BeeLogger) Alert(format string, v ...interface{}) {
	if LevelAlert > bl.level {
		return
	}
	msg := fmt.Sprintf("[A] "+format, v...)
	bl.writeMsg(LevelAlert, msg)
}

// Critical Log CRITICAL level message.
func (bl *BeeLogger) Critical(format string, v ...interface{}) {
	if LevelCritical > bl.level {
		return
	}
	msg := fmt.Sprintf("[C] "+format, v...)
	bl.writeMsg(LevelCritical, msg)
}

// Error Log ERROR level message.
func (bl *BeeLogger) Error(format string, v ...interface{}) {
	if LevelError > bl.level {
		return
	}
	msg := fmt.Sprintf("[E] "+format, v...)
	bl.writeMsg(LevelError, msg)
}

// Warning Log WARNING level message.
func (bl *BeeLogger) Warning(format string, v ...interface{}) {
	if LevelWarning > bl.level {
		return
	}
	msg := fmt.Sprintf("[W] "+format, v...)
	bl.writeMsg(LevelWarning, msg)
}

// Notice Log NOTICE level message.
func (bl *BeeLogger) Notice(format string, v ...interface{}) {
	if LevelNotice > bl.level {
		return
	}
	msg := fmt.Sprintf("[N] "+format, v...)
	bl.writeMsg(LevelNotice, msg)
}

// Informational Log INFORMATIONAL level message.
func (bl *BeeLogger) Informational(format string, v ...interface{}) {
	if LevelInformational > bl.level {
		return
	}
	msg := fmt.Sprintf("[I] "+format, v...)
	bl.writeMsg(LevelInformational, msg)
}

// Debug Log DEBUG level message.
func (bl *BeeLogger) Debug(format string, v ...interface{}) {
	if LevelDebug > bl.level {
		return
	}
	msg := fmt.Sprintf("[D] "+format, v...)
	bl.writeMsg(LevelDebug, msg)
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func (bl *BeeLogger) Warn(format string, v ...interface{}) {
	if LevelWarning > bl.level {
		return
	}
	msg := fmt.Sprintf("[W] "+format, v...)
	bl.writeMsg(LevelWarning, msg)
}

// Info Log INFO level message.
// compatibility alias for Informational()
func (bl *BeeLogger) Info(format string, v ...interface{}) {
	if LevelInformational > bl.level {
		return
	}
	msg := fmt.Sprintf("[I] "+format, v...)
	bl.writeMsg(LevelInformational, msg)
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func (bl *BeeLogger) Trace(format string, v ...interface{}) {
	if LevelDebug > bl.level {
		return
	}
	msg := fmt.Sprintf("[D] "+format, v...)
	bl.writeMsg(LevelDebug, msg)
}

// Flush flush all chan data.
func (bl *BeeLogger) Flush() {
	for _, l := range bl.outputs {
		l.Flush()
	}
}

// Close close logger, flush all chan data and destroy all adapters in BeeLogger.
func (bl *BeeLogger) Close() {
	for {
		if len(bl.msgChan) > 0 {
			bm := <-bl.msgChan
			bl.writeToLoggers(bm.when, bm.msg, bm.level)
			logMsgPool.Put(bm)
			continue
		}
		break
	}
	for _, l := range bl.outputs {
		l.Flush()
		l.Destroy()
	}
}

func formatLogTime(when time.Time) string {
	y, mo, d := when.Date()
	h, mi, s := when.Clock()
	//len(2006/01/02 15:03:04)==19
	var buf [20]byte
	t := 3
	for y >= 10 {
		p := y / 10
		buf[t] = byte('0' + y - p*10)
		y = p
		t--
	}
	buf[0] = byte('0' + y)
	buf[4] = '/'
	if mo > 9 {
		buf[5] = '1'
		buf[6] = byte('0' + mo - 9)
	} else {
		buf[5] = '0'
		buf[6] = byte('0' + mo)
	}
	buf[7] = '/'
	t = d / 10
	buf[8] = byte('0' + t)
	buf[9] = byte('0' + d - t*10)
	buf[10] = ' '
	t = h / 10
	buf[11] = byte('0' + t)
	buf[12] = byte('0' + h - t*10)
	buf[13] = ':'
	t = mi / 10
	buf[14] = byte('0' + t)
	buf[15] = byte('0' + mi - t*10)
	buf[16] = ':'
	t = s / 10
	buf[17] = byte('0' + t)
	buf[18] = byte('0' + s - t*10)
	buf[19] = ' '

	return string(buf[0:])
}

func formatTimeHeader(when time.Time) ([]byte, int) {
	y, mo, d := when.Date()
	h, mi, s := when.Clock()
	//len("2006/01/02 15:04:05 ")==20
	var buf [20]byte

	buf[0] = y1[y/1000%10]
	buf[1] = y2[y/100]
	buf[2] = y3[y-y/100*100]
	buf[3] = y4[y-y/100*100]
	buf[4] = '/'
	buf[5] = mo1[mo-1]
	buf[6] = mo2[mo-1]
	buf[7] = '/'
	buf[8] = d1[d-1]
	buf[9] = d2[d-1]
	buf[10] = ' '
	buf[11] = h1[h]
	buf[12] = h2[h]
	buf[13] = ':'
	buf[14] = mi1[mi]
	buf[15] = mi2[mi]
	buf[16] = ':'
	buf[17] = s1[s]
	buf[18] = s2[s]
	buf[19] = ' '

	return buf[0:], d
}