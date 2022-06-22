resource "citrixadc_sslservicegroup_sslcertkey_binding" "tf_sslservicegroup_sslcertkey_binding" {
	ca = false
	certkeyname = citrixadc_sslcertkey.tf_sslcertkey.certkey
	servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
}

resource "citrixadc_sslcertkey" "tf_sslcertkey" {
	certkey = "tf_sslcertkey"
	cert = "/var/tmp/certificate1.crt"
	key = "/var/tmp/key1.pem"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
	servicegroupname = "tf_servicegroup"
	servicetype = "SSL"
}
