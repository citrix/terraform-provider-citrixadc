resource "citrixadc_vpnvserver" "tf_vpnvserver" {
  name        = "tf_example.com"
  servicetype = "SSL"
  ipv46       = "3.3.3.3"
  port        = 443
}
resource "citrixadc_vpnclientlessaccesspolicy" "tf_vpnclientlessaccesspolicy" {
  name        = "tf_vpnclientlessaccesspolicy"
  profilename = "ns_cvpn_default_profile"
  rule        = "true"
}
resource "citrixadc_vpnvserver_vpnclientlessaccesspolicy_binding" "tf_bind" {
  name      = citrixadc_vpnvserver.tf_vpnvserver.name
  policy    = citrixadc_vpnclientlessaccesspolicy.tf_vpnclientlessaccesspolicy.name
  priority  = 20
  bindpoint = "REQUEST" // unbind doesnot work for other values
}