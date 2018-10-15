package ns

type Nsaptlicense struct {
	Bindtype       string      `json:"bindtype,omitempty"`
	Countavailable string      `json:"countavailable,omitempty"`
	Counttotal     string      `json:"counttotal,omitempty"`
	Dateexp        string      `json:"dateexp,omitempty"`
	Datepurchased  string      `json:"datepurchased,omitempty"`
	Datesa         string      `json:"datesa,omitempty"`
	Features       interface{} `json:"features,omitempty"`
	Id             string      `json:"id,omitempty"`
	Licensedir     string      `json:"licensedir,omitempty"`
	Name           string      `json:"name,omitempty"`
	Relevance      string      `json:"relevance,omitempty"`
	Response       string      `json:"response,omitempty"`
	Serialno       string      `json:"serialno,omitempty"`
	Sessionid      string      `json:"sessionid,omitempty"`
	Useproxy       string      `json:"useproxy,omitempty"`
}
