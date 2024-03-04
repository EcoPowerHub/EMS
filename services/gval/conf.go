package eval

type conf struct {
	Expression string `json:"expression"`
}

type refWithAttr struct {
	Ref  string `mapstructure:"ref"`
	Attr string `mapstructure:"attr"`
}
