resource "citrixadc_vpnvserver" "tf_tfvpnvserver" {
  name        = "tf_example_vserver"
  servicetype = "SSL"
  ipv46       = "9.6.77.8"
  port        = 443
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
resource "citrixadc_vpnvserver_vpnurlpolicy_binding" "tf_bind" {
  name                   = citrixadc_vpnvserver.tf_tfvpnvserver.name
  policy                 = citrixadc_vpnurlpolicy.tf_vpnurlpolicy.name
  priority               = 20
  gotopriorityexpression = "next"
  bindpoint              = "REQUEST" // doesnot unbind for other values
}