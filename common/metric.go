package common

type Metric struct {
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

func (m *Metric) Copy() Metric {
	return Metric{
		Value:     m.Value,
		Timestamp: m.Timestamp,
	}
}
