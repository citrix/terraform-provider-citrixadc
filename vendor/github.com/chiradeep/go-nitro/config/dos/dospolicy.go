package dos

type Dospolicy struct {
	Cltdetectrate int    `json:"cltdetectrate,omitempty"`
	Name          string `json:"name,omitempty"`
	Qdepth        int    `json:"qdepth,omitempty"`
}
