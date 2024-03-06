package driver

import (
	"fmt"

	"github.com/EcoPowerHub/EMS/config"
	equipment "github.com/EcoPowerHub/EMS/driver/factory"
	context "github.com/EcoPowerHub/context/pkg"
	"github.com/EcoPowerHub/shared/pkg/objects"
	"golang.org/x/sync/errgroup"
)

type Manager struct {
	Objects []equipment.ManagerObject
	ctx     *context.Context
}

// New takes a list of equipments from config file and returns a managerEquipment
func New(equipments []config.Equipment, ctx *context.Context) (*Manager, error) {
	var (
		err     error
		manager = &Manager{
			ctx: ctx,
		}
	)
	manager.Objects, err = equipment.Instanciate(equipments)
	return manager, err
}

// Read triggers the Read method of all equipments and returns the results
func (m *Manager) Read() error {
	var (
		read map[string]map[string]any
		err  error
	)

	for _, d := range m.Objects {
		read = d.Driver.Read()
		for key1, value1 := range d.Equipement.Outputs {
			for key, value := range value1 {
				if read[key1][key] == nil {
					return fmt.Errorf("object [%s] key [%s] does not exists", key1, key)
				}
				err = m.ctx.Set(value.Ref, read[key1][key])
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Read triggers the Read method of all equipments and returns the results
func (m *Manager) Write() error {
	var (
		ctxValue any
		err      error
		writings map[string]map[string]any
	)

	for _, d := range m.Objects {
		for key1, value1 := range d.Equipement.Inputs {
			for key, value := range value1 {
				ctxValue, err = m.ctx.Get(value.Ref)
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
		if d.Driver.Write(writings) != nil {
			continue
		}
	}
	return nil
}

// SetupEquipments triggers the Configure method of all equipments
func (m *Manager) SetupEquipments() (err error) {
	for _, e := range m.Objects {
		if err = e.Driver.Configure(); err != nil {
			return
		}
	}
	return
}

func (m *Manager) InitCycle() (err error) {
	var (
		g = errgroup.Group{}
	)
	// Launch all Equipments
	for _, e := range m.Objects {
		if e.Driver.State().Value != objects.DriverStateOnline {
			// #8
			fmt.Printf("Equipment is not online, skipping")
			continue
		}
		g.Go(e.Driver.AddOrRefreshData)
	}
	// Wait for all Equipments to finish
	if err = g.Wait(); err != nil {
		return err
	}
	return
}
