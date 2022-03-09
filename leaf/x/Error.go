package x

import "fmt"

var (
	errorsMapping map[IErrno]string
	UNKNOWN_ERRNO FatalErrno = -9999
)

func ErrorsInit(errors map[IErrno]string) {
	if errorsMapping == nil {
		errorsMapping = errors
		errorsMapping[UNKNOWN_ERRNO] = "unknown error"
	}
}

func ErrorSave(no IErrno, message string) {
	errorsMapping[no] = message
}

type Errors map[Errno]string

type IErrno interface {
	Level() string
}

type Errno int

type InfoErrno Errno

func (e InfoErrno) Level() string {
	return "info"
}

type TraceErrno Errno

func (e TraceErrno) Level() string {
	return "trace"
}

type DebugErrno Errno

func (e DebugErrno) Level() string {
	return "debug"
}

type WarnErrno Errno

func (e WarnErrno) Level() string {
	return "warning"
}

type ErrorErrno Errno

func (e ErrorErrno) Level() string {
	return "error"
}

type FatalErrno Errno

func (e FatalErrno) Level() string {
	return "fatal"
}

type PanicErrno Errno

func (e PanicErrno) Level() string {
	return "panic"
}

//type Error struct {
//	Code IErrno `json:"code"`
//	Message string `json:"message"`
//
//}

func ThrowError(errno IErrno) {
	message := errorMessage(errno)
	log.Auto(errno, message, nil)
	panic(errno)
}

func errorMessage(errno IErrno, def ...string) string {
	err, ok := errorsMapping[errno]
	if ok {
		return err
	}

	if len(def) == 0 {
		return errorsMapping[UNKNOWN_ERRNO]
	}

	return def[0]
}

func Recover(fn func(r interface{}, message string)) {
	if r := recover(); r != nil {
		switch r.(type) {
		case IErrno:
			fn(r, errorsMapping[r.(IErrno)])
			break
		default:
			fn(UNKNOWN_ERRNO, fmt.Sprintf("%+v", r))
		}
	}
}
