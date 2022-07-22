resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_examplevserver"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnurl" "url" {
  urlname          = "Firsturl"
  actualurl        = "www.citrix.com"
  appjson          = "xyz"
  applicationtype  = "CVPN"
  clientlessaccess = "OFF"
  comment          = "Testing"
  linkname         = "Description"
  ssotype          = "unifiedgateway"
  vservername      = "server1"
}
resource "citrixadc_vpnvserver_vpnurl_binding" "tf_bind" {
  name    = citrixadc_vpnvserver.tf_vpnvserver.name
  urlname = citrixadc_vpnurl.url.urlname
}
