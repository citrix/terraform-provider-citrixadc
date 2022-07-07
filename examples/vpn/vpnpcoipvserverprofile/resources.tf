resource "citrixadc_vpnpcoipvserverprofile" "tf_vpnpcoipvserverprofile" {
  name        = "tf_vpnpcoipvserverprofile"
  logindomain = "domainname"
  udpport     = "802"
}