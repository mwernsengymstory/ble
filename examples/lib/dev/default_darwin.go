package dev

import (
	"github.com/mwernsengymstory/ble"
	"github.com/mwernsengymstory/ble/darwin"
)

// DefaultDevice ...
func DefaultDevice(opts ...ble.Option) (d ble.Device, err error) {
	return darwin.NewDevice(opts...)
}
