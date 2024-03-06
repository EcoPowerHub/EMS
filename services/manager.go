package services

import (
	"fmt"
	"sort"

	context "github.com/EcoPowerHub/context/pkg"
	"golang.org/x/sync/errgroup"

	"github.com/EcoPowerHub/EMS/config"
	"github.com/EcoPowerHub/EMS/services/modes"
	eval "github.com/EcoPowerHub/EMS/services/services/gval"
	"github.com/EcoPowerHub/EMS/services/services/peakshaving"
	"github.com/rs/zerolog"
)

type Service interface {
	Configure(configuration any, inputsConf any, outputsConf any) error
	Update() error
}

type Manager struct {
	ctx            *context.Context
	conf           map[string]config.Service
	services       map[string]Service
	sortedServices map[int][]string
	modeManager    *modes.Manager
}

func New(conf map[string]config.Service, ctx *context.Context, modeConf modes.Conf) (*Manager, error) {
	modeManager, err := modes.New(modeConf, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create mode manager: %s", err)
	}

	return &Manager{
		conf:        conf,
		services:    make(map[string]Service),
		ctx:         ctx,
		modeManager: modeManager,
	}, nil
}

// SetupServices sets up the services based on the configuration provided.
// It creates and configures each service, sorts them by priority, and stores them in the Manager.
// Returns an error if there is any failure during the setup process.
// SetupServices initializes and configures the services based on the provided configuration.
// It iterates over the service configurations, creates new service instances, and configures them.
// The services are then sorted by priority and stored in the `sortedServices` map.
// Returns an error if there is any failure in creating or configuring the services.
func (m *Manager) SetupServices() error {
	for name, serviceConf := range m.conf {
		service, err := m.newService(serviceConf.Id)
		if err != nil {
			return fmt.Errorf("failed to create service %s: %s", name, err)
		}
		err = service.Configure(serviceConf.Conf.Conf, serviceConf.Conf.Inputs, serviceConf.Conf.Outputs)
		if err != nil {
			return fmt.Errorf("failed to configure service %s: %s", name, err)
		}

		m.services[name] = service
	}

	priorityArray := make([]int, 0)
	// Sort services by priority
	for _, service := range m.conf {
		priorityArray = append(priorityArray, int(service.Priority))
	}

	sort.Ints(priorityArray)

	uniquePriorityArray := make([]int, 0)
	for _, priority := range priorityArray {
		found := false
		for _, uniquePriority := range uniquePriorityArray {
			if priority == uniquePriority {
				found = true
				break
			}
		}
		if !found {
			uniquePriorityArray = append(uniquePriorityArray, priority)
		}
	}

	for _, uniquePriority := range uniquePriorityArray {
		serviceArray := []string{}
		for service := range m.services {
			if int(m.conf[service].Priority) == uniquePriority {
				serviceArray = append(serviceArray, service)
			}
		}
		m.sortedServices[uniquePriority] = serviceArray
	}

	return nil
}

func (m *Manager) UpdateServices() (err error) {
	var (
		g = errgroup.Group{}
	)
	// Update the mode
	if err := m.modeManager.Update(); err != nil {
		return fmt.Errorf("failed to update mode: %s", err)
	}

	// Loop through services and update thoses contained in the actual mode
	for _, k := range m.sortedServices {
		for _, service := range k {
			// If the service is not enabled, skip it
			if m.modeManager.Runnable(service) {
				g.Go(m.services[service].Update)
			}
		}
		// Wait for all Equipments to finish
		if err = g.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) newService(id string) (Service, error) {
	switch id {
	case "peakshaving":
		return peakshaving.New(zerolog.Logger{}, m.ctx), nil
	case "gval":
		return eval.New(zerolog.Logger{}, m.ctx), nil
	default:
		return nil, fmt.Errorf("unknown service id: %s", id)
	}
}
