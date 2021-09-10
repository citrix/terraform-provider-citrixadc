# Since the vpnvserver resource is not yet available on Terraform,
# the tf_vpnvserver vserver must be created by hand in order for the script to run correctly.
# You can do that by using the following Citrix ADC cli command:
# add vpn vserver tf_vpnvserver SSL

resource "citrixadc_csvserver_vpnvserver_binding" "tf_csvserver_vpnvserver_binding" {
	name = citrixadc_csvserver.tf_csvserver.name
	vserver = "tf_vpnvserver"
}

resource "citrixadc_csvserver" "tf_csvserver" {
	name = "tf_csvserver"
	ipv46 = "10.202.11.11"
	port = 8080
	servicetype = "SSL"
	sslprofile = citrixadc_sslprofile.tf_sslprofile.name
}

resource "citrixadc_sslprofile" "tf_sslprofile" {
	name = "tf_sslprofile"
	ecccurvebindings = []
}