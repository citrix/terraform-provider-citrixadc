package ns

type Nshmackey struct {
	Comment  string `json:"comment,omitempty"`
	Digest   string `json:"digest,omitempty"`
	Keyvalue string `json:"keyvalue,omitempty"`
	Name     string `json:"name,omitempty"`
}
