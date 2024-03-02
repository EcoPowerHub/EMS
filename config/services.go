package config

type Service struct {
	Id       string      `json:"id"`
	Priority uint        `json:"priority"`
	Conf     ServiceConf `json:"conf"`
}

type ServiceConf struct {
	Inputs  map[string]any `json:"inputs"`
	Outputs map[string]any `json:"outputs"`
	Conf    map[string]any `json:"conf"`
}
