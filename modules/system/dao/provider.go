package dao

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	_ = container.Container.Provide(NewUserDao)
	return nil
}
