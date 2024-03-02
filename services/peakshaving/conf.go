package peakshaving

type conf struct {
	PID pidConf `mapstructure:"pid_controller"`
}

type pidConf struct {
	Kp     float64 `json:"kp"`
	Ki     float64 `json:"ki"`
	Kd     float64 `json:"kd"`
	Period string  `json:"period"`
}

type inputs struct {
	Poc_kW   refWithAttr `json:"poc_kW"`
	Limit_kW refWithAttr `json:"limit_kW"`
}

type outputs struct {
	Setpoint ref         `json:"setpoint"`
	PidError refWithAttr `json:"pid_error"`
}

type refWithAttr struct {
	Ref  string `json:"ref"`
	Attr string `json:"attr"`
}

type ref struct {
	Ref string `json:"ref"`
}
