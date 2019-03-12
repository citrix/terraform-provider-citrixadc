package ns

type Nstcpparam struct {
	Ackonpush                           string      `json:"ackonpush,omitempty"`
	Autosyncookietimeout                int         `json:"autosyncookietimeout,omitempty"`
	Builtin                             interface{} `json:"builtin,omitempty"`
	Connflushifnomem                    string      `json:"connflushifnomem,omitempty"`
	Connflushthres                      int         `json:"connflushthres,omitempty"`
	Delayedack                          int         `json:"delayedack,omitempty"`
	Downstaterst                        string      `json:"downstaterst,omitempty"`
	Initialcwnd                         int         `json:"initialcwnd,omitempty"`
	Kaprobeupdatelastactivity           string      `json:"kaprobeupdatelastactivity,omitempty"`
	Learnvsvrmss                        string      `json:"learnvsvrmss,omitempty"`
	Limitedpersist                      string      `json:"limitedpersist,omitempty"`
	Maxburst                            int         `json:"maxburst,omitempty"`
	Maxdynserverprobes                  int         `json:"maxdynserverprobes,omitempty"`
	Maxpktpermss                        int         `json:"maxpktpermss,omitempty"`
	Maxsynackretx                       int         `json:"maxsynackretx,omitempty"`
	Maxsynhold                          int         `json:"maxsynhold,omitempty"`
	Maxsynholdperprobe                  int         `json:"maxsynholdperprobe,omitempty"`
	Maxtimewaitconn                     int         `json:"maxtimewaitconn,omitempty"`
	Minrto                              int         `json:"minrto,omitempty"`
	Mptcpchecksum                       string      `json:"mptcpchecksum,omitempty"`
	Mptcpclosemptcpsessiononlastsfclose string      `json:"mptcpclosemptcpsessiononlastsfclose,omitempty"`
	Mptcpconcloseonpassivesf            string      `json:"mptcpconcloseonpassivesf,omitempty"`
	Mptcpimmediatesfcloseonfin          string      `json:"mptcpimmediatesfcloseonfin,omitempty"`
	Mptcpmaxpendingsf                   int         `json:"mptcpmaxpendingsf,omitempty"`
	Mptcpmaxsf                          int         `json:"mptcpmaxsf,omitempty"`
	Mptcppendingjointhreshold           int         `json:"mptcppendingjointhreshold,omitempty"`
	Mptcprtostoswitchsf                 int         `json:"mptcprtostoswitchsf,omitempty"`
	Mptcpsfreplacetimeout               int         `json:"mptcpsfreplacetimeout,omitempty"`
	Mptcpsftimeout                      int         `json:"mptcpsftimeout,omitempty"`
	Mptcpusebackupondss                 string      `json:"mptcpusebackupondss,omitempty"`
	Msslearndelay                       int         `json:"msslearndelay,omitempty"`
	Msslearninterval                    int         `json:"msslearninterval,omitempty"`
	Nagle                               string      `json:"nagle,omitempty"`
	Oooqsize                            int         `json:"oooqsize,omitempty"`
	Pktperretx                          int         `json:"pktperretx,omitempty"`
	Recvbuffsize                        int         `json:"recvbuffsize,omitempty"`
	Sack                                string      `json:"sack,omitempty"`
	Slowstartincr                       int         `json:"slowstartincr,omitempty"`
	Synattackdetection                  string      `json:"synattackdetection,omitempty"`
	Synholdfastgiveup                   int         `json:"synholdfastgiveup,omitempty"`
	Tcpfastopencookietimeout            int         `json:"tcpfastopencookietimeout,omitempty"`
	Tcpmaxretries                       int         `json:"tcpmaxretries,omitempty"`
	Ws                                  string      `json:"ws,omitempty"`
	Wsval                               int         `json:"wsval,omitempty"`
}
