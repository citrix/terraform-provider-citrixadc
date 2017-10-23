package ssl

type Sslprofileecccurvebinding struct {
	Cipherpriority int    `json:"cipherpriority,omitempty"`
	Ecccurvename   string `json:"ecccurvename,omitempty"`
	Name           string `json:"name,omitempty"`
}
