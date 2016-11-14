package lb

type Lbgroup struct {
	Backuppersistencetimeout int    `json:"backuppersistencetimeout,omitempty"`
	Cookiedomain             string `json:"cookiedomain,omitempty"`
	Cookiename               string `json:"cookiename,omitempty"`
	Name                     string `json:"name,omitempty"`
	Newname                  string `json:"newname,omitempty"`
	Persistencebackup        string `json:"persistencebackup,omitempty"`
	Persistencetype          string `json:"persistencetype,omitempty"`
	Persistmask              string `json:"persistmask,omitempty"`
	Rule                     string `json:"rule,omitempty"`
	Td                       int    `json:"td,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	V6persistmasklen         int    `json:"v6persistmasklen,omitempty"`
}
