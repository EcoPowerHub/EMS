package manager

import (
	"os"

	"github.com/EcoPowerHub/EMS/common"
	"github.com/EcoPowerHub/EMS/config"
	equipment "github.com/EcoPowerHub/EMS/driver/factory"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"
)

type Manager struct {
	drivers []equipment.Driver
}

func New(conf config.Equipments) (*Manager, error) {
	var (
		err     error
		manager = &Manager{}
	)

	manager.drivers, err = equipment.Instanciate(conf)

	return manager, err
}

func (m *Manager) Read() map[string]map[string]any {
	var read map[string]map[string]any
	for _, d := range m.drivers {
		read = d.Read()
	}
	return read
}

func (m *Manager) Configure() (err error) {
	for _, d := range m.drivers {
		if err = d.Configure(); err != nil {
			return
		}
	}
	return
}

func (m *Manager) Cycle() (err error) {
	var (
		g      = errgroup.Group{}
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	)

	// Launch all drivers
	for _, d := range m.drivers {
		if d.State().Value != common.DriverStateOnline {
			// TODO replace logger instanciation, see #8
			logger.Warn().Str("driver", "TODO").Msg("Driver is not online, skipping")
			continue
		}
		g.Go(d.Cycle)
	}

	// Wait for all drivers to finish
	if err = g.Wait(); err != nil {
		return err
	}
	return
}
