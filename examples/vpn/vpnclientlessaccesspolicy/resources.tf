resource "citrixadc_vpnclientlessaccesspolicy" "tf_vpnclientlessaccesspolicy" {
	name = "tf_vpnclientlessaccesspolicy"
	profilename = "ns_cvpn_default_profile"
	rule = "true"
}