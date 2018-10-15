package cache

type Cacheparameter struct {
	Disklimit          int    `json:"disklimit,omitempty"`
	Enablebypass       string `json:"enablebypass,omitempty"`
	Enablehaobjpersist string `json:"enablehaobjpersist,omitempty"`
	Maxdisklimit       int    `json:"maxdisklimit,omitempty"`
	Maxmemlimit        int    `json:"maxmemlimit,omitempty"`
	Maxpostlen         int    `json:"maxpostlen,omitempty"`
	Memlimit           int    `json:"memlimit,omitempty"`
	Memlimitactive     int    `json:"memlimitactive,omitempty"`
	Prefetchcur        int    `json:"prefetchcur,omitempty"`
	Prefetchmaxpending int    `json:"prefetchmaxpending,omitempty"`
	Undefaction        string `json:"undefaction,omitempty"`
	Verifyusing        string `json:"verifyusing,omitempty"`
	Via                string `json:"via,omitempty"`
}
