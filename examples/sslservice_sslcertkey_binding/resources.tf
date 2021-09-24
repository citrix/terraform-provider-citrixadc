resource "citrixadc_sslcertkey" "tf_certkey" {
	certkey = "tf_certkey"
	cert = "/nsconfig/ssl/ns-root.cert"
	key = "/nsconfig/ssl/ns-root.key"
	notificationperiod = 40
	expirymonitor = "ENABLED"
}

resource "citrixadc_sslservice" "demo_sslservice" {
	cipherredirect = "DISABLED"
	clientauth = "DISABLED"
	dh = "DISABLED"
	dhcount = 0
	dhkeyexpsizelimit = "DISABLED"
	dtls12 = "DISABLED"
	ersa = "DISABLED"
	redirectportrewrite = "DISABLED"
	serverauth = "ENABLED"
	servicename = citrixadc_service.tf_service.name
	sessreuse = "ENABLED"
	sesstimeout = 300
	snienable = "DISABLED"
	ssl2 = "DISABLED"
	ssl3 = "ENABLED"
	sslredirect = "DISABLED"
	sslv2redirect = "DISABLED"
	tls1 = "ENABLED"
	tls11 = "ENABLED"
	tls12 = "ENABLED"
	tls13 = "DISABLED"
	
}

resource "citrixadc_lbvserver" "tf_lbvserver" {
	ipv46       = "10.10.10.44"
	name        = "tf_lbvserver"
	port        = 443
	servicetype = "SSL"
	sslprofile  = "ns_default_ssl_profile_frontend"
}

resource "citrixadc_service" "tf_service" {
	name = "tf_service"
	servicetype = "SSL"
	port = 443 
	lbvserver = citrixadc_lbvserver.tf_lbvserver.name
	ip = "10.77.33.22"

	depends_on = [citrixadc_sslcertkey.tf_certkey]

}
resource "citrixadc_sslservice_sslcertkey_binding" "tf_sslservice_sslcertkey_binding" {
	certkeyname = citrixadc_sslcertkey.tf_certkey.certkey
	servicename = citrixadc_service.tf_service.name
	ca = true
	ocspcheck = "Optional"
}