package models

type Worksheet struct {
	Secrets []Secret `yaml:"secrets"`
	Items   []Item   `yaml:"items"`
}

type Item struct {
	Name     string `yaml:"name"`
	Data     string `yaml:"data"`
	Template string `yaml:"template"`
	Engine   string `yaml:"engine" default:"go"`
	Output   string `yaml:"output"`
}

type Secret struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
	Env    string `yaml:"env"`
	Region string `yaml:"region"`
}
