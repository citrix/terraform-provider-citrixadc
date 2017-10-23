package ssl

type Sslkeyfile struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Src      string `json:"src,omitempty"`
}
