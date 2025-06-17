package types

type Module struct {
	Name     string `json:"name"`
	Author   string `json:"author"`
	Version  string `json:"version"`
	Path     string `json:"path"`
	IsActive bool   `json:"is_active"`
}
