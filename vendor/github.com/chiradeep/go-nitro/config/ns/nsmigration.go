package ns

type Nsmigration struct {
	Migrationendtime           string `json:"migrationendtime,omitempty"`
	Migrationrollbackstarttime string `json:"migrationrollbackstarttime,omitempty"`
	Migrationstarttime         string `json:"migrationstarttime,omitempty"`
	Migrationstatus            string `json:"migrationstatus,omitempty"`
}
