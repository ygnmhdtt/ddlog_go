package ddlog_go

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

var timeNowFunc = time.Now

type logLevel int

const (
	DEBUG logLevel = iota + 1
	INFO
	WARN
	ERROR
	FATAL
)

func (l logLevel) ddAttr() string {
	return "loglevel=" + l.String()
}

// String returns string value of logLevel
func (l logLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "INFO"
	}
}

// Datadog log format
// metric unix_timestamp value [attribute1=v1 attributes2=v2 ...]
type ddLogger struct {
	logger *log.Logger
	attr   map[string]string
}

// NewddLogger returns ddLogger which has specified name and empty attributes
func NewddLogger(name string, out io.Writer) *ddLogger {
	ddl := new(ddLogger)
	logger := log.New(out, name, 0)
	ddl.logger = logger
	ddl.ClearAttr()
	return ddl
}

// Attr sets attributes
func (ddl *ddLogger) Attr(key string, value string) *ddLogger {
	ddl.attr[key] = value
	return ddl
}

// ClearAttr clears all attributes
func (ddl *ddLogger) ClearAttr() {
	ddl.attr = map[string]string{}
}

// INFO prints INFO log
func (ddl *ddLogger) INFO(val string) {
	ddl.logger.Println(ddl.newline(val, INFO))
}

// DEBUG prints DEBUG log
func (ddl *ddLogger) DEBUG(val string) {
	ddl.logger.Println(ddl.newline(val, DEBUG))
}

// WARN prints WARN log
func (ddl *ddLogger) WARN(val string) {
	ddl.logger.Println(ddl.newline(val, WARN))
}

// ERROR prints ERROR log
func (ddl *ddLogger) ERROR(val string) {
	ddl.logger.Println(ddl.newline(val, ERROR))
}

// FATAL prints FATAL log
func (ddl *ddLogger) FATAL(val string) {
	ddl.logger.Println(ddl.newline(val, FATAL))
}

func (ddl *ddLogger) newline(val string, l logLevel) string {
	return fmt.Sprintf(" %v %v %v", timestamp(), val, ddl.attrStr(l))
}

func (ddl *ddLogger) attrStr(l logLevel) string {
	// this func returns "env=staging hoge=fuga"
	attrArr := []string{}
	attrArr = append(attrArr, l.ddAttr())
	for key, value := range ddl.attr {
		attrArr = append(attrArr, key+"="+value)
	}
	return strings.Join(attrArr, " ")
}

func timestamp() string {
	return fmt.Sprint(timeNowFunc().Unix())
}
