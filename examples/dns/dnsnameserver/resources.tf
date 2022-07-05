resource "citrixadc_dnsnameserver" "dnsnameserver" {
	ip = "192.0.2.0"
    local = true
    state = "DISABLED"
    type = "UDP"
    dnsprofilename = "tf_pr"
}