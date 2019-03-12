package network

type Route6 struct {
	Active            bool        `json:"active,omitempty"`
	Advertise         string      `json:"advertise,omitempty"`
	Bgp               bool        `json:"bgp,omitempty"`
	Connected         bool        `json:"connected,omitempty"`
	Cost              int         `json:"cost,omitempty"`
	Data              bool        `json:"data,omitempty"`
	Data1             string      `json:"data1,omitempty"`
	Detail            bool        `json:"detail,omitempty"`
	Distance          int         `json:"distance,omitempty"`
	Dynamic           bool        `json:"dynamic,omitempty"`
	Failedprobes      int         `json:"failedprobes,omitempty"`
	Flags             bool        `json:"flags,omitempty"`
	Gateway           string      `json:"gateway,omitempty"`
	Gatewayname       string      `json:"gatewayname,omitempty"`
	Isis              bool        `json:"isis,omitempty"`
	Monitor           string      `json:"monitor,omitempty"`
	Monstatcode       int         `json:"monstatcode,omitempty"`
	Monstatparam1     int         `json:"monstatparam1,omitempty"`
	Monstatparam2     int         `json:"monstatparam2,omitempty"`
	Monstatparam3     int         `json:"monstatparam3,omitempty"`
	Msr               string      `json:"msr,omitempty"`
	Network           string      `json:"network,omitempty"`
	Ospfv3            bool        `json:"ospfv3,omitempty"`
	Ownergroup        string      `json:"ownergroup,omitempty"`
	Permanent         bool        `json:"permanent,omitempty"`
	Raroute           bool        `json:"raroute,omitempty"`
	Retain            int         `json:"retain,omitempty"`
	Rip               bool        `json:"rip,omitempty"`
	Routeowners       interface{} `json:"routeowners,omitempty"`
	Routetype         string      `json:"routetype,omitempty"`
	State             int         `json:"state,omitempty"`
	Static            bool        `json:"Static,omitempty"`
	Td                int         `json:"td,omitempty"`
	Totalfailedprobes int         `json:"totalfailedprobes,omitempty"`
	Totalprobes       int         `json:"totalprobes,omitempty"`
	Type              bool        `json:"type,omitempty"`
	Vlan              int         `json:"vlan,omitempty"`
	Vxlan             int         `json:"vxlan,omitempty"`
	Weight            int         `json:"weight,omitempty"`
}
