resource "citrixadc_aaagroup_vpnurl_binding" "tf_aaagroup_vpnurl_binding" {
  groupname = "my_group"
  urlname   = citrixadc_vpnurl.tf_url.urlname
}

resource "citrixadc_vpnurl" "tf_url" {
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