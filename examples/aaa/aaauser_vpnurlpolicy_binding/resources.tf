resource "citrixadc_aaauser_vpnurlpolicy_binding" "tf_aaauser_vpnurlpolicy_binding" {
  username = "user1"
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