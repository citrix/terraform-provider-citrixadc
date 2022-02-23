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
resource "citrixadc_vpnurlpolicy" "citrixadc_vpnurlpolicy" {
  name   = "new_policy"
  rule   = "true"
  action = citrixadc_vpnurlaction.tf_vpnurlaction.name
}
resource "citrixadc_vpnglobal_vpnurlpolicy_binding" "tf_bind" {
  policyname = citrixadc_vpnurlpolicy.citrixadc_vpnurlpolicy.name
  priority = "20"
}
