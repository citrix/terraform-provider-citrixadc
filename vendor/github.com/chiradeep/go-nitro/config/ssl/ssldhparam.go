package ssl

type Ssldhparam struct {
	Bits   int    `json:"bits,omitempty"`
	Dhfile string `json:"dhfile,omitempty"`
	Gen    string `json:"gen,omitempty"`
}
