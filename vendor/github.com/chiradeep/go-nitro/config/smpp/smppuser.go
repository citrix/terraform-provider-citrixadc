package smpp

type Smppuser struct {
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}
