package x

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type Logger interface {
	GenerateTraceId()
	GetId() string
	Auto(no Errno, message string, data interface{}, extras ...interface{})
	LogTrace(message string, data interface{}, extras ...interface{})
	LogDebug(message string, data interface{}, extras ...interface{})
	LogFatal(message string, data interface{}, extras ...interface{})
	LogPanic(message string, data interface{}, extras ...interface{})
	LogError(message string, data interface{}, extras ...interface{})
	LogInfo(message string, data interface{}, extras ...interface{})
	LogWarning(message string, data interface{}, extras ...interface{})
}

var (
	logger *logrus.Logger
	log    *Log
)

func DefaultLogger() *Log {
	if log == nil {
		log = &Log{}
	}

	return log
}

// LogConfig 配置log
type LogConfig struct {
	Path            string `json:"path" toml:"path"`
	Level           string `json:"level" toml:"level"`
	Stdout          bool   `json:"stdout" toml:"stdout"`
	SaveDay         uint   `json:"save_day" toml:"save_day"`
	TimestampFormat string `json:"timestamp_format" toml:"timestamp_format"`
}

func (l LogConfig) Initialize() {
	logger = logrus.New()

	if err := l.display(); err != nil {
		panic(fmt.Sprintf("log initialize failed: %v", err))
	}

	level, _ := logrus.ParseLevel(l.Level)
	logger.SetLevel(level)

	formatter := &Formatter{l.TimestampFormat}
	lfHook := lfshook.NewHook(l.outputMap(), formatter)
	logger.AddHook(lfHook)
}

// display 关闭控制输出
func (l LogConfig) display() (err error) {
	if l.Stdout {
		return
	}

	file, err := os.OpenFile(os.DevNull, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	logger.SetOutput(file)
	return
}

// outputMap 切割字典
func (l LogConfig) outputMap(levels ...logrus.Level) lfshook.WriterMap {
	writers := make(lfshook.WriterMap)
	var b []byte
	for _, level := range levels {
		b, _ = level.MarshalText()
		writers[level] = getWriter(
			fmt.Sprintf("%s.%s", l.Path, b), l.SaveDay,
		)
	}

	return writers
}

// Log 应用log
type Log struct {
	id      string
	message string
}

func (l *Log) Auto(no IErrno, message string, data interface{}, extras ...interface{}) {
	level, _ := logrus.ParseLevel(no.Level())
	l.Log(level, l.format(message, data, extras))
}

func (l *Log) GenerateTraceId() {
	buf := make([]byte, 36)
	u := uuid.NewV4().Bytes()
	hex.Encode(buf, u)
	l.id = string(buf)
}

func (l *Log) Log(level logrus.Level, message string) {
	logger.Log(level, message)
}

func (l *Log) LogTrace(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.TraceLevel, l.format(message, data, extras))
}

func (l *Log) LogDebug(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.DebugLevel, l.format(message, data, extras))
}

func (l *Log) LogFatal(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.FatalLevel, l.format(message, data, extras))
}

func (l *Log) LogPanic(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.PanicLevel, l.format(message, data, extras))
}

func (l *Log) LogError(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.ErrorLevel, l.format(message, data, extras))
}

func (l *Log) LogInfo(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.InfoLevel, l.format(message, data, extras))
}

func (l *Log) LogWarning(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.WarnLevel, l.format(message, data, extras))
}

// GetId 获取追踪ID
func (l *Log) GetId() string {
	return l.id
}

// Inject 注入初始数据
func (l *Log) Inject(action, ip string, user, params interface{}) {
	l.GenerateTraceId()
	l.message = fmt.Sprintf("trace_id:%s action:%s ip:%s user:%s params:%s",
		l.id, action, ip, l.json(user), l.json(params))
}

// format 格式化数据
func (l *Log) format(message string, extras ...interface{}) string {
	prefix := "%s"
	if l.message != "" {
		prefix += " "
	}

	if len(extras) >= 2 {
		extras = extras[:2]
		return fmt.Sprintf(prefix+"%s data:%s traces:%v\n", l.message, message, l.json(extras[0]), l.json(extras[1]))
	}

	if len(extras) == 0 {
		return fmt.Sprintf(prefix+" %s\n", l.message, message)
	}

	return fmt.Sprintf(prefix+"%s data:%s\n", l.message, message, l.json(extras[0]))
}

// json 将数据转换成json
func (l *Log) json(i interface{}) string {
	if i != nil {
		b, _ := json.Marshal(i)
		return string(b)
	}

	return "{}"
}

func getWriter(path string, rotationCount uint) *rotate.RotateLogs {
	writer, err := rotate.New(
		path+".%Y%m%d%H",
		rotate.WithLinkName(path),
		rotate.WithRotationCount(rotationCount),
		rotate.WithRotationTime(time.Hour),
	)

	if err != nil {
		panic(fmt.Sprintf("InitLog writer error:%s, logPath:%s",
			err.Error(), path))
	}

	return writer
}

type Formatter struct {
	timestampFormat string
}

func (f Formatter) buf(entry *logrus.Entry) *bytes.Buffer {
	if entry.Buffer == nil {
		return &bytes.Buffer{}
	}

	return entry.Buffer
}

func (f *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	buf := f.buf(entry)

	buf.WriteString(
		fmt.Sprintf("[%v] %v %v",
			strings.ToUpper(entry.Level.String()), entry.Time.Format(f.timestampFormat), entry.Message,
		),
	)

	buf.WriteByte('\n')
	return buf.Bytes(), nil
}
