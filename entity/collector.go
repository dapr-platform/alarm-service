package entity

type Collector struct {
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	ParseScript string `json:"parse_script"`
}
