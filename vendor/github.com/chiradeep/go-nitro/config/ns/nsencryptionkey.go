package ns

type Nsencryptionkey struct {
	Comment  string `json:"comment,omitempty"`
	Iv       string `json:"iv,omitempty"`
	Keyvalue string `json:"keyvalue,omitempty"`
	Method   string `json:"method,omitempty"`
	Name     string `json:"name,omitempty"`
	Padding  string `json:"padding,omitempty"`
}
