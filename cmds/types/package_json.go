package types

type PackageJSON struct {
	Name        string            `json:"name"`
	Version     string            `json:"version"`
	Author      string            `json:"author"`
	Description string            `json:"description"`
	License     string            `json:"license"`
	Scripts     map[string]string `json:"scripts"`
}
