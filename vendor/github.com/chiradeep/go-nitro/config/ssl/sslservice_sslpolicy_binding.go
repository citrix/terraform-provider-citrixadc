package ssl

type Sslservicesslpolicybinding struct {
	Gotopriorityexpression string `json:"gotopriorityexpression,omitempty"`
	Invoke                 bool   `json:"invoke,omitempty"`
	Labelname              string `json:"labelname,omitempty"`
	Labeltype              string `json:"labeltype,omitempty"`
	Policyname             string `json:"policyname,omitempty"`
	Polinherit             int    `json:"polinherit,omitempty"`
	Priority               int    `json:"priority,omitempty"`
	Servicename            string `json:"servicename,omitempty"`
}
