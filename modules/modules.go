package modules

import (
	system "atom/modules/system/container"

	"github.com/rogeecn/atom/utils/opt"
)

func Provide(opts ...opt.Option) error {
	return system.Provide()
}
