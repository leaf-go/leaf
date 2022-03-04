package x

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

type WarningErrno Errno

func (e WarningErrno) Level() string {
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
