package mounts

type ILimiter interface {
	Allow() bool
	Remain() int
}

type Limit struct {
}

func NewLimit() *Limit {
	return &Limit{}
}
