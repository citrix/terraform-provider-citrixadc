resource "citrixadc_aaagroup_vpnurlpolicy_binding" "tf_aaagroup_vpnurlpolicy_binding" {
  groupname = "my_group"
  policy    = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
  priority  = 100
}

resource "citrixadc_vpnurlaction" "tf_vpnurlaction" {
  name             = "tf_vpnurlaction"
  linkname         = "new_link"
  actualurl        = "www.citrix.com"
  applicationtype  = "CVPN"
  clientlessaccess = "OFF"
  comment          = "Testing"
  ssotype          = "unifiedgateway"
  vservername      = "vserver1"
}
resource "citrixadc_vpnurlpolicy" "tf_vpnurlpolicy" {
  name   = "new_policy"
  rule   = "true"
  action = citrixadc_vpnurlaction.tf_vpnurlaction.name
}
