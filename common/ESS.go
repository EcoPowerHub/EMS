package common

type ESS struct {
	P_kW        float64 `json:"p_kw"`
	Soc         float64 `json:"soc"`
	Soh         float64 `json:"soh"`
	Capacity_Wh float64 `json:"capacity_wh"`
	Timestamp   int64   `json:"timestamp"`
}

func (e *ESS) Copy() ESS {
	return ESS{
		P_kW:        e.P_kW,
		Soc:         e.Soc,
		Soh:         e.Soh,
		Capacity_Wh: e.Capacity_Wh,
	}
}
