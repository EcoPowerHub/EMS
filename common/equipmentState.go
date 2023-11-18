package common

const (
	EquipmentStateInit        = 1
	EquipmentStateOnline      = 2
	EquipmentStateUnreachable = 3
	EquipmentStateOffline     = 4
	EquipmentStateError       = 5
)

type EquipmentState struct {
	Value uint8 `json:"value"`
}
