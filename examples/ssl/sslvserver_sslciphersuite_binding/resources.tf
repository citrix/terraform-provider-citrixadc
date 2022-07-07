resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding" {
	ciphername = "TLS1.2-ECDHE-RSA-AES128-GCM-SHA256"
	vservername = citrixadc_lbvserver.tf_sslvserver.name
}

resource "citrixadc_sslvserver_sslciphersuite_binding" "tf_sslvserver_sslciphersuite_binding2" {
	ciphername = "TLS1.3-CHACHA20-POLY1305-SHA256"
	vservername = citrixadc_lbvserver.tf_sslvserver.name
}

resource "citrixadc_lbvserver" "tf_sslvserver" {
	name = "tf_sslvserver"
	servicetype = "SSL"
	ipv46 = "5.5.5.5"
	port = 80
}