package types

type Config struct {
	Name    string            `json:"name"`
	Modules map[string]Module `json:"modules"`
}
