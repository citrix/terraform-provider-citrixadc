package ns

type Nsmode struct {
	Bridgebpdus         bool        `json:"bridgebpdus,omitempty"`
	Cka                 bool        `json:"cka,omitempty"`
	Dradv               bool        `json:"dradv,omitempty"`
	Dradv6              bool        `json:"dradv6,omitempty"`
	Edge                bool        `json:"edge,omitempty"`
	Fr                  bool        `json:"fr,omitempty"`
	Iradv               bool        `json:"iradv,omitempty"`
	L2                  bool        `json:"l2,omitempty"`
	L3                  bool        `json:"l3,omitempty"`
	Mbf                 bool        `json:"mbf,omitempty"`
	Mediaclassification bool        `json:"mediaclassification,omitempty"`
	Mode                interface{} `json:"mode,omitempty"`
	Pmtud               bool        `json:"pmtud,omitempty"`
	Riseapbr            bool        `json:"rise_apbr,omitempty"`
	Riserhi             bool        `json:"rise_rhi,omitempty"`
	Sradv               bool        `json:"sradv,omitempty"`
	Sradv6              bool        `json:"sradv6,omitempty"`
	Tcpb                bool        `json:"tcpb,omitempty"`
	Ulfd                bool        `json:"ulfd,omitempty"`
	Usip                bool        `json:"usip,omitempty"`
	Usnip               bool        `json:"usnip,omitempty"`
}
