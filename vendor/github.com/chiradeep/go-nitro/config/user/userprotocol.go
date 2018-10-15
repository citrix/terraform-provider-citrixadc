package user

type Userprotocol struct {
	Comment   string `json:"comment,omitempty"`
	Extension string `json:"extension,omitempty"`
	Name      string `json:"name,omitempty"`
	Transport string `json:"transport,omitempty"`
}
