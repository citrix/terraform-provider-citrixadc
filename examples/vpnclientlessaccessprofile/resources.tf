resource "citrixadc_vpnclientlessaccessprofile" "tf_vpnclientlessaccessprofile" {
	profilename = "tf_vpnclientlessaccessprofile"
	requirepersistentcookie = "ON"
}