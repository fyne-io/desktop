// +build openbsd freebsd netbsd

package status

import "syscall"

func (b *battery) value() (float64, error) {
	val, err := syscall.Sysctl("hw.acpi.battery.life")
	if err != nil {
		return 0, err
	}

	percent := int(val[0])
	if percent == 0 { // avoid 0/100 below
		return 0, nil
	}

	return float64(percent) / 100, nil
}