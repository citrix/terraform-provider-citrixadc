resource "citrixadc_policymap" "tf_policymap" {
	mappolicyname = "tf_policymap"
	sd = "www.citrix.com"
	td = "www.google.com"
	su = "/www.citrix.com"
	tu = "/www.google.com"
}