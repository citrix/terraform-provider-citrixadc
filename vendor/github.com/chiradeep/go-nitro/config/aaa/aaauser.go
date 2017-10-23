package aaa

type Aaauser struct {
	Loggedin bool   `json:"loggedin,omitempty"`
	Password string `json:"password,omitempty"`
	Username string `json:"username,omitempty"`
}
