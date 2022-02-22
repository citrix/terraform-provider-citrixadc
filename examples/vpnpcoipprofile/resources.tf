resource "citrixadc_vpnpcoipprofile" "tf_vpnpcoipprofile" {
  name               = "tf_vpnpcoipprofile"
  conserverurl       = "http://www.example.com"
  sessionidletimeout = 80
}
