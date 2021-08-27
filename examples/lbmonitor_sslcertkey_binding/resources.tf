resource "citrixadc_lbmonitor_sslcertkey_binding" "tf_lbmonitor_sslcertkey_binding" {
	monitorname = citrixadc_lbmonitor.tf_monitor.monitorname
	certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
}

resource "citrixadc_lbmonitor" "tf_monitor" {
	monitorname = "tf_monitor"
	type = "HTTP"
	sslprofile = "ns_default_ssl_profile_backend"
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	certkey = "tf_sslcertkey"
	cert = "/var/tmp/certificate1.crt"
	key = "/var/tmp/key1.pem"
}