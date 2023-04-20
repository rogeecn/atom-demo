package modules

import (
	system "atom/modules/system/container"
)

func Provide() error {
	return system.Provide()
}
