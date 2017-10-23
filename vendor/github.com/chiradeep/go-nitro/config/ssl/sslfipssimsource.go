package ssl

type Sslfipssimsource struct {
	Certfile     string `json:"certfile,omitempty"`
	Sourcesecret string `json:"sourcesecret,omitempty"`
	Targetsecret string `json:"targetsecret,omitempty"`
}
