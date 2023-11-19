package driver

type Driver interface {
	Configure()
	Start()
	Cycle()
}
