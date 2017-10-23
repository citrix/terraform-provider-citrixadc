package network

type Arpparam struct {
	Spoofvalidation string `json:"spoofvalidation,omitempty"`
	Timeout         int    `json:"timeout,omitempty"`
}
