package app

type Application struct {
	Appname             string `json:"appname,omitempty"`
	Apptemplatefilename string `json:"apptemplatefilename,omitempty"`
	Deploymentfilename  string `json:"deploymentfilename,omitempty"`
}
