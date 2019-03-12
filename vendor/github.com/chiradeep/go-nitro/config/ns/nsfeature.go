package ns

type Nsfeature struct {
	Aaa                bool        `json:"aaa,omitempty"`
	Adaptivetcp        bool        `json:"adaptivetcp,omitempty"`
	Appflow            bool        `json:"appflow,omitempty"`
	Appfw              bool        `json:"appfw,omitempty"`
	Appqoe             bool        `json:"appqoe,omitempty"`
	Bgp                bool        `json:"bgp,omitempty"`
	Cf                 bool        `json:"cf,omitempty"`
	Ch                 bool        `json:"ch,omitempty"`
	Ci                 bool        `json:"ci,omitempty"`
	Cloudbridge        bool        `json:"cloudbridge,omitempty"`
	Cmp                bool        `json:"cmp,omitempty"`
	Contentaccelerator bool        `json:"contentaccelerator,omitempty"`
	Cqa                bool        `json:"cqa,omitempty"`
	Cr                 bool        `json:"cr,omitempty"`
	Cs                 bool        `json:"cs,omitempty"`
	Feature            interface{} `json:"feature,omitempty"`
	Feo                bool        `json:"feo,omitempty"`
	Forwardproxy       bool        `json:"forwardproxy,omitempty"`
	Gslb               bool        `json:"gslb,omitempty"`
	Hdosp              bool        `json:"hdosp,omitempty"`
	Htmlinjection      bool        `json:"htmlinjection,omitempty"`
	Ic                 bool        `json:"ic,omitempty"`
	Ipv6pt             bool        `json:"ipv6pt,omitempty"`
	Isis               bool        `json:"isis,omitempty"`
	Lb                 bool        `json:"lb,omitempty"`
	Lsn                bool        `json:"lsn,omitempty"`
	Ospf               bool        `json:"ospf,omitempty"`
	Pq                 bool        `json:"pq,omitempty"`
	Push               bool        `json:"push,omitempty"`
	Rdpproxy           bool        `json:"rdpproxy,omitempty"`
	Rep                bool        `json:"rep,omitempty"`
	Responder          bool        `json:"responder,omitempty"`
	Rewrite            bool        `json:"rewrite,omitempty"`
	Rip                bool        `json:"rip,omitempty"`
	Rise               bool        `json:"rise,omitempty"`
	Sc                 bool        `json:"sc,omitempty"`
	Sp                 bool        `json:"sp,omitempty"`
	Ssl                bool        `json:"ssl,omitempty"`
	Sslinterception    bool        `json:"sslinterception,omitempty"`
	Sslvpn             bool        `json:"sslvpn,omitempty"`
	Urlfiltering       bool        `json:"urlfiltering,omitempty"`
	Videooptimization  bool        `json:"videooptimization,omitempty"`
	Wl                 bool        `json:"wl,omitempty"`
}
