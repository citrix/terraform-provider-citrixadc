package ssl

type Sslglobalsslpolicybinding struct {
	Policyname string `json:"policyname,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	Type       string `json:"type,omitempty"`
}
