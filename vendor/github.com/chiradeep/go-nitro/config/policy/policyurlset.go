package policy

type Policyurlset struct {
	Canaryurl           string `json:"canaryurl,omitempty"`
	Comment             string `json:"comment,omitempty"`
	Delimiter           string `json:"delimiter,omitempty"`
	Imported            bool   `json:"imported,omitempty"`
	Interval            int    `json:"interval,omitempty"`
	Name                string `json:"name,omitempty"`
	Overwrite           bool   `json:"overwrite,omitempty"`
	Patterncount        int    `json:"patterncount,omitempty"`
	Privateset          bool   `json:"privateset,omitempty"`
	Rowseparator        string `json:"rowseparator,omitempty"`
	Subdomainexactmatch bool   `json:"subdomainexactmatch,omitempty"`
	Url                 string `json:"url,omitempty"`
}
