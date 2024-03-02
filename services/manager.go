package services

import (
	"fmt"
	"sort"

	context "github.com/EcoPowerHub/context/pkg"

	"github.com/EcoPowerHub/EMS/config"
	"github.com/EcoPowerHub/EMS/services/peakshaving"
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
	sortedServices []string
}

func New(conf map[string]config.Service, ctx *context.Context) (*Manager, error) {
	return &Manager{
		conf:     conf,
		services: make(map[string]Service),
		ctx:      ctx,
	}, nil
}

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

	// Sort services by priority
	for k := range m.conf {
		m.sortedServices = append(m.sortedServices, k)
	}

	sort.Slice(m.sortedServices, func(i, j int) bool {
		return m.conf[m.sortedServices[i]].Priority < m.conf[m.sortedServices[j]].Priority
	})

	return nil
}

func (m *Manager) UpdateServices() error {
	for _, k := range m.sortedServices {
		if err := m.services[k].Update(); err != nil {
			return fmt.Errorf("failed to update service %s: %s", k, err)
		}
	}
	return nil
}

func (m *Manager) newService(id string) (Service, error) {
	switch id {
	case "peakshaving":
		return peakshaving.New(zerolog.Logger{}, m.ctx), nil
	default:
		return nil, fmt.Errorf("unknown service id: %s", id)
	}
}
