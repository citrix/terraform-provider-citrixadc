package ns

type Nsversion struct {
	Installedversion bool   `json:"installedversion,omitempty"`
	Mode             int    `json:"mode,omitempty"`
	Version          string `json:"version,omitempty"`
}
