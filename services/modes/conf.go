package modes

type Conf struct {
	Modes []mode `json:"modes"`
}

type mode struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	EnabledServices []string `json:"enabledServices"`
	ConditionRef    string   `json:"condition"`
}
