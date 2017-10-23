package policy

type Policymap struct {
	Mappolicyname string `json:"mappolicyname,omitempty"`
	Sd            string `json:"sd,omitempty"`
	Su            string `json:"su,omitempty"`
	Targetname    string `json:"targetname,omitempty"`
	Td            string `json:"td,omitempty"`
	Tu            string `json:"tu,omitempty"`
}
