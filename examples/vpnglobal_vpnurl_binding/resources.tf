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
resource "citrixadc_vpnglobal_vpnurl_binding" "tf_bind" {
  urlname = citrixadc_vpnurl.url.urlname
}
