resource "citrixadc_vpnalwaysonprofile" "tf_vpnalwaysonprofile" {
	name = "tf_vpnalwaysonprofile"
	clientcontrol = "ALLOW"
	locationbasedvpn = "Everywhere"
	networkaccessonvpnfailure = "fullAccess"
}