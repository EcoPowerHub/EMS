package common

const (
	DriverStateInit        = 1
	DriverStateOnline      = 2
	DriverStateUnreachable = 3
	DriverStateOffline     = 4
	DriverStateError       = 5
)

type DriverState struct {
	Value uint8 `json:"value"`
}
