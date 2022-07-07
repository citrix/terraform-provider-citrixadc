resource "citrixadc_csvserver_auditsyslogpolicy_binding" "tf_csvserver_auditsyslogpolicy_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	policyname = citrixadc_auditsyslogpolicy.tf_auditsyslogpolicy.name
	priority = 5
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "HTTP"
}

resource "citrixadc_auditsyslogpolicy" "tf_auditsyslogpolicy" {
	name = "tf_auditsyslogpolicy"
	rule = "ns_true"
	action = citrixadc_auditsyslogaction.tf_syslogaction.name
}

resource "citrixadc_auditsyslogaction" "tf_syslogaction" {
	name = "tf_syslogaction"
	serverip = "10.78.60.33"
	serverport = 514
	loglevel = [
		"ERROR",
		"NOTICE",
	]
}