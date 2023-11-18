package common

type Status struct {
	Value     float64 `json:"value"`
	Timestamp int64   `json:"timestamp"`
}

func (s *Status) Copy() Status {
	return Status{
		Value:     s.Value,
		Timestamp: s.Timestamp,
	}
}
