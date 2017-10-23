package ns

type Nstcpprofile struct {
	Ackonpush                 string `json:"ackonpush,omitempty"`
	Buffersize                int    `json:"buffersize,omitempty"`
	Delayedack                int    `json:"delayedack,omitempty"`
	Dynamicreceivebuffering   string `json:"dynamicreceivebuffering,omitempty"`
	Establishclientconn       string `json:"establishclientconn,omitempty"`
	Flavor                    string `json:"flavor,omitempty"`
	Initialcwnd               int    `json:"initialcwnd,omitempty"`
	Ka                        string `json:"ka,omitempty"`
	Kaconnidletime            int    `json:"kaconnidletime,omitempty"`
	Kamaxprobes               int    `json:"kamaxprobes,omitempty"`
	Kaprobeinterval           int    `json:"kaprobeinterval,omitempty"`
	Kaprobeupdatelastactivity string `json:"kaprobeupdatelastactivity,omitempty"`
	Maxburst                  int    `json:"maxburst,omitempty"`
	Maxpktpermss              int    `json:"maxpktpermss,omitempty"`
	Minrto                    int    `json:"minrto,omitempty"`
	Mptcp                     string `json:"mptcp,omitempty"`
	Mss                       int    `json:"mss,omitempty"`
	Nagle                     string `json:"nagle,omitempty"`
	Name                      string `json:"name,omitempty"`
	Oooqsize                  int    `json:"oooqsize,omitempty"`
	Pktperretx                int    `json:"pktperretx,omitempty"`
	Refcnt                    int    `json:"refcnt,omitempty"`
	Sack                      string `json:"sack,omitempty"`
	Sendbuffsize              int    `json:"sendbuffsize,omitempty"`
	Slowstartincr             int    `json:"slowstartincr,omitempty"`
	Syncookie                 string `json:"syncookie,omitempty"`
	Ws                        string `json:"ws,omitempty"`
	Wsval                     int    `json:"wsval,omitempty"`
}
