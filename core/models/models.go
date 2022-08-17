package models

type Worksheet struct {
	Items []Item `yaml:"items"`
}

type Item struct {
	Name     string `yaml:"name"`
	Data     string `yaml:"data"`
	Template string `yaml:"template"`
	Engine   string `yaml:"engine" default:"gotemp"`
	Output   string `yaml:"output"`
}
