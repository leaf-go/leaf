package xsrv

const (
	HTTP Kind = iota + 1
	Websocket
	Task
	RPC
)

type Kind int

// XServer 服务接口定义
type XServer interface {
	GetConfig() Config
	Start() error
}

type Config struct {
	Host string
	Port int8
}

func New(kind Kind, config Config) {

}

type DefaultHttpServer struct {
}

type TaskServer struct {
	config Config
}

func (t TaskServer) GetConfig() Config {
	return t.config
}

func (t TaskServer) Start() error {

}

