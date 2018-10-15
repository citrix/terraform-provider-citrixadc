package urlfiltering

type Urlfilteringparameter struct {
	Clouddblookuptimeout      int    `json:"clouddblookuptimeout,omitempty"`
	Cloudhost                 string `json:"cloudhost,omitempty"`
	Cloudkeepalivetimeout     int    `json:"cloudkeepalivetimeout,omitempty"`
	Cloudserverconnecttimeout int    `json:"cloudserverconnecttimeout,omitempty"`
	Hoursbetweendbupdates     int    `json:"hoursbetweendbupdates,omitempty"`
	Localdatabasethreads      int    `json:"localdatabasethreads,omitempty"`
	Maxnumberofcloudthreads   int    `json:"maxnumberofcloudthreads,omitempty"`
	Proxyhostip               string `json:"proxyhostip,omitempty"`
	Proxypassword             string `json:"proxypassword,omitempty"`
	Proxyport                 int    `json:"proxyport,omitempty"`
	Proxyusername             string `json:"proxyusername,omitempty"`
	Seeddbpath                string `json:"seeddbpath,omitempty"`
	Seeddbsizelevel           int    `json:"seeddbsizelevel,omitempty"`
	Timeofdaytoupdatedb       string `json:"timeofdaytoupdatedb,omitempty"`
}
