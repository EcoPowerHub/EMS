package manager

import "github.com/EcoPowerHub/EMS/driver"

type Manager struct {
	drivers map[string]driver.Driver
}
