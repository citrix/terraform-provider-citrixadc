package ssl

type Sslcertkeysslocspresponderbinding struct {
	Ca            bool   `json:"ca,omitempty"`
	Certkey       string `json:"certkey,omitempty"`
	Ocspresponder string `json:"ocspresponder,omitempty"`
	Priority      int    `json:"priority,omitempty"`
}
