package x

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/fatih/color"
	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type Logger interface {
	// GenerateTraceId 创建追踪id
	GenerateTraceId()

	// GetId 获取追踪ID
	GetId() string

	// User 注入数据
	User(user interface{}) Logger

	Params(params interface{}) Logger

	Result(result interface{}) Logger

	// Auto 自动识别log等级
	Auto(no IErrno, message string, data interface{}, extras ...interface{})

	// LogInfo LogTrace LogDebug LogFatal LogPanic LogError LogWarning 不同等级的Log
	LogInfo(message string, data interface{}, extras ...interface{})
	LogTrace(message string, data interface{}, extras ...interface{})
	LogDebug(message string, data interface{}, extras ...interface{})
	LogWarning(message string, data interface{}, extras ...interface{})
	LogFatal(message string, data interface{}, extras ...interface{})
	LogPanic(message string, data interface{}, extras ...interface{})
	LogError(message string, data interface{}, extras ...interface{})
}

var (
	logger *logrus.Logger
	log    *defaultLog
)

func NewLogger(method, action, ip string, toFile bool) *defaultLog {
	log = &defaultLog{}
	log.GenerateTraceId()
	log.method = method
	log.action = action
	log.ip = ip
	log.toFile = toFile
	return log
}

// defaultLog 应用log
type defaultLog struct {
	id     string // 追踪id
	toFile bool
	method string
	action string
	ip     string
	user   interface{}
	params interface{}
	result interface{}
}

func (l *defaultLog) User(user interface{}) Logger {
	l.user = user
	return l
}

func (l *defaultLog) Params(params interface{}) Logger {
	l.params = params
	return l
}

func (l *defaultLog) Result(result interface{}) Logger {
	l.result = result
	return l
}

func (l *defaultLog) Auto(no IErrno, message string, data interface{}, extras ...interface{}) {
	level, _ := logrus.ParseLevel(no.Level())
	go l.Log(level, l.format(message, data, extras))
}

func (l *defaultLog) GenerateTraceId() {
	buf := make([]byte, 36)
	u := uuid.NewV4().Bytes()
	hex.Encode(buf, u)
	l.id = string(buf)
}

func (l *defaultLog) Log(level logrus.Level, message string) {
	if !l.toFile {
		color.Yellow("[%s]: %s , stack:%s", message, debug.Stack())
		return
	}

	go logger.Log(level, message)
}

func (l *defaultLog) LogTrace(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.TraceLevel, l.format(message, data, extras))
}

func (l *defaultLog) LogDebug(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.DebugLevel, l.format(message, data, extras))
}

func (l *defaultLog) LogFatal(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.FatalLevel, l.format(message, data, extras))
}

func (l *defaultLog) LogPanic(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.PanicLevel, l.format(message, data, extras))
}

func (l *defaultLog) LogError(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.ErrorLevel, l.format(message, data, extras))
}

func (l *defaultLog) LogInfo(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.InfoLevel, l.format(message, data, extras))
}

func (l *defaultLog) LogWarning(message string, data interface{}, extras ...interface{}) {
	l.Log(logrus.WarnLevel, l.format(message, data, extras))
}

// GetId 获取追踪ID
func (l *defaultLog) GetId() string {
	return l.id
}

// format 格式化数据
func (l *defaultLog) format(message string, data interface{}, extras ...interface{}) string {
	js, _ := json.Marshal(H{
		"user":   l.user,
		"params": l.params,
		"result": l.result,
		"data":   data,
		"extras": extras,
	})

	return fmt.Sprintf("TRACE:%s %s %s IP:%s <%s> CONTEXT:%s\ndebug:%s", l.id, l.method, l.action, l.ip, message, js, debug.Stack())
}

// json 将数据转换成json
func (l *defaultLog) json(i interface{}) string {
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

// LogConfig 配置log
type LogConfig struct {
	Path            string `json:"path" toml:"path"`
	Level           string `json:"level" toml:"level"`
	Stdout          bool   `json:"stdout" toml:"stdout"`
	SaveDay         uint   `json:"save_day" toml:"save_day"`
	TimestampFormat string `json:"timestamp_format" toml:"timestamp_format"`
}

func (l LogConfig) Init() {
	logger = logrus.New()

	if err := l.display(); err != nil {
		panic(fmt.Sprintf("log initialize failed: %v", err))
	}

	level, _ := logrus.ParseLevel(l.Level)
	logger.SetLevel(level)

	formatter := &Formatter{l.TimestampFormat}
	levels := []logrus.Level{logrus.DebugLevel, logrus.InfoLevel, logrus.WarnLevel, logrus.ErrorLevel, logrus.FatalLevel}
	lfHook := lfshook.NewHook(l.outputMap(levels...), formatter)
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
