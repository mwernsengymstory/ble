package dev

import (
	"github.com/mwernsengymstory/ble"
	"github.com/mwernsengymstory/ble/linux"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return linux.NewDevice(opts...)
}
