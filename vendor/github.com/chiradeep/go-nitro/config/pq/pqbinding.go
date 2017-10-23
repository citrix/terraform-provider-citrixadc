package pq

type Pqbinding struct {
	Hits        int    `json:"hits,omitempty"`
	Policyname  string `json:"policyname,omitempty"`
	Polqdepth   int    `json:"polqdepth,omitempty"`
	Priority    int    `json:"priority,omitempty"`
	Qdepth      int    `json:"qdepth,omitempty"`
	Rule        string `json:"rule,omitempty"`
	Vservername string `json:"vservername,omitempty"`
	Weight      int    `json:"weight,omitempty"`
}
