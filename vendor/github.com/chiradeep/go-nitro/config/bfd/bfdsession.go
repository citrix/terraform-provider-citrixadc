package bfd

type Bfdsession struct {
	Admindown                         bool   `json:"admindown,omitempty"`
	Currentownerpe                    int    `json:"currentownerpe,omitempty"`
	Localdiagnotic                    int    `json:"localdiagnotic,omitempty"`
	Localdiscriminator                int    `json:"localdiscriminator,omitempty"`
	Localip                           string `json:"localip,omitempty"`
	Localport                         int    `json:"localport,omitempty"`
	Minimumreceiveinterval            int    `json:"minimumreceiveinterval,omitempty"`
	Minimumtransmitinterval           int    `json:"minimumtransmitinterval,omitempty"`
	Multihop                          bool   `json:"multihop,omitempty"`
	Multiplier                        int    `json:"multiplier,omitempty"`
	Negotiatedminimumreceiveinterval  int    `json:"negotiatedminimumreceiveinterval,omitempty"`
	Negotiatedminimumtransmitinterval int    `json:"negotiatedminimumtransmitinterval,omitempty"`
	Originalownerpe                   int    `json:"originalownerpe,omitempty"`
	Ownernode                         int    `json:"ownernode,omitempty"`
	Passive                           bool   `json:"passive,omitempty"`
	Remotediscriminator               int    `json:"remotediscriminator,omitempty"`
	Remoteip                          string `json:"remoteip,omitempty"`
	Remotemultiplier                  int    `json:"remotemultiplier,omitempty"`
	Remoteport                        int    `json:"remoteport,omitempty"`
	State                             string `json:"state,omitempty"`
	Vlan                              int    `json:"vlan,omitempty"`
}
