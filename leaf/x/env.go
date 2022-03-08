package x

var (
	IsRelease bool
	Release   string
	Env       string
)

func InitEnv(env string, release string) {
	Env = env
	Release = release
	IsRelease = env == release
}
