package network

type Netprofile struct {
	Mbf              string `json:"mbf,omitempty"`
	Name             string `json:"name,omitempty"`
	Overridelsn      string `json:"overridelsn,omitempty"`
	Srcip            string `json:"srcip,omitempty"`
	Srcippersistency string `json:"srcippersistency,omitempty"`
	Td               int    `json:"td,omitempty"`
}
