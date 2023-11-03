package common

type PV struct {
	P_kW       float64 `json:"p_kw"`
	Capacity_W float64 `json:"capacity_w"`
}

func (p *PV) Copy() PV {
	return PV{
		P_kW:       p.P_kW,
		Capacity_W: p.Capacity_W,
	}
}
