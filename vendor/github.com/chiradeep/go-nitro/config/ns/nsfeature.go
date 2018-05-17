package ns

type Nsfeature struct {
	Aaa           bool        `json:"aaa,omitempty"`
	Appflow       bool        `json:"appflow,omitempty"`
	Appfw         bool        `json:"appfw,omitempty"`
	Appqoe        bool        `json:"appqoe,omitempty"`
	Bgp           bool        `json:"bgp,omitempty"`
	Cf            bool        `json:"cf,omitempty"`
	Ch            bool        `json:"ch,omitempty"`
	Cloudbridge   bool        `json:"cloudbridge,omitempty"`
	Cmp           bool        `json:"cmp,omitempty"`
	Cr            bool        `json:"cr,omitempty"`
	Cs            bool        `json:"cs,omitempty"`
	Feature       interface{} `json:"feature,omitempty"`
	Gslb          bool        `json:"gslb,omitempty"`
	Hdosp         bool        `json:"hdosp,omitempty"`
	Htmlinjection bool        `json:"htmlinjection,omitempty"`
	Ic            bool        `json:"ic,omitempty"`
	Ipv6pt        bool        `json:"ipv6pt,omitempty"`
	Isis          bool        `json:"isis,omitempty"`
	Lb            bool        `json:"lb,omitempty"`
	Ospf          bool        `json:"ospf,omitempty"`
	Pq            bool        `json:"pq,omitempty"`
	Push          bool        `json:"push,omitempty"`
	Responder     bool        `json:"responder,omitempty"`
	Rewrite       bool        `json:"rewrite,omitempty"`
	Rip           bool        `json:"rip,omitempty"`
	Sc            bool        `json:"sc,omitempty"`
	Sp            bool        `json:"sp,omitempty"`
	Ssl           bool        `json:"ssl,omitempty"`
	Sslvpn        bool        `json:"sslvpn,omitempty"`
	Vpath         bool        `json:"vpath,omitempty"`
	Wl            bool        `json:"wl,omitempty"`
}
