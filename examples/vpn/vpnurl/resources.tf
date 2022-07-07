resource "citrixadc_vpnurl" "url" {
  actualurl        = "www.citrix.com"
  appjson          = "xyz"
  applicationtype  = "CVPN"
  clientlessaccess = "OFF"
  comment          = "Testing"
  linkname         = "Description"
  ssotype          = "unifiedgateway"
  urlname          = "Firsturl"
  vservername      = "server1"
}
