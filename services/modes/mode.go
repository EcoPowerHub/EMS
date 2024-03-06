package modes

import (
	"fmt"

	context "github.com/EcoPowerHub/context/pkg"
	"github.com/EcoPowerHub/shared/pkg/objects"
)

func New(conf Conf, ctx *context.Context) (*Manager, error) {
	return &Manager{
		conf: conf,
		ctx:  ctx,
	}, nil
}

type Manager struct {
	conf       Conf
	ctx        *context.Context
	actualMode mode
}

func (m *Manager) Runnable(name string) bool {
	return false
}

func (m *Manager) Update() error {
	var (
		status *objects.Status
		err    error
	)

	for key, mode := range m.conf.Modes {
		if mode.ConditionRef == "default" {
			m.actualMode = mode
			return nil
		}
		status, err = m.ctx.Status(mode.ConditionRef)
		if err != nil {
			return fmt.Errorf("failed to get condition for mode %s: %s", key, err)
		}

		if status.Value == 1 {
			m.actualMode = mode
			return nil
		}
	}
	return nil
}
