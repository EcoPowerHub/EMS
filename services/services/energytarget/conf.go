package energytarget

type conf struct {
	PID    pidConf `mapstructure:"pid_controller"`
	Target int8    `mapstructure:"target"`
}

type pidConf struct {
	Kp     float64 `json:"kp"`
	Ki     float64 `json:"ki"`
	Kd     float64 `json:"kd"`
	Period string  `json:"period"`
}

type inputs struct {
	Ess ref `json:"poc_kW"`
}

type outputs struct {
	Setpoint ref `json:"setpoint"`
}

type ref struct {
	Ref string `json:"ref"`
}
