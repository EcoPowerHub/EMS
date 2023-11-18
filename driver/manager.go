package driver

import (
	"fmt"

	"github.com/EcoPowerHub/EMS/common"
	"github.com/EcoPowerHub/EMS/config"
	equipment "github.com/EcoPowerHub/EMS/driver/factory"
	"golang.org/x/sync/errgroup"
)

type Manager struct {
	Equipments []equipment.Driver
}

// New takes a list of equipments from config file and returns a managerEquipment
func New(equipments []config.Equipment) (*Manager, error) {
	var (
		err     error
		manager = &Manager{}
	)
	manager.Equipments, err = equipment.Instanciate(equipments)
	return manager, err
}

// Read triggers the Read method of all equipments and returns the results
func (m *Manager) Read() map[string]map[string]any {
	var read map[string]map[string]any
	for _, d := range m.Equipments {
		read = d.Read()
	}
	return read
}

// SetupEquipments triggers the Configure method of all equipments
func (m *Manager) SetupEquipments() (err error) {
	for _, e := range m.Equipments {
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
	// Launch all Equipments
	for _, e := range m.Equipments {
		if e.State().Value != common.EquipmentStateOnline {
			// #8
			fmt.Printf("Equipment is not online, skipping")
			continue
		}
		g.Go(e.AddOrRefreshData)
	}
	// Wait for all Equipments to finish
	if err = g.Wait(); err != nil {
		return err
	}
	return
}
