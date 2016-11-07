package basic

type Location struct {
	Ipfrom            string `json:"ipfrom,omitempty"`
	Ipto              string `json:"ipto,omitempty"`
	Latitude          int    `json:"latitude,omitempty"`
	Longitude         int    `json:"longitude,omitempty"`
	Preferredlocation string `json:"preferredlocation,omitempty"`
	Q1label           string `json:"q1label,omitempty"`
	Q2label           string `json:"q2label,omitempty"`
	Q3label           string `json:"q3label,omitempty"`
	Q4label           string `json:"q4label,omitempty"`
	Q5label           string `json:"q5label,omitempty"`
	Q6label           string `json:"q6label,omitempty"`
}
