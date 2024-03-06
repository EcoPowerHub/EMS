package driver

import (
	"fmt"

	"github.com/EcoPowerHub/EMS/config"
	driver "github.com/EcoPowerHub/EMS/driver/factory"
	context "github.com/EcoPowerHub/context/pkg"
	"github.com/EcoPowerHub/shared/pkg/objects"
	"golang.org/x/sync/errgroup"
)

type Manager struct {
	Manager driver.Manager
	ctx     *context.Context
}

// New takes a list of drivers from config file and returns a managerDriver
func New(drivers []config.Driver, ctx *context.Context) (*Manager, error) {
	var (
		err     error
		manager = &Manager{
			ctx: ctx,
		}
	)
	manager.Manager.Drivers, err = manager.Manager.Instanciate(drivers)
	return manager, err
}

// Read triggers the Read method of all drivers and returns the results
func (m *Manager) Read() error {
	var (
		read map[string]map[string]any
		err  error
	)

	for _, d := range m.Manager.Drivers {
		read = d.Read()
		for key1, value1 := range read {
			for key, value := range value1 {
				if read[key1][key] == nil {
					return fmt.Errorf("object [%s] key [%s] does not exists", key1, key)
				}
				err = m.ctx.Set(key, value)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Read triggers the Read method of all drivers and returns the results
func (m *Manager) Write() error {
	var (
		ctxValue any
		err      error
		writings map[string]map[string]any
	)

	for _, d := range m.Manager.Drivers {
		for key1, value1 := range d.Read() {
			for key, _ := range value1 {
				ctxValue, err = m.ctx.Get(key)
				writings = map[string]map[string]any{
					key1: {
						key: ctxValue,
					},
				}
				if err != nil {
					return err
				}
			}
		}
		if d.Write(writings) != nil {
			return err
		}
	}
	return nil
}

// SetupDrivers triggers the Configure method of all Drivers
func (m *Manager) SetupDrivers() (err error) {
	for _, e := range m.Manager.Drivers {
		if err = e.Configure(); err != nil {
			return
		}
	}
	return
}

func (m *Manager) InitCycle() (err error) {
	var (
		g = errgroup.Group{}
	)
	// Launch all Drivers
	for _, e := range m.Manager.Drivers {
		if e.State().Value != objects.EquipmentStateOnline {
			// #8
			fmt.Printf("Driver is not online, skipping")
			continue
		}
		g.Go(e.AddOrRefreshData)
	}
	// Wait for all Drivers to finish
	if err = g.Wait(); err != nil {
		return err
	}
	return
}
