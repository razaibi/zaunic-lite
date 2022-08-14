package models

type Worksheet struct {
	Items []Item `yaml:"items"`
}

type Item struct {
	Name     string `yaml:"name"`
	Data     string `yaml:"data"`
	Template string `yaml:"template"`
	Output   string `yaml:"output"`
}
