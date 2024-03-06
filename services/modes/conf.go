package modes

type Conf struct {
	Modes map[string]mode `json:"modes"`
}

type mode struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	EnabledServices []string `json:"enabledServices"`
	ConditionRef    string   `json:"condition"`
}
