resource "citrixadc_sslservicegroup" "tf_sslservicegroup" {
	servicegroupname = citrixadc_servicegroup.tf_servicegroup.servicegroupname
	sesstimeout = 50
	sessreuse = "ENABLED"
	ssl3 = "ENABLED"
	snienable = "ENABLED"
	serverauth = "ENABLED"
	sendclosenotify = "YES"
	strictsigdigestcheck = "ENABLED"
}

resource "citrixadc_servicegroup" "tf_servicegroup" {
	servicegroupname = "tf_servicegroup"
	servicetype = "SSL"
}